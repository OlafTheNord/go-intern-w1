package main

import (
	"fmt"
	"go-intern-w1/funcs"
	"os"
	"strconv"
)

func main() {
	list := make([]int, 0, 0)
	for _, elem := range os.Args[1:] {
		i, err := strconv.Atoi(elem)
		if err != nil {
			panic("1")
		}
		list = append(list, i)
	}
	if len(list) == 0 {
		panic("1")
	}

	fmt.Println("Sum:", funcs.Sum(list))
	fmt.Println("Avg:", funcs.AvgSum(list))
	fmt.Println("Min:", funcs.Min(list))
	fmt.Println("Max:", funcs.Max(list))
}
