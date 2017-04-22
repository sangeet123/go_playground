package url

const zero = '0'
const two = '2'
const percentage = '%'

// Urlify replaces all the ocurrances of space in between
// characters with "%20". The url should have space for the
// "%20" at it back.
func Urlify(url []uint8, length int) {
	space := uint8(' ')
	i := length - 1
	for i >= 0 && url[i] == space {
		i--
	}
	j := length - 1
	for i = i; i >= 0; i-- {
		if url[i] == space {
			url[j] = zero
			url[j-1] = two
			url[j-2] = percentage
			j = j - 3
		} else {
			url[j] = url[i]
			j--
		}
	}

}
