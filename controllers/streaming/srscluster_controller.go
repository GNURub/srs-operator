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
	"context"
	"reflect"

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

	// Check if the stateful already exists, if not create a new one
	found := &appsv1.StatefulSet{}
	err = r.Get(ctx, types.NamespacedName{Name: srsInstance.Name, Namespace: srsInstance.Namespace}, found)
	if err != nil && errors.IsNotFound(err) {
		// Define a new deployment
		dep := r.statefulSetForSRS(srsInstance)
		log.Info("Creating a new StatefulSet", "StatefulSet.Namespace", dep.Namespace, "StatefulSet.Name", dep.Name)
		err = r.Create(ctx, dep)
		if err != nil {
			log.Error(err, "Failed to create new StatefulSet", "StatefulSet.Namespace", dep.Namespace, "StatefulSet.Name", dep.Name)
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
	if *found.Spec.Replicas != size {
		found.Spec.Replicas = &size
		err = r.Update(ctx, found)
		if err != nil {
			log.Error(err, "Failed to update StatefulSet", "StatefulSet.Namespace", found.Namespace, "StatefulSet.Name", found.Name)
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

// statefulSetForSRS returns a srs StatefulSet with the same name/namespace as the cr
func (r *SRSClusterReconciler) statefulSetForSRS(cr *streamingv1alpha1.SRSCluster) *appsv1.StatefulSet {

	ls := labelsForSRSCluster(cr.Name)
	img := "ossrs/srs:v4.0.56"
	if len(cr.Spec.Image) > 0 {
		img = cr.Spec.Image
	}

	serviceName := "srs_cluster_svc"

	if len(cr.Spec.ServiceName) > 0 {
		serviceName = cr.Spec.ServiceName
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
