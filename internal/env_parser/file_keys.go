package env_parser

import (
	"gopkg.in/yaml.v3"
)

type cfgKeysBuilder map[string]interface{}

func StructToEnvNames(prefix string, src interface{}) ([]string, error) {
	cfgBytes, err := yaml.Marshal(src)
	if err != nil {
		return nil, err
	}

	cfg := make(cfgKeysBuilder)
	err = yaml.Unmarshal(cfgBytes, cfg)
	if err != nil {
		return nil, err
	}

	return cfg.extractVariables(prefix, cfg)
}

func (c cfgKeysBuilder) extractVariables(prefix string, m map[string]interface{}) (out []string, err error) {
	for k, v := range m {
		if newMap, ok := v.(cfgKeysBuilder); ok {
			values, err := c.extractVariables(prefix+"_"+k, newMap)
			if err != nil {
				return nil, err
			}
			out = append(out, values...)
		} else {
			k = prefix + "_" + k

			out = append(out, k[1:])
		}
	}
	return out, nil
}
