package matreshka

import (
	"bytes"
	"reflect"
	"sort"
	"strings"

	"github.com/Red-Sock/evon"
	errors "github.com/Red-Sock/trace-errors"

	"github.com/godverv/matreshka/environment"
	"github.com/godverv/matreshka/internal/cases"
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

		if len(v.Enum) != 0 {
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

func (a *Environment) GenerateCustomGoStruct() []byte {
	structBuffer := bytes.NewBuffer(nil)
	imports := make(map[string]struct{})

	structBuffer.WriteString("type EnvironmentConfig struct {\n")
	for _, env := range *a {
		structBuffer.WriteByte('\t')
		name := strings.ReplaceAll(env.Name, " ", "_")
		structBuffer.WriteString(cases.SnakeToPascal(name))
		structBuffer.WriteByte(' ')
		typeName, importName := environment.MapVariableToGoType(*env)
		structBuffer.WriteString(typeName)
		structBuffer.WriteByte('\n')

		if importName != "" {
			imports[importName] = struct{}{}
		}
	}

	structBuffer.WriteByte('}')

	fileBuffer := bytes.NewBuffer(nil)
	fileBuffer.WriteString("package config\n\n")

	if len(imports) != 0 {
		fileBuffer.WriteString("import (\n")
		for importName := range imports {
			fileBuffer.WriteByte('\t')
			fileBuffer.WriteByte('"')
			fileBuffer.WriteString(importName)
			fileBuffer.WriteByte('"')
			fileBuffer.WriteString("\n")
		}
		fileBuffer.WriteString(")\n\n")
	}

	fileBuffer.Write(structBuffer.Bytes())
	return fileBuffer.Bytes()
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
		v := dstMapping[name]

		v.Set(reflect.ValueOf(env.Value))
	}

	return nil
}
