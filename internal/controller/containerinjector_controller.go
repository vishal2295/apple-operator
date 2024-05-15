/*
Copyright 2024.

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

package controller

import (
	"context"

	appsv1 "k8s.io/api/apps/v1"
	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/handler"
	"sigs.k8s.io/controller-runtime/pkg/log"
	"sigs.k8s.io/controller-runtime/pkg/source"

	applev1 "github.com/vishal2295/apple-operator/api/v1"
	v1 "github.com/vishal2295/apple-operator/api/v1"
)

// ContainerInjectorReconciler reconciles a ContainerInjector object
type ContainerInjectorReconciler struct {
	client.Client
	Scheme *runtime.Scheme
}

//+kubebuilder:rbac:groups=apple.dev.vishal2295,resources=containerinjectors,verbs=get;list;watch;create;update;patch;delete
//+kubebuilder:rbac:groups=apple.dev.vishal2295,resources=containerinjectors/status,verbs=get;update;patch
//+kubebuilder:rbac:groups=apple.dev.vishal2295,resources=containerinjectors/finalizers,verbs=update

// Reconcile is part of the main kubernetes reconciliation loop which aims to
// move the current state of the cluster closer to the desired state.
// TODO(user): Modify the Reconcile function to compare the state specified by
// the ContainerInjector object against the actual cluster state, and then
// perform operations to make the cluster state reflect the state specified by
// the user.
//
// For more details, check Reconcile and its Result here:
// - https://pkg.go.dev/sigs.k8s.io/controller-runtime@v0.17.3/pkg/reconcile
func (r *ContainerInjectorReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	_ = log.FromContext(ctx)

	// TODO(user): your logic here

	return ctrl.Result{}, nil
}

// SetupWithManager sets up the controller with the Manager.
func (r *ContainerInjectorReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&applev1.ContainerInjector{}).
		Watches(
			&source.Kind{Type: &appsv1.Deployment{}},
			handler.EnqueueRequestsFromMapFunc{r.GetAll}).
		Complete(r)
}

func (r *ContainerInjectorReconciler) GetAll(o client.Object) []ctrl.Request {
	result := []ctrl.Request{}

	injectorList := v1.ContainerInjectorList{}
	r.Client.List(context.Background(), &injectorList)

	for _, labeler := range injectorList.Items {
		result = append(result, ctrl.Request{NamespacedName: client.ObjectKey{Namespace: labeler.Namespace, Name: labeler.Name}})
	}

	return result
}
