package instance

import (
	"context"
	"time"

	"github.com/entgigi/bundle-operator/api/v1alpha1"
	"github.com/entgigi/bundle-operator/common"
	"github.com/entgigi/bundle-operator/controllers/services"
	"github.com/entgigi/bundle-operator/utility"

	pluginapi "github.com/entgigi/plugin-operator/api/v1alpha1"

	"k8s.io/apimachinery/pkg/api/errors"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/types"
	ctrl "sigs.k8s.io/controller-runtime"
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

	return d.Conditions.IsPluginCrApplied(ctx, cr, d.GenPluginCode(cr))
}

func (d *PluginManager) IsPluginReady(ctx context.Context, cr *v1alpha1.EntandoBundleInstanceV2) bool {

	return d.Conditions.IsPluginCrReady(ctx, cr)
}

func (d *PluginManager) ApplyPlugin(ctx context.Context, cr *v1alpha1.EntandoBundleInstanceV2, scheme *runtime.Scheme) error {

	basePluginCr := d.buildPluginCr(cr, scheme)
	pluginCr := &pluginapi.EntandoPluginV2{}

	err, isUpgrade := d.isCrUpgrade(ctx, cr, pluginCr)
	if err != nil {
		return err
	}

	var applyError error
	if isUpgrade {
		pluginCr.Spec = basePluginCr.Spec
		applyError = d.Base.Client.Update(ctx, pluginCr)

	} else {
		applyError = d.Base.Client.Create(ctx, basePluginCr)
	}

	if applyError != nil {
		return applyError
	}

	return d.Conditions.SetConditionPluginCrApplied(ctx, cr, d.GenPluginCode(cr))
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

func (d *PluginManager) isCrUpgrade(ctx context.Context, cr *v1alpha1.EntandoBundleInstanceV2, pluginCr *pluginapi.EntandoPluginV2) (error, bool) {
	err := d.Base.Client.Get(ctx, types.NamespacedName{Name: d.GenPluginCode(cr), Namespace: cr.GetNamespace()}, pluginCr)
	if errors.IsNotFound(err) {
		return nil, false
	}
	return err, true
}

func (d *PluginManager) GenPluginId(cr *v1alpha1.EntandoBundleInstanceV2) string {
	pluginFullRepo := cr.Spec.Repository + "@" + cr.Spec.Digest
	s := utility.GenerateSha256(pluginFullRepo)
	return utility.TruncateString(s, 8)
}

func (d *PluginManager) GenPluginCode(cr *v1alpha1.EntandoBundleInstanceV2) string {
	pluginId := d.GenPluginId(cr)
	pluginCode := utility.TruncateString("pn-"+pluginId+"-"+cr.Name, 220)
	return pluginCode
}

func (d *PluginManager) buildPluginCr(cr *v1alpha1.EntandoBundleInstanceV2, scheme *runtime.Scheme) *pluginapi.EntandoPluginV2 {
	pluginCode := d.GenPluginCode(cr)
	pluginCr := &pluginapi.EntandoPluginV2{
		ObjectMeta: metav1.ObjectMeta{
			Name:      pluginCode,
			Namespace: cr.GetNamespace(),
		},
		Spec: pluginapi.EntandoPluginV2Spec{},
	}
	// set owner
	ctrl.SetControllerReference(cr, pluginCr, scheme)
	return pluginCr
}
