package environment

import (
	"testing"
	"time"

	"github.com/stretchr/testify/require"
	"gopkg.in/yaml.v3"
)

func Test_DurationValue(t *testing.T) {
	type testCase struct {
		req               any
		expected          time.Duration
		expectedYAMLValue string
	}

	testCases := map[string]testCase{
		"nano_sec_native": {
			req:               time.Nanosecond,
			expected:          time.Nanosecond,
			expectedYAMLValue: "1ns",
		},
		"micro_sec_native": {
			req:               time.Microsecond,
			expected:          time.Microsecond,
			expectedYAMLValue: "1µs",
		},
		"ml_sec_native": {
			req:               time.Millisecond,
			expected:          time.Millisecond,
			expectedYAMLValue: "1ms",
		},
		"sec_native": {
			req:               time.Second,
			expected:          time.Second,
			expectedYAMLValue: "1s",
		},
		"min_native": {
			req:               time.Minute,
			expected:          time.Minute,
			expectedYAMLValue: "1m0s",
		},
		"hour_native": {
			req:               time.Hour,
			expected:          time.Hour,
			expectedYAMLValue: "1h0m0s",
		},
	}

	for name, tc := range testCases {
		tc := tc
		t.Run(name, func(t *testing.T) {
			actual := MustNewVariable(varName, tc.req)

			expected := &Variable{
				Name: varName,
				Type: VariableTypeDuration,
				Value: Value{
					val: &durationValue{
						v: tc.expected,
					},
				},
			}

			require.Equal(t, expected, actual)

			marshalled, err := yaml.Marshal(actual)
			require.NoError(t, err)
			expectedYAML := `
name: test_variable
type: duration
value: ` + tc.expectedYAMLValue
			require.YAMLEq(t, expectedYAML, string(marshalled))
		})
	}
}
