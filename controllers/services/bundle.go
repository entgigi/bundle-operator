package services

import (
	"context"
	"io/ioutil"
	"strings"

	"github.com/entgigi/bundle-operator/api/v1alpha1"
	"github.com/entgigi/bundle-operator/bundles"
	"github.com/entgigi/bundle-operator/utility"
)

type BundleService struct {
}

func NewBundleService() *BundleService {
	return &BundleService{}
}

func (bs *BundleService) CheckBundleSignature(ctx context.Context, cr *v1alpha1.EntandoBundleV2) (map[string]string, error) {

	return nil, nil
}

func (bs *BundleService) GenerateBundleCode(cr *v1alpha1.EntandoBundleV2) string {
	s := utility.GenerateSha256(cr.Spec.Repository)
	return "bundle-" + strings.ToLower(utility.TruncateString(s, 8))
}

func (bs *BundleService) GetComponents(ctx context.Context, cr *v1alpha1.EntandoBundleInstanceV2) ([]bundles.Component, string, error) {
	/*
		repository := "docker.io/gigiozzz/bundle-test-op"
		concat := "@"
		digest := "sha256:70ba938d4e11f219fc9dc0424e3e55173419a1da51598b341bb2162ea088a8a4"
	*/
	dir, err := ioutil.TempDir("/tmp", "crane-"+cr.Spec.Digest+"-")
	if err != nil {
		return nil, dir, err
	}

	err = bundles.ExtractImageTo(cr.Spec.Repository+"@"+cr.Spec.Digest, dir)
	if err != nil {
		return nil, dir, err
	}

	bundleDescriptor, err := bundles.ReadBundleDescriptor(dir + "/descriptor.yaml")
	if err != nil {
		return nil, dir, err
	}

	return bundleDescriptor.Components, dir, nil

}
