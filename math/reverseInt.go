package math

func Reverse(no int) no {
	r := 0
	for no != 0 {
		r = r*10 + no%10
		no /= 10
	}
	return r
}
