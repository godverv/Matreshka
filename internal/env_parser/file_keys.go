package env_parser

type EnvVal struct {
	Name  string
	Value interface{}
}

func ExtractVariables(prefix string, m map[string]interface{}) (vals []EnvVal, err error) {
	for k, v := range m {
		if newMap, ok := v.(map[string]interface{}); ok {
			vs, err := ExtractVariables(prefix+"_"+k, newMap)
			if err != nil {
				return nil, err
			}

			vals = append(vals, vs...)
		} else {
			k = prefix + "_" + k

			vals = append(vals, EnvVal{
				Name:  k,
				Value: v,
			})
		}
	}

	return vals, nil
}
