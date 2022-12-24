package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func BAdjacencyListMain() {
	var n, m int
	fmt.Scan(&n, &m)

	var rdr = bufio.NewReaderSize(os.Stdin, 10000)

	var abMap = map[int][]int{}

	for i := 0; i < m; i++ {
		abArray := strings.Split(BAdjacencyListReadLine(rdr), " ")
		a, _ := strconv.Atoi(abArray[0])
		b, _ := strconv.Atoi(abArray[1])

		aArray, okA := abMap[a]
		if okA {
			abMap[a] = append(aArray, b)
		} else {
			abMap[a] = []int{b}
		}

		bArray, okB := abMap[b]
		if okB {
			abMap[b] = append(bArray, a)
		} else {
			abMap[b] = []int{a}
		}
	}

	var resultSlice []string
	for i := 1; i <= n; i++ {
		tempResArr, ok := abMap[i]
		if ok {
			sort.Ints(tempResArr)
			var tempResStr []string
			length := len(tempResArr)
			tempResStr = append(tempResStr, strconv.FormatInt(int64(length), 10))
			for j := 0; j < length; j++ {
				tempResStr = append(tempResStr, strconv.FormatInt(int64(tempResArr[j]), 10))
			}
			resultSlice = append(resultSlice, strings.Join(tempResStr, " "))
		} else {
			resultSlice = append(resultSlice, "0")
		}

	}
	fmt.Println(strings.Join(resultSlice, "\n"))

}

func BAdjacencyListReadLine(rdr *bufio.Reader) string {
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
