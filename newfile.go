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
