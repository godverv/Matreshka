package tests

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/godverv/matreshka"
)

func Test_ReadConfig(t *testing.T) {
	c, err := matreshka.ParseConfig([]byte(resourcedConfig))
	require.NoError(t, err)

	fmt.Println(c)
}
