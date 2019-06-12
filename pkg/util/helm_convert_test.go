package util

import (
	"github.com/stretchr/testify/require"
	"strings"
	"testing"
)

func TestHelmBuildValues(t *testing.T) {
	tests := []struct {
		name string
		values string
		expect string
	}{
		{
			name: "empty",
			values: "",
			expect: "{}",
		},
		{
			name: "empty 2",
			values: "{}",
			expect: "{}",
		},
		{
			name: "one value",
			values: `reticulate_splines: please`,
			expect: `reticulate_splines: '{{repl ConfigOption "reticulate_splines" }}'`,
		},
		{
			name: "two values",
			values: `
reticulate_splines: please
deploy_adjunct_frombulator: nope
`,
			expect: `
reticulate_splines: '{{repl ConfigOption "reticulate_splines" }}'
deploy_adjunct_frombulator: '{{repl ConfigOption "deploy_adjunct_frombulator" }}'
`,
		},
		{
			name: "one nested value",
			values: `
reticulate_splines:
  depth: 4
`,
			expect: `
reticulate_splines:
  depth: '{{repl ConfigOption "reticulate_splines.depth" }}'
`,
		},
		{
			name: "multiple nested values",
			values: `
reticulate_splines:
  depth: 4
persistence:
  postgres:
    enable: false
  redis:
    enable: false
    replicas: 2
`,
			expect: `
reticulate_splines:
  depth: '{{repl ConfigOption "reticulate_splines.depth" }}'
persistence:
  postgres:
    enable: '{{repl ConfigOption "persistence.postgres.enable" }}'
  redis:
    enable: '{{repl ConfigOption "persistence.redis.enable" }}'
    replicas: '{{repl ConfigOption "persistence.redis.replicas" }}'
`,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			req := require.New(t)
			hc := &helmConverter{}
			actual, err := hc.ConvertValues(test.values)

			req.NoError(err, "convert helm values")
			req.Equal(strings.TrimSpace(test.expect), strings.TrimSpace(actual))
		})
	}
}

