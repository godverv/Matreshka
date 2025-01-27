package environment

import (
	"math/rand"
	"testing"

	"github.com/stretchr/testify/require"
	"gopkg.in/yaml.v3"
)

func Test_IntVariable(t *testing.T) {
	t.Parallel()

	const (
		varName          = "int_var"
		singleIntValue   = 1
		singleInt8Value  = int8(1)
		singleInt16Value = int16(1)
		singleInt32Value = int32(1)
		singleInt64Value = int64(1)
	)

	t.Run("Single", func(t *testing.T) {

		type testCase struct {
			val any
		}

		testCases := map[string]testCase{
			"int": {
				val: singleIntValue,
			},
			"int8": {
				val: singleInt8Value,
			},
			"int16": {
				val: singleInt16Value,
			},
			"int32": {
				val: singleInt32Value,
			},
			"int64": {
				val: singleInt64Value,
			},
		}

		expect := &Variable{
			Name: varName,
			Type: VariableTypeInt,
			Value: Value{
				val: &intValue{
					v: singleIntValue,
				},
			},
		}

		for name, tc := range testCases {
			tc := tc
			t.Run(name, func(t *testing.T) {
				t.Parallel()
				actual := MustNewVariable(varName, tc.val)

				require.Equal(t, expect, actual)

				marshalled, err := yaml.Marshal(actual)
				require.NoError(t, err)

				expectedYaml := `
name: int_var
type: int
value: 1
`[1:]

				require.YAMLEq(t, string(marshalled), expectedYaml)
			})
		}

	})
}

func Test_MarshalInt(t *testing.T) {
	t.Parallel()

	t.Run("single", func(t *testing.T) {
		t.Parallel()
		v := 1
		resp := marshalInt(v)
		require.Equal(t, resp, "1")
	})

	t.Run("array", func(t *testing.T) {
		t.Parallel()
		v := []int{1, 2, 3, 5, 7, 8, 9, 10, 12}
		rand.Shuffle(len(v), func(i, j int) {
			v[i], v[j] = v[j], v[i]
		})
		resp := marshalInt(v)
		require.Equal(t, resp, "[1-3,5,7-10,12]")
	})

	t.Run("empty_array", func(t *testing.T) {
		t.Parallel()
		v := []int{}
		resp := marshalInt(v)
		require.Equal(t, resp, "[]")
	})

}
