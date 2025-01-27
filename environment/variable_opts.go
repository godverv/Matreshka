package environment

func WithEnum(enum ...any) opt {
	return func(v *Variable) {
		v.Enum = enum
	}
}

func WithType(tp variableType) opt {
	return func(v *Variable) {
		v.Type = tp
	}
}
