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

package helmclient

import (
	"github.com/mittwald/go-helm-client/values"
	"github.com/pkg/errors"
	"helm.sh/helm/v3/pkg/getter"
	"sigs.k8s.io/yaml"
)

func (spec *ChartSpec) GetValuesMap(p getter.Providers) (map[string]interface{}, error) {
	valuesYaml := map[string]interface{}{}

	err := yaml.Unmarshal([]byte(spec.ValuesYaml), &valuesYaml)
	if err != nil {
		return nil, errors.Wrap(err, "Failed to Parse ValuesYaml")
	}

	valuesOptions, err := spec.ValuesOptions.MergeValues(p)
	if err != nil {
		return nil, errors.Wrap(err, "Failed to Parse ValuesOptions")
	}

	return values.MergeMaps(valuesYaml, valuesOptions), nil
}
