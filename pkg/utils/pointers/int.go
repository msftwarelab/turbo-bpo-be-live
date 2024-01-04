package pointers

func Int64(s int64) *int64 {
	return &s
}

func Int(s int) *int {
	return &s
}

func ToInt(s *int) int {
	if s != nil {
		return *s
	}
	return 0
}

func Int64ToInt(s *int64) *int {
	if s != nil {
		return Int(int(*s))

	}
	return Int(0)
}
