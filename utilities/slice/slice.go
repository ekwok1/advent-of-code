package slice

import (
	"fmt"
	"math"
	"sort"
)

func Contains(slice *[]interface{}, value interface{}) bool {
	for _, item := range *slice {
		if item == value {
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
	sliceRef := *slice
	if len(sliceRef) < howMany {
		panic(fmt.Sprintf("You cannot return %d min ints from slice of length %d", howMany, len(sliceRef)))
	}

	sort.Ints(sliceRef)

	for len(ints) < howMany {
		min := sliceRef[0]
		sliceRef = sliceRef[1:]
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
