package matreshka

import (
	"sort"
	"strings"

	"github.com/Red-Sock/evon"

	"github.com/godverv/matreshka/environment"
)

type Environment []*environment.Variable

func (a *Environment) MarshalEnv(prefix string) []evon.Node {
	if prefix != "" {
		prefix += "_"
	}

	out := make([]evon.Node, 0, len(*a))
	for _, v := range *a {
		out = append(out, evon.Node{
			Name:  prefix + strings.ToUpper(v.Name),
			Value: v,
		})
	}

	sort.Slice(out, func(i, j int) bool {
		return out[i].Name > out[j].Name
	})

	return out
}
func (a *Environment) UnmarshalEnv(rootNode *evon.Node) error {
	//for _, n := range rootNode.InnerNodes {
	//a[strings.ToLower(n.Name[len(rootNode.Name)+1:])] = n.Value
	//}

	return nil
}
