package bit

type sumcarry struct {
	s rune
	c rune
}

var sumCarryMap = map[int]sumcarry{
	0: {'0', '0'},
	1: {'1', '0'},
	2: {'0', '1'},
	3: {'1', '1'},
}

func Add(s1 string, s2 string) string {
	if len(s1) == 0 {
		return s2
	}
	if len(s2) == 0 {
		return s1
	}

	diff := len(s1) - len(s2)
	if diff > 0 {
		rr, c := add(s1[diff:], s2)
		rl := increment(s1[0:diff], c)
		return rl + rr
	} else {
		rr, c := add(s2[-diff:], s1)
		rl := increment(s2[0:-diff], c)
		return rl + rr
	}
}

func increment(s string, by rune) string {
	if by == '0' {
		return s
	}
	c := by
	t := '0'
	r := ""
	for i := len(s) - 1; i >= 0; i-- {
		t, c = getSumAndCarry([]rune{rune(s[i]), '0', c})
		r = string(t) + r
	}
	if c == '1' {
		r = "1" + r
	}
	return r
}

func add(s1 string, s2 string) (string, rune) {
	c := '0'
	t := '0'
	r := ""
	for i := len(s1) - 1; i >= 0; i-- {
		t, c = getSumAndCarry([]rune{rune(s1[i]), rune(s2[i]), c})
		r = string(t) + r
	}
	return r, c
}

func getSumAndCarry(arr []rune) (rune, rune) {
	oneCount := 0
	for i := 0; i < len(arr); i++ {
		if arr[i] == '1' {
			oneCount++
		}
	}
	r := sumCarryMap[oneCount]
	return r.s, r.c
}
