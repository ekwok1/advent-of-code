package utilities

func CountTotal(slice *[]int) (total int) {
	for _, position := range *slice {
		total += position
	}
	return
}

func CreateIntGrid(size int) (ret [][]int) {
	for i := 0; i < size; i++ {
		gridRow := make([]int, size)
		ret = append(ret, gridRow)
	}

	return
}
