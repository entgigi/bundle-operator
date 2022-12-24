package services

import (
	"context"

	"github.com/entgigi/bundle-operator/api/v1alpha1"
	"github.com/entgigi/bundle-operator/bundles"
)

type BundleService struct {
}

func NewBundleService() *ConditionService {
	return &ConditionService{}
}

func (bs *BundleService) CheckBundleSignature(ctx context.Context, cr *v1alpha1.EntandoBundleInstanceV2) error {

	return nil
}

func (bs *BundleService) GetComponents(ctx context.Context, cr *v1alpha1.EntandoBundleInstanceV2) ([]bundles.Component, error) {

	return nil, nil
}
