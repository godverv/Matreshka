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

	replacer := strings.NewReplacer("-", " ")

	for _, e := range rootNode.InnerNodes {
		name := e.Name[len(rootNode.Name)+1:]
		name = strings.ToLower(name)
		name = replacer.Replace(name)
		ev := &environment.Variable{
			Name: name,
		}
		err := ev.UnmarshalEnv(e)
		if err != nil {
			return errors.Wrap(err, "error unmarshalling environment variable")
		}
		env = append(env, ev)
	}

	sort.Slice(env, func(i, j int) bool {
		return env[i].Name < env[j].Name
	})

	*a = env
	return nil
}
