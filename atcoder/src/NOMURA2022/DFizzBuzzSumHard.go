package main

import (
	"fmt"
)

func DFizzBuzzSumHardMain() {
	var n, a, b int64
	fmt.Scan(&n, &a, &b)

	result := n * (n + 1) / 2

	var max int64 = 0
	var min int64 = 0

	if a > b {
		max = a
		min = b
	} else {
		max = b
		min = a
	}

	if max%min == 0 {
		minus := DFizzBuzzSumHardTargetSum(n, min)
		fmt.Println(
			result - minus,
		)
	} else {
		maxMinus := DFizzBuzzSumHardTargetSum(n, max)
		minMinus := DFizzBuzzSumHardTargetSum(n, min)
		plus := DFizzBuzzSumHardTargetSum(n, DFizzBuzzSumHardTargeCal(max, min))
		fmt.Println(
			result - maxMinus - minMinus + plus,
		)
	}

}

func DFizzBuzzSumHardTargetSum(n int64, target int64) int64 {
	syou := n / target
	max := target * syou
	return syou * (target + max) / 2
}

// 最大公約数
func DFizzBuzzSumHardTargeGcd(a, b int64) int64 {
	if b == 0 {
		return a
	}
	return DFizzBuzzSumHardTargeGcd(b, a%b)
}

// 最小公約数
func DFizzBuzzSumHardTargeCal(a, b int64) int64 {
	return a * b / DFizzBuzzSumHardTargeGcd(a, b)
}
