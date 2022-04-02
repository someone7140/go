package main

import (
	"fmt"
	"math"
)

func D2variableFunctionMain() {
	var n int64
	fmt.Scan(&n)

	var result int64
	result = 0

	// まずはbが0の時を考える
	rippouN := int64(math.Cbrt(float64(n)))
	tempA := rippouN
	for {
		tempX := int64(math.Pow(float64(tempA), 3))
		if n <= tempX {
			result = tempX
			break
		} else {
			tempA = tempA + 1
		}

	}
	tempA = tempA - 1
	var tempB int64
	tempB = 1

	for {
		if tempA < tempB {
			break
		} else {
			temp := tempA*tempA*tempA + tempA*tempA*tempB + tempB*tempB*tempA + tempB*tempB*tempB
			if temp >= result {
				tempA = tempA - 1
			} else if temp < n {
				tempB = tempB + 1
			} else {
				if temp < result {
					result = temp
					tempA = tempA - 1
				}
			}
		}
	}
	fmt.Println(result)
}
