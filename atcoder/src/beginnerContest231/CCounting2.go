package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func CCounting2Main() {
	var n, q int
	fmt.Scan(&n, &q)

	var rdr = bufio.NewReaderSize(os.Stdin, 10000)

	aStrArray := strings.Split(readLineCCounting2(rdr), " ")
	aArray := make([]int, n)

	for i := 0; i < n; i++ {
		aArray[i], _ = strconv.Atoi(aStrArray[i])
	}

	sort.Slice(aArray, func(i, j int) bool { return aArray[i] < aArray[j] })

	results := make([]int, q)

	for i := 0; i < q; i++ {
		xStr := readLineCCounting2(rdr)
		x, _ := strconv.Atoi(xStr)
		results[i] = nibun(aArray, n, x)
	}

	for i := 0; i < q; i++ {
		fmt.Println(results[i])
	}
}

func readLineCCounting2(rdr *bufio.Reader) string {
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

func nibun(aArrayInput []int, n int, x int) int {

	if x <= aArrayInput[0] {
		return n
	} else if x > aArrayInput[n-1] {
		return 0
	} else {
		start := 0
		end := n - 1
		half := (start + end) / 2
		for {
			if start >= half || end <= half {
				break
			}
			if aArrayInput[half] < x {
				start = half
				half = (start + end) / 2
			} else if aArrayInput[half] >= x {
				end = half
				half = (start + end) / 2
			}
		}
		return n - end
	}
}
