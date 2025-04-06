package main

import (
	"fmt"
	"math"
)

func C2ab2Main() {
	var n int64
	fmt.Scan(&n)

	result := int64(0)
	temp2 := int64(2)

	for {
		if temp2 > n {
			break
		}

		// まずはnを割る
		temp_waru := n / temp2
		// 平方根を切り捨て
		temp_max := int64(math.Sqrt(float64(temp_waru)))
		// すでに大きい場合はマイナス
		temp_total := temp2 * temp_max * temp_max
		if temp_total > n {
			for {
				if temp_total <= n {
					break
				}
				temp_max = temp_max - 1
				temp_total = temp2 * temp_max * temp_max
			}
			// それ以外は足していく
		} else if temp_total < n {
			temp_max2 := temp_max
			for {
				temp_max2 = temp_max2 + 1
				temp_total2 := temp2 * temp_max2 * temp_max2
				if temp_total2 > n {
					break
				}
				temp_max = temp_max + 1
				temp_total = temp2 * temp_max * temp_max
			}
		}
		// あまりがある場合は2で割った結果に1足す
		if temp_max%2 == 0 {
			result = result + temp_max/2
		} else {
			result = result + temp_max/2 + 1
		}

		temp2 = temp2 * 2
	}

	fmt.Println(result)
}
