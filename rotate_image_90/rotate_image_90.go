package rotateimage

// This function works for rectangular image where
// height and width are not of equal dimension
func rotateimageby90(img [][]int) [][]int {
	if len(img) == 0 {
		return img
	}

	rotatedimg := make([][]int, len(img[0]))
	for i := 0; i < len(img); i++ {
		for j := 0; j < len(img[i]); j++ {
			rotatedimg[j] = append([]int{img[i][j]}, rotatedimg[j]...)
		}
	}
	return rotatedimg
}

func assert_square_matrix(arr [][]int) {
	is_square := len(arr) == 0 || len(arr) == len(arr[0])

	if !is_square {
		panic("not a square matrix")
	}
}

// This function is inplace for square image
// There no return type since original image will be
// modified
func rotatedimgby90inplace(img [][]int) {
	assert_square_matrix(img)
	start := 0
	end := len(img) - 1
	for start < end {
		for i := 0; i < end-start; i++ {
			// save the initial element
			start_element := img[start][start+i]
			img[start][start+i] = img[end-i][start]
			img[end-i][start] = img[end][end-i]
			img[end][end-i] = img[start+i][end]
			img[start+i][end] = start_element
		}
		start++
		end--
	}

}

// This function is inplace for square image
// There no return type since original image will be
// modified
func rotatedimgby90inplaceUsingTranspose(img [][]int) {
	assert_square_matrix(img)
	perform_inplace_matrix_transpose(img)
	swap_matrix_columns(img)
}

func swap_matrix_columns(img [][]int) {
	for i := 0; i < len(img); i++ {
		reverse_array(img[i])
	}
}

func reverse_array(arr []int) {
	i := 0
	j := len(arr) - 1
	for i < j {
		temp := arr[i]
		arr[i] = arr[j]
		arr[j] = temp
		i++
		j--
	}
}

//Only for square matrix
func perform_inplace_matrix_transpose(img [][]int) {
	assert_square_matrix(img)
	for i := 0; i < len(img); i++ {
		for j := i; j < len(img); j++ {
			temp := img[i][j]
			img[i][j] = img[j][i]
			img[j][i] = temp
		}
	}
}
