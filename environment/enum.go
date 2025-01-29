package environment

type typedEnum interface {
	typedValue
	isEnum(value typedValue) error
}

type TypedEnum struct {
	v typedEnum
}

func (t *TypedEnum) MarshalYAML() (any, error) {
	return t.v.YamlValue(), nil
}

func (t *TypedEnum) Value() any {
	if t.v != nil {
		return t.v.Val()
	}

	return nil
}
