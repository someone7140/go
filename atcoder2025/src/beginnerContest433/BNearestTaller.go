package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func BNearestTallerMain() {
	var n int
	fmt.Scan(&n)

	var rdr = bufio.NewReaderSize(os.Stdin, 10000)
	aStrArray := strings.Split(BNearestTallerRdr(rdr), " ")
	aArray := make([]int, n)
	for i := 0; i < n; i++ {
		a, _ := strconv.Atoi(aStrArray[i])
		aArray[i] = a
	}

	var resultSlice []string
	for i := 0; i < n; i++ {
		temp := -1
		for j := 0; j < i; j++ {
			if aArray[i] < aArray[j] {
				temp = j
			}
		}
		if temp == -1 {
			resultSlice = append(resultSlice, "-1")

		} else {
			resultSlice = append(resultSlice, strconv.FormatInt(int64(temp+1), 10))
		}
	}

	fmt.Println(strings.Join(resultSlice, "\n"))
}

func BNearestTallerRdr(rdr *bufio.Reader) string {
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
