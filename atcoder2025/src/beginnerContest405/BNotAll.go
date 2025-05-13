package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func BNotAllMain() {
	var n, m int
	fmt.Scan(&n, &m)
	var rdr = bufio.NewReaderSize(os.Stdin, 10000)
	aStrList := strings.Split(BNotAllRdr(rdr), " ")
	aSet := make(map[int]struct{})

	result := 0
	endIndex := -1
	for i, aStr := range aStrList {
		a, _ := strconv.Atoi(aStr)
		aSet[a] = struct{}{}
		if len(aSet) == m {
			endIndex = i
			break
		}
	}

	if endIndex > -1 {
		result = n - endIndex
	}

	fmt.Println(result)
}

func BNotAllRdr(rdr *bufio.Reader) string {
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
