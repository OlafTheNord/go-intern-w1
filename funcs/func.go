package funcs

import (
	"strconv"
)

func CalculateStats(a []int) (sum, avg, min, max float64) {
	min = float64(a[0])
	max = float64(a[0])
	for _, elem := range a {
		sum += float64(elem)
		if float64(elem) > max {
			max = float64(elem)
		}
		if float64(elem) < min {
			min = float64(elem)
		}
	}
	avg = sum / float64(len(a))
	return sum, avg, min, max
}

func ParseArgs(args []string) ([]int, error) {
	list := make([]int, 0, 0)
	for _, elem := range args[1:] {
		i, err := strconv.Atoi(elem)
		if err != nil {
			panic("1")
		}
		list = append(list, i)
	}
	if len(list) == 0 {
		panic("1")
	}
	return list, nil
}
