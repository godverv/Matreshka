package matreshka

import (
	"sort"
)

type ApplicationKeys struct {
	Servers     []string
	DataSources []string
	Environment []string
}

func GenerateKeys(a AppConfig) ApplicationKeys {
	keys := ApplicationKeys{
		Servers:     make([]string, 0, len(a.Servers)),
		DataSources: make([]string, 0, len(a.DataSources)),
		Environment: make([]string, 0, len(a.Environment)),
	}

	for _, s := range a.Servers {
		keys.Servers = append(keys.Servers, s.GetName())
	}
	for _, d := range a.DataSources {
		keys.DataSources = append(keys.DataSources, d.GetName())
	}
	name := a.Name
	if name != "" {
		name += "_"
	}
	//for k := range a.Environment {
	//	keys.Environment = append(keys.Environment, k)
	//}

	sort.Strings(keys.Environment)

	return keys
}
