/*
Copyright 2024 Train Conductor Authors.

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

	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/log"

	trainv1alpha1 "github.com/tedhtchang/train-conductor.git/api/v1alpha1"
)

// TuningJobReconciler reconciles a TuningJob object
type TuningJobReconciler struct {
	client.Client
	Scheme *runtime.Scheme
}

//+kubebuilder:rbac:groups=train.watsonx.ai,resources=tuningjobs,verbs=get;list;watch;create;update;patch;delete
//+kubebuilder:rbac:groups=train.watsonx.ai,resources=tuningjobs/status,verbs=get;update;patch
//+kubebuilder:rbac:groups=train.watsonx.ai,resources=tuningjobs/finalizers,verbs=update

// Reconcile is part of the main kubernetes reconciliation loop which aims to
// move the current state of the cluster closer to the desired state.
// TODO(user): Modify the Reconcile function to compare the state specified by
// the TuningJob object against the actual cluster state, and then
// perform operations to make the cluster state reflect the state specified by
// the user.
//
// For more details, check Reconcile and its Result here:
// - https://pkg.go.dev/sigs.k8s.io/controller-runtime@v0.17.0/pkg/reconcile
func (r *TuningJobReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	l := log.FromContext(ctx)

	// TODO(user): your logic here
	tj := &trainv1alpha1.TuningJob{}
	err := r.Get(ctx, req.NamespacedName, tj)
	if err != nil {
		l.Error(err, "Unable to fetch TuningJob")
	} else {
		l.Info("Reconciling TuningJob",
			"Name", tj.Name,
			"Namespace", tj.Namespace,
			"Model Name", tj.Spec.ModelName)
		return ctrl.Result{}, client.IgnoreNotFound(err)
	}

	return ctrl.Result{}, nil
}

// SetupWithManager sets up the controller with the Manager.
func (r *TuningJobReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&trainv1alpha1.TuningJob{}).
		Complete(r)
}
