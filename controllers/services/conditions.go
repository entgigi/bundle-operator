package services

import (
	"context"

	"github.com/entgigi/bundle-operator/api/v1alpha1"
	"github.com/entgigi/bundle-operator/common"
	"github.com/entgigi/bundle-operator/utility"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"sigs.k8s.io/controller-runtime/pkg/log"
)

const (
	CONDITION_INSTANCE_CR_APPLIED        = "InstanceCrApplied"
	CONDITION_INSTANCE_CR_APPLIED_REASON = "InstanceCrIsApplied"
	CONDITION_INSTANCE_CR_APPLIED_MSG    = "Your instance cr was applied"

	CONDITION_INSTANCE_CR_READY        = "InstanceCrReady"
	CONDITION_INSTANCE_CR_READY_REASON = "InstanceCrIsReady"
	CONDITION_INSTANCE_CR_READY_MSG    = "Your instance cr is ready"
)

type ConditionService struct {
	Base *common.BaseK8sStructure
}

func NewConditionService(base *common.BaseK8sStructure) *ConditionService {
	return &ConditionService{
		Base: base,
	}
}

func (cs *ConditionService) IsInstanceCrReady(ctx context.Context, cr *v1alpha1.EntandoBundleInstanceV2) bool {

	condition, observedGeneration := cs.getConditionStatus(ctx, cr, CONDITION_INSTANCE_CR_READY)

	return metav1.ConditionTrue == condition && observedGeneration == cr.Generation
}

func (cs *ConditionService) SetConditionInstanceCrReady(ctx context.Context, cr *v1alpha1.EntandoBundleInstanceV2) error {

	cs.deleteCondition(ctx, cr, CONDITION_INSTANCE_CR_READY)
	return utility.AppendCondition(ctx, cs.Base.Client, cr,
		CONDITION_INSTANCE_CR_READY,
		metav1.ConditionTrue,
		CONDITION_INSTANCE_CR_READY_REASON,
		CONDITION_INSTANCE_CR_READY_MSG,
		cr.Generation)
}

func (cs *ConditionService) IsInstanceCrApplied(ctx context.Context, cr *v1alpha1.EntandoBundleInstanceV2) bool {

	condition, observedGeneration := cs.getConditionStatus(ctx, cr, CONDITION_INSTANCE_CR_APPLIED)

	return metav1.ConditionTrue == condition && observedGeneration == cr.Generation
}

func (cs *ConditionService) SetConditionInstanceCrApplied(ctx context.Context, cr *v1alpha1.EntandoBundleInstanceV2) error {

	cs.deleteCondition(ctx, cr, CONDITION_INSTANCE_CR_APPLIED)
	return utility.AppendCondition(ctx, cs.Base.Client, cr,
		CONDITION_INSTANCE_CR_APPLIED,
		metav1.ConditionTrue,
		CONDITION_INSTANCE_CR_APPLIED_REASON,
		CONDITION_INSTANCE_CR_APPLIED_MSG,
		cr.Generation)
}

/*
func (cs *ConditionService) SetConditionPluginReadyTrue(ctx context.Context, cr *v1alpha1.EntandoPluginV2) error {
	return cs.setConditionPluginReady(ctx, cr, metav1.ConditionTrue)
}

func (cs *ConditionService) SetConditionPluginReadyUnknow(ctx context.Context, cr *v1alpha1.EntandoPluginV2) error {
	return cs.setConditionPluginReady(ctx, cr, metav1.ConditionUnknown)
}

func (cs *ConditionService) SetConditionPluginReadyFalse(ctx context.Context, cr *v1alpha1.EntandoPluginV2) error {
	return cs.setConditionPluginReady(ctx, cr, metav1.ConditionFalse)
}
*/

func (cs *ConditionService) getConditionStatus(ctx context.Context, cr *v1alpha1.EntandoBundleInstanceV2, typeName string) (metav1.ConditionStatus, int64) {

	var output metav1.ConditionStatus = metav1.ConditionUnknown
	var observedGeneration int64

	for _, condition := range cr.Status.Conditions {
		if condition.Type == typeName {
			output = condition.Status
			observedGeneration = condition.ObservedGeneration
		}
	}
	return output, observedGeneration
}

func (cs *ConditionService) deleteCondition(ctx context.Context, cr *v1alpha1.EntandoBundleInstanceV2, typeName string) error {

	log := log.FromContext(ctx)
	var newConditions = make([]metav1.Condition, 0)
	for _, condition := range cr.Status.Conditions {
		if condition.Type != typeName {
			newConditions = append(newConditions, condition)
		}
	}
	cr.Status.Conditions = newConditions

	err := cs.Base.Client.Status().Update(ctx, cr)
	if err != nil {
		log.Info("Application resource status update failed.")
	}
	return nil
}
