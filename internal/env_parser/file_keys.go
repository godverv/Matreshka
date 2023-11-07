package env_parser

import (
	"gopkg.in/yaml.v3"
)

type cfgKeysBuilder map[string]interface{}

func ExtractEnvNames(prefix string, src interface{}) (keys []string, values []any, err error) {
	cfgBytes, err := yaml.Marshal(src)
	if err != nil {
		return nil, nil, err
	}

	cfg := make(cfgKeysBuilder)
	err = yaml.Unmarshal(cfgBytes, cfg)
	if err != nil {
		return nil, nil, err
	}

	return cfg.extractVariables(prefix, cfg)
}

func (c cfgKeysBuilder) extractVariables(prefix string, m map[string]interface{}) (keys []string, values []any, err error) {
	for k, v := range m {
		if newMap, ok := v.(cfgKeysBuilder); ok {
			upperKeys, upperValues, err := c.extractVariables(prefix+"_"+k, newMap)
			if err != nil {
				return nil, nil, err
			}
			keys = append(keys, upperKeys...)
			values = append(values, upperValues...)
		} else {
			k = prefix + "_" + k

			keys = append(keys, k)
			values = append(values, v)
		}
	}
	return keys, values, nil
}
