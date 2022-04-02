package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func CCouponMain() {

	var n, k, x int
	fmt.Scan(&n, &k, &x)

	var rdr = bufio.NewReaderSize(os.Stdin, 10000)
	aStrArray := strings.Split(readLine(rdr), " ")
	aArray := make([]int, n)

	aArrayAfterKoupon := make([]int, n)

	for i := 0; i < n; i++ {
		aArray[i], _ = strconv.Atoi(aStrArray[i])
	}
	sort.Slice(aArray, func(i, j int) bool { return aArray[j] < aArray[i] })

	for i := 0; i < n; i++ {
		a := aArray[i]
		if k == 0 || a < x {
			aArrayAfterKoupon[i] = aArray[i]
		} else {
			syou := a / x
			if k < syou {
				aArrayAfterKoupon[i] = aArray[i] - k*x
				k = 0
			} else {
				aArrayAfterKoupon[i] = aArray[i] - syou*x
				k = k - syou
			}
		}
	}

	sort.Slice(aArrayAfterKoupon, func(i, j int) bool { return aArrayAfterKoupon[j] < aArrayAfterKoupon[i] })
	result := 0

	for i := 0; i < n; i++ {
		a := aArrayAfterKoupon[i]
		if k == 0 {
			result = result + a
		} else {
			k = k - 1
		}
	}
	fmt.Println(result)
}

func readLine(rdr *bufio.Reader) string {
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
