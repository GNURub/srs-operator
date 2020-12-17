/*


Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package controllers

import (
	"bytes"
	"context"
	"reflect"
	"text/template"

	"github.com/go-logr/logr"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/types"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/reconcile"

	streamingv1alpha1 "github.com/GNURub/srs-operator/apis/streaming/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
)

// SRSClusterReconciler reconciles a SRSCluster object
type SRSClusterReconciler struct {
	client.Client
	Log    logr.Logger
	Scheme *runtime.Scheme
}

// +kubebuilder:rbac:groups=streaming.srs,resources=srsclusters,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=streaming.srs,resources=srsclusters/status,verbs=get;update;patch

func (r *SRSClusterReconciler) Reconcile(req ctrl.Request) (ctrl.Result, error) {
	ctx := context.Background()
	log := r.Log.WithValues("srscluster", req.NamespacedName)

	// Fetch the SRSCluster instance
	srsInstance := &streamingv1alpha1.SRSCluster{}
	err := r.Get(ctx, req.NamespacedName, srsInstance)
	if err != nil {
		if errors.IsNotFound(err) {
			// Request object not found, could have been deleted after reconcile request.
			// Owned objects are automatically garbage collected. For additional cleanup logic use finalizers.
			// Return and don't requeue
			r.Log.Info("SRSCluster resource not found. Ignoring since object must be deleted")
			return reconcile.Result{}, nil
		}
		// Error reading the object - requeue the request.
		r.Log.Error(err, "Failed to get SRSCluster")
		return reconcile.Result{}, err
	}

	serviceName := "srs_cluster_svc"

	if len(srsInstance.Spec.ServiceName) > 0 {
		serviceName = srsInstance.Spec.ServiceName
	}

	// Check if the ConfigMap already exists, if not create a new one
	configFound := &corev1.ConfigMap{}
	err = r.Get(ctx, types.NamespacedName{Name: srsInstance.Name, Namespace: srsInstance.Namespace}, configFound)
	if err != nil && errors.IsNotFound(err) {
		// Define a new deployment
		config := r.statefulSetForSRSConf(srsInstance, serviceName)
		log.Info("Creating a new ConfigMap", "ConfigMap.Namespace", config.Namespace, "ConfigMap.Name", config.Name)
		err = r.Create(ctx, config)
		if err != nil {
			log.Error(err, "Failed to create new ConfigMap", "ConfigMap.Namespace", config.Namespace, "ConfigMap.Name", config.Name)
			return ctrl.Result{}, err
		}
	} else if err != nil {
		log.Error(err, "Failed to get ConfigMap")
		return ctrl.Result{}, err
	}

	// Check if the stateful already exists, if not create a new one
	statefulSetfound := &appsv1.StatefulSet{}
	err = r.Get(ctx, types.NamespacedName{Name: srsInstance.Name, Namespace: srsInstance.Namespace}, statefulSetfound)
	if err != nil && errors.IsNotFound(err) {
		// Define a new deployment
		stat := r.statefulSetForSRS(srsInstance, serviceName)
		log.Info("Creating a new StatefulSet", "StatefulSet.Namespace", stat.Namespace, "StatefulSet.Name", stat.Name)
		err = r.Create(ctx, stat)
		if err != nil {
			log.Error(err, "Failed to create new StatefulSet", "StatefulSet.Namespace", stat.Namespace, "StatefulSet.Name", stat.Name)
			return ctrl.Result{}, err
		}
		// StatefulSet created successfully - return and requeue
		return ctrl.Result{Requeue: true}, nil
	} else if err != nil {
		log.Error(err, "Failed to get StatefulSet")
		return ctrl.Result{}, err
	}

	// Ensure the deployment size is the same as the spec
	size := srsInstance.Spec.Size
	if *statefulSetfound.Spec.Replicas != size {
		statefulSetfound.Spec.Replicas = &size
		err = r.Update(ctx, statefulSetfound)
		if err != nil {
			log.Error(err, "Failed to update StatefulSet", "StatefulSet.Namespace", statefulSetfound.Namespace, "StatefulSet.Name", statefulSetfound.Name)
			return ctrl.Result{}, err
		}
		// Spec updated - return and requeue
		return ctrl.Result{Requeue: true}, nil
	}

	// Update the SRSCluster status with the pod names
	// List the pods for this SRSCluster's StatefulSet
	podList := &corev1.PodList{}
	listOpts := []client.ListOption{
		client.InNamespace(srsInstance.Namespace),
		client.MatchingLabels(labelsForSRSCluster(srsInstance.Name)),
	}
	if err = r.List(ctx, podList, listOpts...); err != nil {
		log.Error(err, "Failed to list pods", "SRSCluster.Namespace", srsInstance.Namespace, "SRSCluster.Name", srsInstance.Name)
		return ctrl.Result{}, err
	}
	podNames := getPodNames(podList.Items)

	// Update status.Nodes if needed
	if !reflect.DeepEqual(podNames, srsInstance.Status.PodNames) {
		srsInstance.Status.PodNames = podNames
		err := r.Status().Update(ctx, srsInstance)
		if err != nil {
			log.Error(err, "Failed to update SRSCluster status")
			return ctrl.Result{}, err
		}
	}

	return ctrl.Result{}, nil
}

func pointerInt32(n int32) *int32 {
	var i = new(int32)
	*i = n
	return i
}

// statefulSetForSRSConf returns a srs ConfigMap with the same name/namespace as the cr
func (r *SRSClusterReconciler) statefulSetForSRSConf(cr *streamingv1alpha1.SRSCluster, serviceName string) *corev1.ConfigMap {
	if cr.Spec.Config == nil {
		cr.Spec.Config = &streamingv1alpha1.SRSConfigSpec{
			Daemon: false,
			Listen: pointerInt32(1935),
			API: &streamingv1alpha1.SRSConfigServerSpec{
				Enabled: true,
				Listen:  pointerInt32(1985),
			},
			Server: &streamingv1alpha1.SRSConfigServerSpec{
				Enabled: true,
				Listen:  pointerInt32(8080),
			},
		}
	}

	var i int32
	origins := []string{}
	for i = 0; i < cr.Spec.Size; i++ {
		origins = append(origins, cr.Name+"-"+string(i))
	}
	cr.Spec.Config.Origins = origins

	// Configuration srs
	tmplConfig, err := template.New("").Parse(`
		listen              1935;
		max_connections     1000;
		daemon              off;

		{{if .API}}
		http_api {
			enabled         on;
			listen          1985;
		}
		{{end}}

		{{if .Server}}
		http_server {
			enabled         on;
			listen          8080;
		}
		{{end}}

		vhost __defaultVhost__ {
			enabled             on;
			mix_correct         on;

			cluster {
				mode            local;
				origin_cluster  on;
				coworkers        {{range .Origins}}{{.}}.` + serviceName + `:1985{{end}};
			}

			hls {
				enabled         on;
				hls_fragment    6;
				hls_window      30;
				hls_path        ./objs/nginx/html;
				hls_m3u8_file   [app]/default/[stream].m3u8;
				hls_ts_file     [app]/default/[stream]-[timestamp].ts;
				hls_cleanup     off;
				hls_nb_notify   64;
				hls_wait_keyframe       on;
			}
		}
	`)

	if err != nil {
		panic(err)
	}

	var tpl bytes.Buffer
	err = tmplConfig.Execute(&tpl, cr.Spec.Config)

	ls := labelsForSRSCluster(cr.Name)
	return &corev1.ConfigMap{
		ObjectMeta: metav1.ObjectMeta{
			Namespace: cr.Namespace,
			Labels:    ls,
		},
		Data: map[string]string{
			"srs.conf": tpl.String(),
		},
	}
}

// statefulSetForSRS returns a srs StatefulSet with the same name/namespace as the cr
func (r *SRSClusterReconciler) statefulSetForSRS(cr *streamingv1alpha1.SRSCluster, serviceName string) *appsv1.StatefulSet {
	ls := labelsForSRSCluster(cr.Name)
	img := "ossrs/srs:v4.0.56"
	if len(cr.Spec.Image) > 0 {
		img = cr.Spec.Image
	}

	return &appsv1.StatefulSet{
		ObjectMeta: metav1.ObjectMeta{
			GenerateName: cr.Name,
			Namespace:    cr.Namespace,
			Labels:       ls,
		},
		Spec: appsv1.StatefulSetSpec{
			ServiceName: serviceName,
			Replicas:    &cr.Spec.Size,
			Selector: &metav1.LabelSelector{
				MatchLabels: ls,
			},
			Template: corev1.PodTemplateSpec{
				ObjectMeta: metav1.ObjectMeta{
					Labels: ls,
				},
				Spec: corev1.PodSpec{
					Containers: []corev1.Container{
						{
							Name:            "srs",
							Image:           img,
							ImagePullPolicy: corev1.PullIfNotPresent,
							Ports: []corev1.ContainerPort{
								{
									ContainerPort: 1935,
								},
								{
									ContainerPort: 1985,
								},
								{
									ContainerPort: 8080,
								},
							},
						},
					},
				},
			},
		},
	}
}

// labelsForSRSCluster returns the labels for selecting the resources
// belonging to the given RSOriginCluster CR name.
func labelsForSRSCluster(name string) map[string]string {
	return map[string]string{"app": "srs_cluster", "srs_cluster_cr": name}
}

// getPodNames returns the pod names of the array of pods passed in
func getPodNames(pods []corev1.Pod) []string {
	var podNames []string
	for _, pod := range pods {
		podNames = append(podNames, pod.Name)
	}
	return podNames
}

func (r *SRSClusterReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&streamingv1alpha1.SRSCluster{}).
		Complete(r)
}
