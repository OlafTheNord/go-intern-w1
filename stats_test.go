package main

import (
	"fmt"
	"go-intern-w1/funcs"
	"testing"
)

func inputFunctionTest(t *testing.T) {
	var newList = make([]string, 0, 0)
	_, err := funcs.ParseArgs(newList)
	if err == nil {
		fmt.Errorf("waiting: panic, resul: %s", err)
	}

	var newList1 []int = []int{1, 1, 1}
	a, b, c, d := funcs.CalculateStats(newList1)
	if a != 1 && b != 1.00 && c != 1 && d != 1 {
		fmt.Errorf("waiting: 1, 1.00, 1, 1, got: %v, %.2f, %v, %v", a, b, c, d)
	}

	var newList2 []int = []int{0, 0, 0}
	a, b, c, d = funcs.CalculateStats(newList2)
	if err == nil {
		fmt.Errorf("waiting: err, got: %v, %.2f, %v, %v", a, b, c, d)
	}
}
