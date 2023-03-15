package v1alpha1

import (
	ctrl "sigs.k8s.io/controller-runtime"
)

func (r *FlagSourceConfiguration) SetupWebhookWithManager(mgr ctrl.Manager) error {
	return ctrl.NewWebhookManagedBy(mgr).
		For(r).
		Complete()
}
