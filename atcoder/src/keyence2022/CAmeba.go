package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func CAmebaMain() {
	var n int
	fmt.Scan(&n)

	aMap := map[int]int{}
	aMap[1] = 0

	var rdr = bufio.NewReaderSize(os.Stdin, 10000)
	aStrArray := strings.Split(CAmebaRdr(rdr), " ")
	aArray := make([]int, n)

	for i, aStr := range aStrArray {
		a, _ := strconv.Atoi(aStr)
		aArray[i] = a
	}

	for i := 1; i <= n; i++ {
		a := aArray[i-1]
		a1 := 2 * i
		a2 := 2*i + 1
		oyaV, ok := aMap[a]
		if ok {
			next := oyaV + 1
			aMap[a1] = next
			aMap[a2] = next
		}
	}

	var resultSlice []string

	resultSlice = append(resultSlice, "0")
	for i := 1; i <= n; i++ {
		a1 := 2 * i
		a2 := 2*i + 1

		a1V, okA1 := aMap[a1]
		if okA1 {
			resultSlice = append(resultSlice, strconv.FormatInt(int64(a1V), 10))
		}

		a2V, okA2 := aMap[a2]
		if okA2 {
			resultSlice = append(resultSlice, strconv.FormatInt(int64(a2V), 10))
		}

	}

	fmt.Println(strings.Join(resultSlice, "\n"))
}

func CAmebaRdr(rdr *bufio.Reader) string {
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
