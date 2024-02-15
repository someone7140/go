package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func CLeftoverRecipesMain() {
	var n int
	fmt.Scan(&n)
	var rdr = bufio.NewReaderSize(os.Stdin, 10000)

	// qを読み込む
	var qSlice = make([]int, n)
	qStrArray := strings.Split(CLeftoverRecipesReadLine(rdr), " ")
	for i := 0; i < n; i++ {
		qSlice[i], _ = strconv.Atoi(qStrArray[i])
	}

	// aを読みこむ
	var aSlice = make([]int, n)
	aStrArray := strings.Split(CLeftoverRecipesReadLine(rdr), " ")
	for i := 0; i < n; i++ {
		a, _ := strconv.Atoi(aStrArray[i])
		aSlice[i] = a
	}

	// bを読みこむ
	var bSlice = make([]int, n)
	bStrArray := strings.Split(CLeftoverRecipesReadLine(rdr), " ")
	for i := 0; i < n; i++ {
		b, _ := strconv.Atoi(bStrArray[i])
		bSlice[i] = b
	}
	result := -1
	for i := 0; i < n; i++ {
		// Aで何個作れるか
		aCount := 0
		if aSlice[i] > 0 {
			aCount = qSlice[i] / aSlice[i]
		}
		bCount := 0
		if bSlice[i] > 0 {
			bCount = qSlice[i] / bSlice[i]
		}

		if aCount < bCount {
			nokoriCount := (qSlice[i] - bSlice[i]*bCount) / aSlice[i]
			bCount = bCount + nokoriCount
		} else {
			nokoriCount := (qSlice[i] - aSlice[i]*aCount) / bSlice[i]
			aCount = aCount + nokoriCount
		}
		if bCount > aCount {
			if result == -1 || bCount < result {
				result = bCount
			}
		} else {
			if result == -1 || aCount < result {
				result = aCount
			}
		}
	}
	fmt.Println(result)

}

func CLeftoverRecipesReadLine(rdr *bufio.Reader) string {
	buf := make([]byte, 0, 10000)
	for {
		l, p, e := rdr.ReadLine()
		if e != nil {
			panic(e)
		}
		buf = append(buf, l...)
		if !p {
			break
		}
	}
	return string(buf)
}
