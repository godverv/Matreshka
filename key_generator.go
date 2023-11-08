package matreshka

import (
	stderrors "errors"

	"github.com/godverv/matreshka/internal/env_parser"
)

var ErrNoAppName = stderrors.New("no app name")

func ExtractKeyValues(c AppConfig) (keys []string, values []any, err error) {
	if c.AppInfo.Name == "" {
		return nil, nil, ErrNoAppName
	}

	keys, values, err = env_parser.ExtractVariables(c.AppInfo.Name, c.Environment)
	if err != nil {
		return nil, nil, err
	}

	return keys, values, nil
}
