package main

import (
	"fmt"

	minimaxsum "github.com/enzonun/hackerrank-go/mini-max-sum"
	"github.com/enzonun/hackerrank-go/plusminus"
)

func main() {
	var positive float64
	positive += 1
	fmt.Println(plusminus.RoundDown(positive/6, 6))
	minimaxsum.MiniMaxSum([]int{1, 2, 3, 4, 5})
}
