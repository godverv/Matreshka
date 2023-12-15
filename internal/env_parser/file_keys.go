package env_parser

type EnvVal struct {
	Name  string
	Value interface{}
}

func ExtractVariables(prefix string, m map[string]interface{}) (vals []EnvVal) {
	for k, v := range m {
		if newMap, ok := v.(map[string]interface{}); ok {
			vals = append(vals, ExtractVariables(prefix+"_"+k, newMap)...)
		} else {
			if prefix != "" {
				k = prefix + "_" + k
			}

			vals = append(vals, EnvVal{
				Name:  k,
				Value: v,
			})
		}
	}

	return vals
}
