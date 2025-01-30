package environment

type typedEnum interface {
	typedValue
	isEnum(value typedValue) error
}

type TypedEnum struct {
	v typedEnum
}

func (t TypedEnum) MarshalYAML() (interface{}, error) {
	if t.v != nil {
		return t.v.YamlValue(), nil
	}

	return nil, nil
}
func (t TypedEnum) IsZero() bool {
	return t.v == nil
}

func (t TypedEnum) Value() any {
	if t.v != nil {
		return t.v.Val()
	}

	return nil
}
