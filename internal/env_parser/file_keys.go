package env_parser

type EnvVal struct {
	Name  string
	Value interface{}
}

func ExtractVariables(prefix string, m map[string]interface{}) (vals []EnvVal) {
	for k, v := range m {
		if prefix != "" {
			k = prefix + "_" + k
		}
		if newMap, ok := v.(map[string]interface{}); ok {
			vals = append(vals, ExtractVariables(k, newMap)...)
		} else {
			vals = append(vals, EnvVal{
				Name:  k,
				Value: v,
			})
		}
	}

	return vals
}
