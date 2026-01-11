package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	var n int
	fmt.Scan(&n)
	var rdr = bufio.NewReaderSize(os.Stdin, 10000)
	aStrs := strings.Split(DTailofSnakeRdr(rdr), " ")
	bStrs := strings.Split(DTailofSnakeRdr(rdr), " ")
	cStrs := strings.Split(DTailofSnakeRdr(rdr), " ")
	aArray := make([]int64, n)
	aSumArray := make([]int64, n)
	bArray := make([]int64, n)
	bSumArray := make([]int64, n)
	cArray := make([]int64, n)
	cSumArray := make([]int64, n)
	for i := 0; i < n; i++ {
		a, _ := strconv.Atoi(aStrs[i])
		aInt64 := int64(a)
		aArray[i] = aInt64
		b, _ := strconv.Atoi(bStrs[i])
		bInt64 := int64(b)
		bArray[i] = bInt64
		c, _ := strconv.Atoi(cStrs[i])
		cInt64 := int64(c)
		cArray[i] = cInt64
		if i == 0 {
			aSumArray[i] = aInt64
			bSumArray[i] = bInt64
			cSumArray[i] = cInt64
		} else {
			aSumArray[i] = aSumArray[i-1] + aInt64
			bSumArray[i] = bSumArray[i-1] + bInt64
			cSumArray[i] = cSumArray[i-1] + cInt64
		}
	}
	/*
		tempResult := aArray[0] + cArray[n-1] + bSumArray[n-2] - bArray[0]
		changeIdx1 := 1
		changeIdx2 := n - 2
		changeIdx1Fix := false
		changeIdx2Fix := false
		// aの判定
		for {
			if changeIdx1 >= changeIdx2 {
				break
			}

			changeIdx1 = changeIdx1 + 1
			if changeIdx1 >= changeIdx2 {
				break
			}

		}
		fmt.Println(result)
	*/
}

func DTailofSnakeRdr(rdr *bufio.Reader) string {
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
