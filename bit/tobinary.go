package bit

import (
	"bytes"
)

func assertLessThanOne(no float64) {
	if no >= 1 {
		panic("should be less than one")
	}
}

// The algorithm on computer always terminate as
// floating point is represented as exact representation
func ToBinary(no float64) string {
	assertLessThanOne(no)
	return toBinary(no)
}

func toBinary(no float64) string {
	var buffer bytes.Buffer
	buffer.WriteString("0.")
	for no > 0 {
		no *= 2
		if no >= 1 {
			buffer.WriteString("1")
			no -= 1
		} else {
			buffer.WriteString("0")
		}
	}
	return buffer.String()
}
