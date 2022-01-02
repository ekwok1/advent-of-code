package grid

func NewIntGrid(height, width int) (grid [][]int) {
	for h := 0; h < height; h++ {
		gridRow := make([]int, width)
		grid = append(grid, gridRow)
	}

	return
}

func GridSum(grid *[][]int) (sum int) {
	gridRef := (*grid)
	for row := 0; row < len(gridRef); row++ {
		for col := 0; col < len(gridRef[0]); col++ {
			sum += gridRef[row][col]
		}
	}

	return
}
