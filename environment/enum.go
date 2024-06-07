package environment

func isValueInEnum(val any, enumSlice []any) bool {
	for _, v := range enumSlice {
		if v == val {
			return true
		}
	}
	return false
}
