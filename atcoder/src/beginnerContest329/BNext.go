package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func BNextMain() {
	var n int
	fmt.Scan(&n)
	var rdr = bufio.NewReaderSize(os.Stdin, 10000)

	aStrs := strings.Split(BNextReadLine(rdr), " ")
	aArray := make([]int, n)

	for i, aMoji := range aStrs {
		a, _ := strconv.Atoi(aMoji)
		aArray[i] = a
	}
	sort.Sort(sort.Reverse(sort.IntSlice(aArray)))

	result := -1
	max := aArray[0]
	for i := 0; i < n; i++ {
		if aArray[i] < max {
			result = aArray[i]
			break
		}
	}

	fmt.Println(result)

}

func BNextReadLine(rdr *bufio.Reader) string {
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
