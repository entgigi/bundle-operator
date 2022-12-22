package instance

import (
	"context"
	"fmt"
	"time"

	"github.com/entgigi/bundle-operator/api/v1alpha1"
	"github.com/entgigi/bundle-operator/common"
	"github.com/entgigi/bundle-operator/controllers/services"
	"github.com/go-logr/logr"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/client-go/tools/record"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

const labelKey = "app"
const serverPortName = "server-port"

type ReconcileInstanceManager struct {
	Base      *common.BaseK8sStructure
	Scheme    *runtime.Scheme
	Recorder  record.EventRecorder
	Condition *services.ConditionService
}

func NewReconcileInstanceManager(client client.Client, log logr.Logger, scheme *runtime.Scheme, recorder record.EventRecorder) *ReconcileInstanceManager {
	base := &common.BaseK8sStructure{Client: client, Log: log}
	return &ReconcileInstanceManager{
		Base:      base,
		Scheme:    scheme,
		Recorder:  recorder,
		Condition: services.NewConditionService(base),
	}
}

func (r *ReconcileInstanceManager) MainReconcile(ctx context.Context, req ctrl.Request, cr *v1alpha1.EntandoBundleInstanceV2) (ctrl.Result, error) {

	log := r.Base.Log
	pluginManager := NewPluginManager(r.Base, r.Condition)

	if err := r.Condition.SetConditionInstanceReadyUnknow(ctx, cr); err != nil {
		log.Info("error on set instance ready unknow")
		return ctrl.Result{}, err
	}

	// plugin done
	applied := pluginManager.IsPluginApplied(ctx, cr)

	if !applied {
		if err := pluginManager.ApplyPlugin(ctx, cr, r.Scheme); err != nil {
			log.Info("error ApplyPlugin reschedule reconcile", "error", err)
			r.Condition.SetConditionInstanceReadyFalse(ctx, cr)
			return ctrl.Result{}, err
		}
	}
	r.Recorder.Eventf(cr, "Normal", "Updated", fmt.Sprintf("Updated plugin cr %s/%s", req.Namespace, req.Name))

	// plugin ready
	var err error
	ready := pluginManager.IsPluginReady(ctx, cr)

	if !ready {
		if ready, err = pluginManager.CheckPluginCr(ctx, cr); err != nil {
			log.Info("error CheckPluginCr reschedule reconcile", "error", err)
			r.Condition.SetConditionInstanceReadyFalse(ctx, cr)
			return ctrl.Result{}, err
		}
		if !ready {
			log.Info("Plugin cr not ready reschedule operator", "seconds", 10)
			r.Recorder.Eventf(cr, "Warning", "NotReady", fmt.Sprintf("Plugin cr not ready %s/%s", req.Namespace, req.Name))
			r.Condition.SetConditionInstanceReadyFalse(ctx, cr)
			return ctrl.Result{Requeue: true, RequeueAfter: 10 * time.Second}, nil
		}
	}

	r.Condition.SetConditionInstanceReadyTrue(ctx, cr)
	return ctrl.Result{}, nil
}
