package main

import (
	"fmt"
	"math"
)

func CMinimizeAbs2Main() {
	var d int64
	fmt.Scan(&d)

	// 平方根を整数でとる
	sqrtInt := int64(math.Sqrt(float64(d)))
	// 差分値
	sabun1 := sqrtInt*sqrtInt - d
	sabun1Zettai := int64(math.Abs(float64(sabun1)))
	sabun2 := (sqrtInt+1)*(sqrtInt+1) - d
	sabun2Zettai := int64(math.Abs(float64(sabun2)))
	if sabun1Zettai <= 1 || sabun2Zettai == 0 {
		fmt.Println(0)
	} else {
		result := int64(0)
		if sabun1Zettai <= sabun2Zettai {
			result = sabun1Zettai
		} else {
			result = sabun2Zettai
		}

		tempSqrtInt := sqrtInt
		for {
			tempSabun := tempSqrtInt*tempSqrtInt - d
			tempSabunZettai := int64(math.Abs(float64(tempSabun)))
			//差分を平方根の近似値
			tempSabunZettaiSqrtInt := int64(math.Sqrt(float64(tempSabunZettai)))

			if tempSqrtInt < 2 {
				break
			} else {
				tempResult1 := (tempSqrtInt*tempSqrtInt + tempSabunZettaiSqrtInt*tempSabunZettaiSqrtInt) - d
				tempResult1Zettai := int64(math.Abs(float64(tempResult1)))
				tempResult2 := (tempSqrtInt*tempSqrtInt + (tempSabunZettaiSqrtInt+1)*(tempSabunZettaiSqrtInt+1)) - d
				tempResult2Zettai := int64(math.Abs(float64(tempResult2)))
				if tempResult1Zettai == 0 || tempResult2Zettai == 0 {
					result = 0
					break
				}
				if tempResult1Zettai < result {
					result = tempResult1Zettai
				}
				if tempResult2Zettai < result {
					result = tempResult2Zettai
				}
				tempSqrtInt = tempSqrtInt - 1
			}
		}

		fmt.Println(result)
	}
}
