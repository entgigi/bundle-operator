package instance

import (
	"context"
	"io/ioutil"

	"github.com/entgigi/bundle-operator/api/v1alpha1"
	"github.com/entgigi/bundle-operator/controllers/applyer"
	"github.com/entgigi/bundle-operator/utility"

	"github.com/entgigi/bundle-operator/common"
	runtime "k8s.io/apimachinery/pkg/runtime"
	"k8s.io/client-go/discovery"
	"k8s.io/client-go/dynamic"
	"k8s.io/client-go/rest"
)

type Manifest struct {
	Base *common.BaseK8sStructure
}

func NewManifest(base *common.BaseK8sStructure) *Manifest {
	return &Manifest{
		Base: base,
	}
}

func (d *Manifest) ApplyManifest(ctx context.Context, cr *v1alpha1.EntandoBundleInstanceV2,
	scheme *runtime.Scheme,
	manifestPath string) error {

	// read yaml
	yfile, err := ioutil.ReadFile(manifestPath)
	if err != nil {
		return err
	}
	config, err := rest.InClusterConfig()
	if err != nil {
		return err
	}

	dynamicClient, err := dynamic.NewForConfig(config)
	if err != nil {
		return err
	}

	discoveryClient, err := discovery.NewDiscoveryClientForConfig(config)
	if err != nil {
		panic(err.Error())
	}
	// apply yaml
	ns, _ := utility.GetWatchNamespace()
	applyOptions := applyer.NewApplyOptions(dynamicClient, discoveryClient)
	if err := applyOptions.Apply(context.TODO(), ns, []byte(yfile)); err != nil {
		return err
	}

	return nil
}
