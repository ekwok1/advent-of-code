package slice

import (
	"fmt"
	"math"
	"sort"
)

func Contains(slice *[]interface{}, item interface{}) bool {
	for _, element := range *slice {
		if element == item {
			return true
		}
	}

	return false
}

func MinInt(slice *[]int) int {
	min := math.MaxInt

	for _, i := range *slice {
		if i < min {
			min = i
		}
	}

	return min
}

func MinInts(slice *[]int, howMany int) (ints []int) {
	if len(*slice) < howMany {
		panic(fmt.Sprintf("You cannot return %d min ints from slice of length %d", howMany, len(*slice)))
	}

	sort.Ints(*slice)

	for len(ints) < howMany {
		min := (*slice)[0]
		(*slice) = (*slice)[1:]
		ints = append(ints, min)
	}

	return
}

func SumInts(slice *[]int) (sum int) {
	for _, position := range *slice {
		sum += position
	}

	return
}
