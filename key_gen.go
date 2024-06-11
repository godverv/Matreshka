package matreshka

type ApplicationKeys struct {
	Servers     []string
	DataSources []string
}

func GenerateKeys(a AppConfig) ApplicationKeys {
	keys := ApplicationKeys{
		Servers:     make([]string, 0, len(a.Servers)),
		DataSources: make([]string, 0, len(a.DataSources)),
	}

	for _, s := range a.Servers {
		keys.Servers = append(keys.Servers, s.GetName())
	}
	for _, d := range a.DataSources {
		keys.DataSources = append(keys.DataSources, d.GetName())
	}

	return keys
}
