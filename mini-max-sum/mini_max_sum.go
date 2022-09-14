package minimaxsum

import (
	"fmt"
	"sort"
)

func MiniMaxSum(s []int) {
	fmt.Println("MiniMaxSum")
	// sort.Ints(s)
	sort.Ints(s)
	// fmt.Println(s)
	var maxSum int
	var minSum int
	var sLen = len(s)
	// fmt.Println(sLen)
	for _, v := range s[:sLen-1] {
		// fmt.Println(v)
		minSum += v
	}
	fmt.Println(minSum)

	for _, v := range s[1:] {
		// fmt.Println(v)
		maxSum += v
	}
	fmt.Println(maxSum)

}
