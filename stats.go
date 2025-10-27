package main

import (
	"fmt"
	"go-intern-w1/funcs"
	"os"
)

func main() {
	//nums := make([]int, 0, 0)
	nums, err := funcs.ParseArgs(os.Args)
	if err != nil {
		fmt.Println(err)
	}
	a, b, c, d := funcs.CalculateStats(nums)

	fmt.Printf("Sum: %v\nAvg: %.2f\nMin: %v\nMax: %v", a, b, c, d)
}
