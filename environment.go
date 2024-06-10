package matreshka

import (
	"sort"
	"strings"

	"github.com/Red-Sock/evon"
	errors "github.com/Red-Sock/trace-errors"

	"github.com/godverv/matreshka/environment"
)

type Environment []*environment.Variable

func (a *Environment) MarshalEnv(prefix string) []evon.Node {
	if prefix != "" {
		prefix += "_"
	}

	out := make([]evon.Node, 0, len(*a))
	for _, v := range *a {
		pref := prefix + strings.NewReplacer(" ", "-", "_", "-").Replace(strings.ToUpper(v.Name))
		out = append(out,
			evon.Node{
				Name:  pref,
				Value: v.ValueString(),
			},
			evon.Node{
				Name:  pref + "_TYPE",
				Value: v.Type,
			},
		)

		if len(v.Enum) != 0 {
			out = append(out, evon.Node{
				Name:  pref + "_ENUM",
				Value: v.EnumString(),
			})
		}
	}

	sort.Slice(out, func(i, j int) bool {
		return out[i].Name > out[j].Name
	})

	return out
}
func (a *Environment) UnmarshalEnv(rootNode *evon.Node) error {
	env := make([]*environment.Variable, 0, len(rootNode.InnerNodes))

	for _, e := range rootNode.InnerNodes {
		ev := &environment.Variable{
			Name: strings.ToLower(e.Name[len(rootNode.Name)+1:]),
		}
		err := ev.UnmarshalEnv(e)
		if err != nil {
			return errors.Wrap(err, "error unmarshalling environment variable")
		}
		env = append(env, ev)
	}

	*a = env
	return nil
}
