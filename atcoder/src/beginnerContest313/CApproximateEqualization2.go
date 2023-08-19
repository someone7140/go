package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func CApproximateEqualization2Main() {
	var n int
	fmt.Scan(&n)
	var rdr = bufio.NewReaderSize(os.Stdin, 10000)
	aListStr := strings.Split(CApproximateEqualization2ReadLine(rdr), " ")
	aArray := make([]int, n)

	var sum int64
	for i := 0; i < n; i++ {
		a, _ := strconv.Atoi(aListStr[i])
		aArray[i] = a
		sum = sum + int64(a)
	}
	sort.Ints(aArray)

	heikin := sum / int64(n)
	sabunArray := make([]int, n)
	heikinArray := make([]int, n)
	var heikinSum int64
	for i := 0; i < n; i++ {
		heikinArray[i] = int(heikin)
		heikinSum = heikinSum + heikin
	}
	var minusSum int64
	var plusSum int64
	for i := 0; i < int(sum-heikinSum); i++ {
		heikinArray[n-i-1] = heikinArray[n-i-1] + 1
	}

	for i := 0; i < n; i++ {
		sabunArray[i] = heikinArray[i] - aArray[i]
	}

	for i := 0; i < n; i++ {
		if sabunArray[i] < 0 {
			minusSum = minusSum - int64(sabunArray[i])
		}
		if sabunArray[i] > 0 {
			plusSum = plusSum + int64(sabunArray[i])
		}
	}
	if minusSum < plusSum {
		fmt.Println(minusSum)
	} else {
		fmt.Println(plusSum)
	}
}

func CApproximateEqualization2ReadLine(rdr *bufio.Reader) string {
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
