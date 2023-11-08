package env_parser

func ExtractVariables(prefix string, m map[string]interface{}) (keys []string, values []any, err error) {
	for k, v := range m {
		if newMap, ok := v.(map[string]interface{}); ok {
			upperKeys, upperValues, err := ExtractVariables(prefix+"_"+k, newMap)
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
