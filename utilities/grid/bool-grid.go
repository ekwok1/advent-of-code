package grid

func NewBoolGrid(height, width int) (grid [][]bool) {
	for h := 0; h < height; h++ {
		gridRow := make([]bool, width)
		grid = append(grid, gridRow)
	}

	return
}

func CountBool(grid *[][]bool, value bool) (count int) {
	gridRef := (*grid)
	for row := 0; row < len(gridRef); row++ {
		for col := 0; col < len(gridRef[0]); col++ {
			if gridRef[row][col] == value {
				count++
			}
		}
	}

	return
}
