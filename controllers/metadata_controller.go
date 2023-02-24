/*
Copyright 2023.

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
	corev1 "k8s.io/api/core/v1"

	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/log"

	foundationsv1alpha1 "Le0exxx/leo-operator/api/v1alpha1"
)

// MetadataReconciler reconciles a Metadata object
type MetadataReconciler struct {
	client.Client
	Scheme *runtime.Scheme
}

//+kubebuilder:rbac:groups=foundations.le0exxx,resources=metadata,verbs=get;list;watch;create;update;patch;delete
//+kubebuilder:rbac:groups=foundations.le0exxx,resources=metadata/status,verbs=get;update;patch
//+kubebuilder:rbac:groups=foundations.le0exxx,resources=metadata/finalizers,verbs=update
//+kubebuilder:rbac:groups="",resources=pods,verbs=get;list;watch;create;update;patch;delete

// Reconcile is part of the main kubernetes reconciliation loop which aims to
// move the current state of the cluster closer to the desired state.
// TODO(user): Modify the Reconcile function to compare the state specified by
// the Metadata object against the actual cluster state, and then
// perform operations to make the cluster state reflect the state specified by
// the user.
//
// For more details, check Reconcile and its Result here:
// - https://pkg.go.dev/sigs.k8s.io/controller-runtime@v0.14.1/pkg/reconcile
func (r *MetadataReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	_ = log.FromContext(ctx)

	var myMetadata foundationsv1alpha1.Metadata

	if err := r.Get(ctx, req.NamespacedName, &myMetadata); err != nil {
		// The resource may have been deleted since the request was queued, so don't requeue
		return ctrl.Result{}, client.IgnoreNotFound(err)
	}

	// Retrieve all the pods under the specified namespace
	podList := &corev1.PodList{}
	if err := r.List(context.Background(), podList, client.InNamespace(req.NamespacedName.Name)); err != nil {
		panic(err)
	}

	for _, pod := range podList.Items {
		labels := pod.GetLabels()
		if labels == nil {
			labels = make(map[string]string)
		}
		labels["Team"] = myMetadata.Spec.Team
		labels["Contact"] = myMetadata.Spec.Contact
		pod.SetLabels(labels)
	}

	return ctrl.Result{}, nil
}

// SetupWithManager sets up the controller with the Manager.
func (r *MetadataReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&foundationsv1alpha1.Metadata{}).
		Complete(r)
}
