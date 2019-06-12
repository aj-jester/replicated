package util

import (
	"fmt"
	"github.com/pkg/errors"
	"gopkg.in/yaml.v2"
	"strings"
)

type HelmConverter interface {
	ConvertValues(helmValues string) (string, error)
}

func NewHelmConverter() HelmConverter {
	return &helmConverter{}
}


var _ HelmConverter = &helmConverter{}
type helmConverter struct {

}


func (c helmConverter) ConvertValues(in string) (string, error) {
	var values yaml.MapSlice
	err := yaml.Unmarshal([]byte(in), &values)
	if err != nil {
		return "", errors.Wrap(err, "unmarshal values")
	}

	var path []string

	result := c.convertValuesRec(values, path)

	marshalled, err := yaml.Marshal(result)
	if err != nil {
		return "", errors.Wrap(err, "marshal result")
	}
	return string(marshalled), nil
}

func (c helmConverter) convertValuesRec(in yaml.MapSlice, path []string) yaml.MapSlice {
	var acc yaml.MapSlice

	for _, item := range in {
		key, ok := item.Key.(string); if !ok {
			// skip non-string keys (log me?)
			continue
		}

		newPath := append(path, key)

		// todo support for more types
		switch t := item.Value.(type) {
		case int, string, bool:
			acc = append(acc, yaml.MapItem{Key: key,
				Value: fmt.Sprintf(
					"{{repl ConfigOption %q }}",
					strings.Join(newPath, "."),
					),
			})
		case yaml.MapSlice:
			acc = append(acc, yaml.MapItem{
				Key: key,
				Value: c.convertValuesRec(t, newPath),
			})
		default:
			// todo need a real logger here
			fmt.Printf("\nUnknown Helm type:\n%s: %T\n", key, t)
		}
	}

	return acc
}
