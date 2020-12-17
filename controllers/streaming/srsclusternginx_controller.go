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

	"github.com/go-logr/logr"
	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/resource"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"

	streamingv1alpha1 "github.com/GNURub/srs-operator/apis/streaming/v1alpha1"
)

// SRSClusterNginxReconciler reconciles a SRSClusterNginx object
type SRSClusterNginxReconciler struct {
	client.Client
	Log    logr.Logger
	Scheme *runtime.Scheme
}

// +kubebuilder:rbac:groups=streaming.srs,resources=srsclusternginxes,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=streaming.srs,resources=srsclusternginxes/status,verbs=get;update;patch

func (r *SRSClusterNginxReconciler) Reconcile(req ctrl.Request) (ctrl.Result, error) {
	_ = context.Background()
	_ = r.Log.WithValues("srsclusternginx", req.NamespacedName)

	// your logic here

	return ctrl.Result{}, nil
}

// statefulSetForSRSNginx returns a srs StatefulSet with the same name/namespace as the cr
func (r *SRSClusterNginxReconciler) statefulSetForSRSNginx(cr *streamingv1alpha1.SRSClusterNginxSpec, serviceName string) *appsv1.Deployment {
	ls := labelsForSRSClusterNginx(cr.Name)

	return &appsv1.Deployment{
		ObjectMeta: metav1.ObjectMeta{
			GenerateName: cr.Name,
			Namespace:    cr.Namespace,
			Labels:       ls,
		},
		Spec: appsv1.DeploymentSpec{
			Replicas: cr.Spec.MinReplicas,
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
							Name:            "nginx",
							Image:           "nginx:latest",
							ImagePullPolicy: corev1.PullIfNotPresent,
							Ports: []corev1.ContainerPort{
								{
									ContainerPort: 80,
								},
							},
							Resources: corev1.ResourceRequirements{
								Requests: corev1.ResourceList{
									corev1.ResourceMemory: resource.MustParse("200Mi"),
									corev1.ResourceCPU:    resource.MustParse("10m"),
								},
								Limits: corev1.ResourceList{
									corev1.ResourceMemory: resource.MustParse("600Mi"),
									corev1.ResourceCPU:    resource.MustParse("300m"),
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
func labelsForSRSClusterNginx(name string) map[string]string {
	return map[string]string{"app": "srs_cluster_nginx", "srs_cluster_nginx_cr": name}
}

func (r *SRSClusterNginxReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&streamingv1alpha1.SRSClusterNginx{}).
		Complete(r)
}
