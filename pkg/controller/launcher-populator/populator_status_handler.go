package launcherpopulator

import (
	"context"
	"slices"

	fmav1alpha1 "github.com/llm-d-incubation/llm-d-fast-model-actuation/api/fma/v1alpha1"
	applyfmav1alpha1 "github.com/llm-d-incubation/llm-d-fast-model-actuation/pkg/generated/applyconfiguration/fma/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// setLPPStatusErrors sets the LauncherPopulationPolicy's Status.Errors to desiredErrors using
// Server-Side Apply. Pass nil or an empty slice to clear all errors.
// The call is skipped when the current Status already matches (same errors and same
// ObservedGeneration), avoiding unnecessary API round-trips.
func (ctl *controller) setLPPStatusErrors(ctx context.Context, lpp *fmav1alpha1.LauncherPopulationPolicy, desiredErrors []string) error {
	if lpp.Status.ObservedGeneration == lpp.Generation && slices.Equal(lpp.Status.Errors, desiredErrors) {
		// Status already reflects the desired state; nothing to do.
		return nil
	}
	apply := applyfmav1alpha1.LauncherPopulationPolicy(lpp.Name, lpp.Namespace).
		WithStatus(applyfmav1alpha1.LauncherPopulationPolicyStatus().
			WithObservedGeneration(lpp.Generation).
			WithErrors(desiredErrors...))
	_, err := ctl.fmaclient.LauncherPopulationPolicies(lpp.Namespace).ApplyStatus(ctx, apply, metav1.ApplyOptions{
		FieldManager: ControllerName,
		Force:        true,
	})
	return err
}

// setLCStatusErrors sets the LauncherConfig's Status.Errors to desiredErrors using
// Server-Side Apply. Pass nil or an empty slice to clear all errors.
// The call is skipped when the current Status already matches (same errors and same
// ObservedGeneration), avoiding unnecessary API round-trips.
func (ctl *controller) setLCStatusErrors(ctx context.Context, lc *fmav1alpha1.LauncherConfig, desiredErrors []string) error {
	if lc.Status.ObservedGeneration == lc.Generation && slices.Equal(lc.Status.Errors, desiredErrors) {
		// Status already reflects the desired state; nothing to do.
		return nil
	}
	apply := applyfmav1alpha1.LauncherConfig(lc.Name, lc.Namespace).
		WithStatus(applyfmav1alpha1.LauncherConfigStatus().
			WithObservedGeneration(lc.Generation).
			WithErrors(desiredErrors...))
	_, err := ctl.fmaclient.LauncherConfigs(lc.Namespace).ApplyStatus(ctx, apply, metav1.ApplyOptions{
		FieldManager: ControllerName,
		Force:        true,
	})
	return err
}
