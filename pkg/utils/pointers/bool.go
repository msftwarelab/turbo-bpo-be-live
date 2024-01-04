package pointers

func Bool(b bool) *bool {
	return &b
}

func ObjectTOBool(b *bool) bool {
	if b != nil {
		return *b
	}
	return false
}

func Float64(b float64) *float64 {
	return &b
}

func Float32(b float32) *float32 {
	return &b
}
