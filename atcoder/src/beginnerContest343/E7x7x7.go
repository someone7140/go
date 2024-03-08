package main

import (
	"fmt"
	"strconv"
	"strings"
)

type E7x7x7Result struct {
	v1 int
	v2 int
	v3 int
}

func E7x7x7Main() {
	var v1, v2, v3 int
	fmt.Scan(&v1, &v2, &v3)

	byResult := 0
	cxResult := 0
	findFlag := false

	var getTaiseki func(by int, cx int) E7x7x7Result
	getTaiseki = func(by int, cx int) E7x7x7Result {
		// b単独
		bTandoku := by * 7 * 7
		// c単独
		cTandoku := cx * 7 * 7
		// a単独
		aTandoku := by * cx * 7

		// aとcとbが被る領域
		bc := (7 - cx) * (7 - by) * 7
		// aとbが被る領域
		ab := (7-by)*7*7 - bc
		// aとcが被る領域
		ac := (7-cx)*7*7 - bc

		return E7x7x7Result{
			v1: aTandoku + bTandoku + cTandoku,
			v2: ab + ac,
			v3: bc,
		}
	}

	for by := 0; by <= 7; by++ {
		for cx := 0; cx <= 7; cx++ {
			temp := getTaiseki(by, cx)
			if v1 == temp.v1 && v2 == temp.v2 && v3 == temp.v3 {
				byResult = by
				cxResult = cx
				findFlag = true
				break
			}
		}
	}
	if findFlag {
		resultSlice := []string{"0", "0", "0", "0", strconv.FormatInt(int64(byResult), 10), "0", strconv.FormatInt(int64(cxResult), 10), "0", "0"}
		fmt.Println("Yes")
		fmt.Println(strings.Join(resultSlice, " "))
	} else {
		fmt.Println("No")
	}

}
