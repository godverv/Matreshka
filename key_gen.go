package matreshka

func GenerateKeys(a AppConfig) []string {
	keys := make([]string, 0, len(a.Servers)+len(a.DataSources)+len(a.Environment))

	for _, s := range a.Servers {
		keys = append(keys, s.GetName())
	}
	for _, d := range a.DataSources {
		keys = append(keys, d.GetName())
	}
	name := a.Name
	if name != "" {
		name += "_"
	}
	for k := range a.Environment {
		keys = append(keys, k)
	}

	return keys
}
