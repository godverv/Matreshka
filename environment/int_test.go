package environment

import (
	"math/rand"
	"testing"

	"github.com/stretchr/testify/require"
)

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
