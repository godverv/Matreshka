package matreshka

import (
	"strings"

	"github.com/godverv/matreshka/internal/cases"
	"github.com/godverv/matreshka/internal/env_parser"
)

func GenerateGoConfigKeys(prefix string, c *AppConfig) ([]byte, error) {
	envKeys, err := env_parser.StructToEnvNames("", c)
	if err != nil {
		return nil, err
	}

	keysFromCfg := make(map[string]string, len(envKeys))
	for _, v := range envKeys {
		keysFromCfg[cases.SnakeToPascal(v)] = v
	}

	for goName, envName := range keysFromCfg {
		if _, ok := keysFromCfg[goName]; !ok {
			keysFromCfg[goName] = envName
		}
	}
	sb := &strings.Builder{}

	for key, v := range keysFromCfg {
		sb.WriteString(key + " = \"" + v + "\"\n")
	}

	return []byte(sb.String()), nil
}
