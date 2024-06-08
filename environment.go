package matreshka

import (
	"fmt"
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
		pref := prefix + strings.ReplaceAll(strings.ToUpper(v.Name), " ", "_")
		out = append(out,
			evon.Node{
				Name:  pref,
				Value: fmt.Sprint(v.Value),
			},
			evon.Node{
				Name:  pref + "_TYPE",
				Value: v.Type,
			},
		)

		if len(v.Enum) != 0 {
			out = append(out, evon.Node{
				Name:  pref + "_ENUM",
				Value: fmt.Sprint(v.Enum),
			})
		}
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
