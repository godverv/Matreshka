package environment

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
	"gopkg.in/yaml.v3"
)

const varName = "test_variable"

func Test_StringVariable(t *testing.T) {
	t.Parallel()

	const singleValue = "single_value"

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
name: test_variable
type: string
value: single_value
`[1:]

	require.YAMLEq(t, string(marshalled), expectedYaml)

	var newVar Variable

	err = yaml.Unmarshal(marshalled, &newVar)
	require.NoError(t, err)
	require.Equal(t, *expect, newVar)

}

func Test_StringSliceVariable(t *testing.T) {
	type testCase struct {
		valueToPass any
		opts        []opt
	}

	multipleValue := []string{"1", "2", "3"}

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
name: test_variable
type: string
value: ["1", "2", "3"]
`

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
}

type Anon struct {
	B any `yaml:"b,flow"`
	A any `yaml:"a,flow"`
}

func (a *Anon) UnmarshalYAML(node *yaml.Node) error {
	return nil
}

func Test_TestingMarshalling(t *testing.T) {

	a := Anon{A: []int{1, 2, 3}, B: 2}

	b, err := yaml.Marshal(a)
	require.NoError(t, err)
	_ = b

	c := Anon{}
	err = yaml.Unmarshal(b, &c)
	require.NoError(t, err)
}
