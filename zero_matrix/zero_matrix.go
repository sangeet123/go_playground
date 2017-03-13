package zeromatrix

func zero_matrix(arr [][]int){
	if len(arr) == 0{
		return
	}

	set_rows := make([]int, len(arr))
	set_column := make([]int, len(arr[0]))

	initialize(set_rows, 1)
	initialize(set_column, 1)

	for i := 0 ; i < len(arr); i++{
		for j := 0 ; j < len(arr[i]) ; j++{
			if arr[i][j] ==0 {
			set_rows[i] = 0
			set_column[j] = 0
			}
		}
	}

	for i := 0 ; i < len(arr); i++{
		for j := 0 ; j < len(arr[i]) ; j++{
			arr[i][j] = set_rows[i]&set_column[j]
		}
	}
}

func initialize(arr []int, val int){
	for i:=0;i<len(arr);i++{
		arr[i] = val
	}
}

//[0 1 1] 
//[0 1 1] 
//[0 0 0]

//[0 1 1]
//[1 1 1] 
//[0 0 0]

// This is way too complicated
// Taking memory O(2n) is more good enough
func zero_matrix_inplace(arr [][]int){
	set_first_row := true
	set_first_column := true

	for i:=0;i< len(arr);i++ {
		if !set_first_row && !set_first_column{
			break
		}

		if arr[0][i] == 0{
			set_first_row = false
		}

		if arr[i][0] == 0{
			set_first_column = false
		}
	}

	for i:=0;i<len(arr);i++{
		for j:=0;j<len(arr);j++{
			if(arr[i][j] == 0){
				arr[0][j] = 0
				arr[i][0] = 0
			}
		}
	}

	for i:=1;i<len(arr);i++{
		for j:=1;j<len(arr[i]);j++{
			if(arr[i][0] == 0 || arr[0][j] == 0){
				arr[i][j] = 0
			}
		}
	}

	if !set_first_column{
		for i:=0;i<len(arr);i++{
			arr[i][0] = 0
		}
	}

	if !set_first_row{
		for i:=0;i<len(arr);i++{
			arr[0][i]=0
		}
	}


}