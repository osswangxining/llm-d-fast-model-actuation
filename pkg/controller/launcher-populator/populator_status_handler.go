package launcherpopulator

import (
	"context"

	fmav1alpha1 "github.com/llm-d-incubation/llm-d-fast-model-actuation/api/fma/v1alpha1"
	applyfmav1alpha1 "github.com/llm-d-incubation/llm-d-fast-model-actuation/pkg/generated/applyconfiguration/fma/v1alpha1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// reportLPPSelectorError writes a selector validation error into the LauncherPopulationPolicy's
// Status.Errors field using Server-Side Apply so the user can observe it via kubectl.
func (ctl *controller) reportLPPSelectorError(ctx context.Context, lpp *fmav1alpha1.LauncherPopulationPolicy, selectorErr error) error {
	apply := applyfmav1alpha1.LauncherPopulationPolicy(lpp.Name, lpp.Namespace).
		WithStatus(applyfmav1alpha1.LauncherPopulationPolicyStatus().
			WithObservedGeneration(lpp.Generation).
			WithErrors(selectorErr.Error()))
	_, err := ctl.fmaclient.LauncherPopulationPolicies(lpp.Namespace).ApplyStatus(ctx, apply, metav1.ApplyOptions{
		FieldManager: ControllerName,
		Force:        true,
	})
	return err
}

// clearLPPSelectorError clears any previously reported selector errors from the
// LauncherPopulationPolicy's Status.Errors field using Server-Side Apply.
func (ctl *controller) clearLPPSelectorError(ctx context.Context, lpp *fmav1alpha1.LauncherPopulationPolicy) error {
	if len(lpp.Status.Errors) == 0 {
		// Nothing to clear.
		return nil
	}
	apply := applyfmav1alpha1.LauncherPopulationPolicy(lpp.Name, lpp.Namespace).
		WithStatus(applyfmav1alpha1.LauncherPopulationPolicyStatus().
			WithObservedGeneration(lpp.Generation))
	_, err := ctl.fmaclient.LauncherPopulationPolicies(lpp.Namespace).ApplyStatus(ctx, apply, metav1.ApplyOptions{
		FieldManager: ControllerName,
		Force:        true,
	})
	return err
}
