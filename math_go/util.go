package math_go

func Full(n int, x interface{}) Series {
	s := make(Series, 0, n)
	for i := 0; i < n; i++ {
		s = append(s, x)
	}
	return s
}
