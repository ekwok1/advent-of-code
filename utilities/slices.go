package utilities

import "strings"

func CountTotal(slice *[]int) (total int) {
	for _, position := range *slice {
		total += position
	}
	return
}

func CreateBoolGrid(row int, col int) (grid [][]bool) {
	for i := 0; i < row; i++ {
		gridRow := make([]bool, col)
		grid = append(grid, gridRow)
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

func IndexOf(substr string, data *[]string) int {
	for i, e := range *data {
		if strings.Contains(e, substr) {
			return i
		}
	}

	return -1
}
