package matreshka

import (
	stderrors "errors"
	"sort"

	"github.com/godverv/matreshka/internal/env_parser"
)

var ErrNoAppName = stderrors.New("no app name")

func GenerateEnvironmentKeys(c AppConfig) (keys []string, values []any, err error) {
	if c.AppInfo.Name == "" {
		return nil, nil, ErrNoAppName
	}

	keys, values, err = env_parser.ExtractVariables(c.AppInfo.Name, c.Environment)
	if err != nil {
		return nil, nil, err
	}

	sort.Slice(values, func(i, j int) bool {
		return keys[i] < keys[j]
	})

	sort.Slice(keys, func(i, j int) bool {
		return keys[i] < keys[j]
	})

	return keys, values, nil
}
