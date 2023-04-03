package main

import (
	"fmt"
)

func CFourVariablesMain() {
	var n int
	fmt.Scan(&n)

	var result int64
	yakusuuMap := map[int]int64{}

	// 1からn-1までの数で約数の組み合わせ数をメモ
	for i := 1; i < n; i++ {
		var yakusuu int64
		for j := 1; j*j <= i; j++ {
			if i%j == 0 {
				yakusuu = yakusuu + 1
				if j*j != i {
					yakusuu = yakusuu + 1
				}
			}
		}
		yakusuuMap[i] = yakusuu
	}

	// abの数とcdの数で組み合わせ数を求める
	for i := 1; i < n; i++ {
		temp1 := yakusuuMap[i]
		temp2 := yakusuuMap[n-i]
		result = result + temp1*temp2
	}

	fmt.Println(result)

}
