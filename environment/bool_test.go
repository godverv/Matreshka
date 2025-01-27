package environment

import (
	"testing"

	"github.com/stretchr/testify/require"
	"gopkg.in/yaml.v3"
)

func Test_BoolValue(t *testing.T) {
	t.Parallel()

	type testCase struct {
		val      any
		expected bool
	}

	testCases := map[string]testCase{
		"true": {
			val:      true,
			expected: true,
		},
		"false": {
			val:      false,
			expected: false,
		},
	}

	for name, tc := range testCases {
		tc := tc
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			actual := MustNewVariable(varName, tc.val)

			expected := &Variable{
				Name: varName,
				Type: VariableTypeBool,
				Value: Value{
					val: &boolValue{
						v: tc.expected,
					},
				},
			}

			require.Equal(t, expected, actual)

			marshalled, err := yaml.Marshal(actual)
			require.NoError(t, err)

			var expectedYaml string
			if tc.expected {
				expectedYaml = `
name: test_variable
type: bool
value: true
`
			} else {
				expectedYaml = `
name: test_variable
type: bool
value: false
`
			}

			require.YAMLEq(t, string(marshalled), expectedYaml)

			var newVar Variable
			err = yaml.Unmarshal(marshalled, &newVar)
			require.NoError(t, err)

			require.Equal(t, expected, &newVar)
		})
	}

}

func Test_BoolSliceVariable(t *testing.T) {
	expectedYaml := `
name: test_variable
type: bool
value: [true, false]
`

	actual := MustNewVariable(varName, []bool{true, false})

	expected := &Variable{
		Name: varName,
		Type: VariableTypeBool,
		Value: Value{
			val: &boolSliceValue{
				v: []bool{true, false},
			},
		},
	}

	require.Equal(t, expected, actual)

	marshalled, err := yaml.Marshal(actual)
	require.NoError(t, err)

	require.YAMLEq(t, string(marshalled), expectedYaml)

	var newVar Variable
	err = yaml.Unmarshal(marshalled, &newVar)
	require.NoError(t, err)

	require.Equal(t, expected, &newVar)
}
