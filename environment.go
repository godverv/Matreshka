package matreshka

import (
	"reflect"
	"sort"
	"strings"

	"go.redsock.ru/evon"
	errors "go.redsock.ru/rerrors"

	"go.verv.tech/matreshka/environment"
	"go.verv.tech/matreshka/internal/cases"
)

var ErrNotAPointer = errors.New("not a pointer")

type Environment []*environment.Variable

func (a *Environment) MarshalEnv(prefix string) ([]*evon.Node, error) {
	if prefix != "" {
		prefix += "_"
	}

	out := make([]*evon.Node, 0, len(*a))
	for _, v := range *a {
		pref := prefix + strings.NewReplacer(" ", "-", "_", "-").Replace(strings.ToUpper(v.Name))
		root := &evon.Node{
			Name:  pref,
			Value: v.ValueString(),
			InnerNodes: []*evon.Node{{
				Name:  pref + "_TYPE",
				Value: v.Type,
			},
			},
		}

		if v.Enum != nil {
			root.InnerNodes = append(root.InnerNodes,
				&evon.Node{
					Name:  pref + "_ENUM",
					Value: v.EnumString(),
				})
		}
		out = append(out, root)
	}

	sort.Slice(out, func(i, j int) bool {
		return out[i].Name < out[j].Name
	})

	return out, nil
}
func (a *Environment) UnmarshalEnv(rootNode *evon.Node) error {
	env := make([]*environment.Variable, 0, len(rootNode.InnerNodes))

	replacer := strings.NewReplacer("-", "_")

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

func (a *Environment) ParseToStruct(dst any) error {
	dstRef := reflect.ValueOf(dst)
	if dstRef.Kind() != reflect.Ptr {
		return errors.Wrap(ErrNotAPointer, "expected destination to be a pointer ")
	}

	dstRef = dstRef.Elem()
	numFields := dstRef.NumField()

	dstMapping := make(map[string]reflect.Value)

	for i := 0; i < numFields; i++ {
		field := dstRef.Type().Field(i)
		dstMapping[field.Name] = dstRef.Field(i)
	}

	for _, env := range *a {
		name := env.Name
		name = strings.ReplaceAll(name, " ", "_")
		name = cases.SnakeToPascal(name)
		v, ok := dstMapping[name]
		if !ok {
			return errors.Wrap(ErrNotFound, "field with name "+name+" can't be found in target struct")
		}

		v.Set(reflect.ValueOf(env.Value.Value()))

	}

	return nil
}
