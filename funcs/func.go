package funcs

func Sum(a []int) int {
	res := 0
	for _, elem := range a {
		res += elem
	}
	return res
}

func AvgSum(a []int) float64 {
	res := 0
	for _, elem := range a {
		res += elem
	}
	return float64(res / len(a))
}

func Min(a []int) int {
	res := a[0]
	for _, elem := range a {
		if elem < res {
			res = elem
		}
	}
	return res
}

func Max(a []int) int {
	res := a[0]
	for _, elem := range a {
		if elem > res {
			res = elem
		}
	}
	return res
}
