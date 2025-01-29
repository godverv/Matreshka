package environment

import (
	"fmt"
)

func WithEnum(enums ...any) opt {
	return func(v *Variable) {
		if len(enums) == 0 {
			return
		}

		switch enums[0].(type) {
		case int:
			is := &intSliceValue{
				v: make([]int, 0, len(enums)),
			}
			for _, ev := range enums {
				is.v = append(is.v, ev.(int))
			}

			v.Enum.v = is
		case string:
			is := &stringSliceValue{
				v: make([]string, 0, len(enums)),
			}
			for _, ev := range enums {
				is.v = append(is.v, ev.(string))
			}

			v.Enum.v = is
		default:
			panic(fmt.Sprintf("unsupported enum type %T", enums[0]))
		}

	}
}

func WithType(tp variableType) opt {
	return func(v *Variable) {
		v.Type = tp
	}
}
