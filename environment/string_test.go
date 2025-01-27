package environment

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
	"gopkg.in/yaml.v3"
)

func Test_StringVariable(t *testing.T) {
	t.Parallel()

	const (
		varName     = "string_variable"
		singleValue = "single_value"
	)
	multipleValue := []string{"1", "2", "3"}
	_ = multipleValue

	t.Run("Single", func(t *testing.T) {
		actual := MustNewVariable(varName, singleValue)

		expect := &Variable{
			Name: varName,
			Type: VariableTypeStr,
			Value: Value{
				val: &stringValue{
					v: singleValue,
				},
			},
		}

		require.Equal(t, expect, actual)

		marshalled, err := yaml.Marshal(actual)
		require.NoError(t, err)

		expectedYaml := `
name: string_variable
type: string
value: single_value
`[1:]

		require.YAMLEq(t, string(marshalled), expectedYaml)
	})

	t.Run("Slice", func(t *testing.T) {
		type testCase struct {
			valueToPass any
			opts        []opt
		}

		testCases := map[string]testCase{
			"FromStringSlice": {
				valueToPass: multipleValue,
			},
			"FromStringSliceAsOneString": {
				valueToPass: "[" + strings.Join(multipleValue, ",") + "]",
			},
			"FromAnySlice": {
				valueToPass: func() any {
					anySlice := make([]any, 0, len(multipleValue))
					for _, v := range multipleValue {
						anySlice = append(anySlice, v)
					}

					return anySlice
				}(),
				opts: []opt{
					WithType(VariableTypeStr),
				},
			},
		}

		expect := &Variable{
			Name: varName,
			Type: VariableTypeStr,
			Value: Value{
				val: &stringSliceValue{
					v: multipleValue,
				},
			},
		}
		expectedYaml := `
name: string_variable
type: string
value: '[1,2,3]'
`[1:]

		for name, tc := range testCases {
			tc := tc
			t.Run(name, func(t *testing.T) {
				actual := MustNewVariable(varName, tc.valueToPass, tc.opts...)
				require.Equal(t, expect, actual)

				marshalled, err := yaml.Marshal(actual)
				require.NoError(t, err)

				require.YAMLEq(t, string(marshalled), expectedYaml)
			})
		}
	})
}
