package environment

import (
	"testing"

	"github.com/stretchr/testify/require"
	"gopkg.in/yaml.v3"
)

func Test_FloatVariable(t *testing.T) {
	t.Parallel()

	const (
		singleFloatValue   = 1.1
		singleFloat32Value = float32(singleFloatValue)
	)

	type testCase struct {
		val      any
		expected float64
	}

	testCases := map[string]testCase{
		"float64": {
			val:      singleFloatValue,
			expected: singleFloatValue,
		},
		"float32": {
			val:      singleFloat32Value,
			expected: singleFloatValue,
		},

		"negative_float64": {
			val:      -singleFloatValue,
			expected: -singleFloatValue,
		},
		"negative_float32": {
			val:      -singleFloat32Value,
			expected: -singleFloatValue,
		},
	}

	for name, tc := range testCases {
		tc := tc
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			actual := MustNewVariable(varName, tc.val)

			expected := &Variable{
				Name: varName,
				Type: VariableTypeFloat,
				Value: Value{
					val: &floatValue{
						v: tc.expected,
					},
				},
			}

			require.Equal(t, expected, actual)

			marshalled, err := yaml.Marshal(actual)
			require.NoError(t, err)

			var expectedYaml string
			if tc.expected > 0 {
				expectedYaml = `
name: test_variable
type: float
value: 1.1
`
			} else {
				expectedYaml = `
name: test_variable
type: float
value: -1.1
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

func Test_FloatSliceVariable(t *testing.T) {
	t.Parallel()

	targetSlice := []float64{
		-1.1,
		1.1,
	}

	type testCase struct {
		val      any
		expected []float64
	}

	testCases := map[string]testCase{
		"float64": {
			val:      targetSlice,
			expected: targetSlice,
		},
		"float32": {
			val: []float32{
				-1.1,
				1.1,
			},
			expected: targetSlice,
		},
	}

	expectedYaml := `
name: test_variable
type: float
value: [-1.1, 1.1]
`

	for name, tc := range testCases {
		tc := tc
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			actual := MustNewVariable(varName, tc.val)

			expected := &Variable{
				Name: varName,
				Type: VariableTypeFloat,
				Value: Value{
					val: &floatSliceValue{
						v: tc.expected,
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
		})
	}

}
