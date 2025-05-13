package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func CSumofProductMain() {
	var n int
	fmt.Scan(&n)
	var rdr = bufio.NewReaderSize(os.Stdin, 10000)
	aStrList := strings.Split(CSumofProductRdr(rdr), " ")
	aList := make([]int64, n)
	aSumList := make([]int64, n)

	var result int64 = 0
	for i, aStr := range aStrList {
		aInt, _ := strconv.Atoi(aStr)
		a := int64(aInt)
		aList[i] = a

	}

	for i := n - 1; i >= 0; i-- {
		if i == n-1 {
			aSumList[i] = aList[i]
		} else {
			aSumList[i] = aSumList[i+1] + aList[i]
		}
	}
	for i, a := range aList {
		if i < n-1 {
			result = result + a*aSumList[i+1]
		}
	}
	fmt.Println(result)
}

func CSumofProductRdr(rdr *bufio.Reader) string {
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
