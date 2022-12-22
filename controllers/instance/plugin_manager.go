package instance

import (
	"context"
	"time"

	"github.com/entgigi/bundle-operator/api/v1alpha1"

	"github.com/entgigi/bundle-operator/common"
	"github.com/entgigi/bundle-operator/controllers/services"

	"k8s.io/apimachinery/pkg/runtime"
)

type PluginManager struct {
	Base       *common.BaseK8sStructure
	Conditions *services.ConditionService
}

func NewPluginManager(base *common.BaseK8sStructure, conditions *services.ConditionService) *PluginManager {
	return &PluginManager{
		Base:       base,
		Conditions: conditions,
	}
}

func (d *PluginManager) IsPluginApplied(ctx context.Context, cr *v1alpha1.EntandoBundleInstanceV2) bool {

	return d.Conditions.IsPluginCrApplied(ctx, cr)
}

func (d *PluginManager) IsPluginReady(ctx context.Context, cr *v1alpha1.EntandoBundleInstanceV2) bool {

	return d.Conditions.IsPluginCrReady(ctx, cr)
}

func (d *PluginManager) ApplyPlugin(ctx context.Context, cr *v1alpha1.EntandoBundleInstanceV2, scheme *runtime.Scheme) error {
	/* FIXME
	applyError := d.ApplyKubeDeployment(ctx, cr, scheme)
	if applyError != nil {
		return applyError
	}
	*/
	return d.Conditions.SetConditionPluginCrApplied(ctx, cr)
}

func (d *PluginManager) CheckPluginCr(ctx context.Context, cr *v1alpha1.EntandoBundleInstanceV2) (bool, error) {
	time.Sleep(time.Second * 10)
	ready := true
	// check condition "Available" is "True"
	if ready {
		return ready, d.Conditions.SetConditionPluginCrReady(ctx, cr)
	}

	return ready, nil

}
