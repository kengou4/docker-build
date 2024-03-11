// Licensed to the Apache Software Foundation (ASF) under one
// or more contributor license agreements.  See the NOTICE file
// distributed with this work for additional information
// regarding copyright ownership.  The ASF licenses this file
// to you under the Apache License, Version 2.0 (the
// "License"); you may not use this file except in compliance
// with the License.  You may obtain a copy of the License at
//
//   http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing,
// software distributed under the License is distributed on an
// "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
// KIND, either express or implied.  See the License for the
// specific language governing permissions and limitations
// under the License.

package mockhelmclient

import (
	"testing"

	"go.uber.org/mock/gomock"

	"helm.sh/helm/v3/pkg/release"
)

var mockedRelease = release.Release{Name: "test"}

// TestHelmClientInterfaces performs checks on the clients interface functions.
func TestHelmClientInterfaces(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockClient := NewMockClient(ctrl)
	if mockClient == nil {
		t.Fail()
	}

	t.Run("UpdateChartRepos", func(t *testing.T) {
		mockClient.EXPECT().UpdateChartRepos().Return(nil)
		err := mockClient.UpdateChartRepos()
		if err != nil {
			panic(err)
		}
	})

	t.Run("ListReleases", func(t *testing.T) {
		mockClient.EXPECT().ListDeployedReleases().Return([]*release.Release{&mockedRelease}, nil)
		r, err := mockClient.ListDeployedReleases()
		if err != nil {
			panic(err)
		}
		if len(r) == 0 {
			panic(err)
		}
	})

	t.Run("GetRelease", func(t *testing.T) {
		mockClient.EXPECT().GetRelease(mockedRelease.Name).Return(&release.Release{Name: mockedRelease.Name}, nil)
		r, err := mockClient.GetRelease(mockedRelease.Name)
		if err != nil {
			panic(err)
		}
		if r == nil {
			panic(err)
		}
	})

	t.Run("GetReleaseValues", func(t *testing.T) {
		m := make(map[string]interface{})
		mockClient.EXPECT().GetReleaseValues(mockedRelease.Name, true).
			Return(m, nil)
		rv, err := mockClient.GetReleaseValues(mockedRelease.Name, true)
		if err != nil {
			panic(err)
		}
		if rv == nil {
			panic(err)
		}
	})
}
