//
//
// Tencent is pleased to support the open source community by making tRPC available.
//
// Copyright (C) 2023 Tencent.
// All rights reserved.
//
// If you have downloaded a copy of the tRPC source code from Tencent,
// please note that tRPC source code is licensed under the Apache 2.0 License,
// A copy of the Apache 2.0 License is included in this file.
//
//

package discovery

import (
	"testing"

	"github.com/NHYCRaymond/trpc-go/naming/discovery"
	"github.com/NHYCRaymond/trpc-naming-polarismesh/mock/mock_api"
	"github.com/NHYCRaymond/trpc-naming-polarismesh/mock/mock_model"

	"github.com/golang/mock/gomock"
	"github.com/NHYCRaymond/polaris-go/pkg/model"
	"github.com/stretchr/testify/assert"
)

func TestSetup(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	m := mock_api.NewMockSDKContext(ctrl)

	assert.Nil(t, Setup(m, &Config{Name: "polarismesh"}, true))

	assert.NotNil(t, discovery.Get("polarismesh"))
	assert.NotNil(t, discovery.DefaultDiscovery)
}

func TestList(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	m := mock_api.NewMockConsumerAPI(ctrl)

	insts := []model.Instance{mock_model.NewMockInstance(ctrl)}
	resp := &model.InstancesResponse{
		Instances: insts,
	}
	m.EXPECT().GetInstances(gomock.Any()).Return(resp, nil).AnyTimes()
	d := &Discovery{
		consumer: m,
	}

	list, err := d.List("service", discovery.WithNamespace("namespace"))
	assert.Nil(t, err)
	assert.Len(t, list, 1)

	_, err = d.List("service")
	assert.NotNil(t, err)
}
