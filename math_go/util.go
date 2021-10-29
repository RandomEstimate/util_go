package math_go

func nanFill(n int) []interface{} {
	s := make([]interface{}, 0, n)
	for i := 0; i < n; i++ {
		s = append(s, nil)
	}
	return s
}
