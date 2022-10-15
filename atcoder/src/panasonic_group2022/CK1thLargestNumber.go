package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func CK1thLargestNumberMain() {
	var n int
	fmt.Scan(&n)

	var rdr = bufio.NewReaderSize(os.Stdin, 10000)
	aStrArray := strings.Split(CK1thLargestNumberReadLine(rdr), " ")
	aSet := make(map[int]struct{})
	var aArray = make([]int, n)
	for i := 0; i < n; i++ {
		a, _ := strconv.Atoi(aStrArray[i])
		aSet[a] = struct{}{}
		aArray[i] = a
	}

	var aSetArray []int
	var lenA = 0
	for k := range aSet {
		lenA = lenA + 1
		aSetArray = append(aSetArray, k)
	}
	sort.Sort(sort.Reverse(sort.IntSlice(aSetArray)))
	aCountMap := make(map[int]int)
	for i, a := range aSetArray {
		aCountMap[a] = i
	}

	kosuuMap := make(map[int]int)
	for _, a := range aArray {
		v := aCountMap[a]
		kosuu, ok := kosuuMap[v]
		if ok {
			kosuuMap[v] = kosuu + 1
		} else {
			kosuuMap[v] = 1
		}
	}
	var resultSlice []string
	for i := 0; i < n; i++ {
		kosuu, ok := kosuuMap[i]
		if ok {
			resultSlice = append(resultSlice, strconv.FormatInt(int64(kosuu), 10))
		} else {
			resultSlice = append(resultSlice, "0")
		}
	}
	fmt.Println(strings.Join(resultSlice, "\n"))
}

func CK1thLargestNumberReadLine(rdr *bufio.Reader) string {
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
