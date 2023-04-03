package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func CMergeSequencesMain() {
	var n, m int
	fmt.Scan(&n, &m)

	var rdr = bufio.NewReaderSize(os.Stdin, 10000)
	aStrArray := strings.Split(CMergeSequencesReadLine(rdr), " ")
	bStrArray := strings.Split(CMergeSequencesReadLine(rdr), " ")

	abArray := make([]int, n+m)
	aOrderMap := map[int]int{}
	bOrderMap := map[int]int{}
	for i := 0; i < n; i++ {
		a, _ := strconv.Atoi(aStrArray[i])
		abArray[i] = a
		aOrderMap[a] = -1
	}
	for i := 0; i < m; i++ {
		b, _ := strconv.Atoi(bStrArray[i])
		abArray[i+n] = b
		bOrderMap[b] = -1
	}
	sort.Ints(abArray)
	var resultSliceA []string
	var resultSliceB []string
	for index, val := range abArray {
		// aに存在
		_, ok := aOrderMap[val]
		if ok {
			resultSliceA = append(resultSliceA, strconv.FormatInt(int64(index+1), 10))
		} else {
			resultSliceB = append(resultSliceB, strconv.FormatInt(int64(index+1), 10))
		}
	}

	fmt.Println(strings.Join(resultSliceA, " "))
	fmt.Println(strings.Join(resultSliceB, " "))

}

func CMergeSequencesReadLine(rdr *bufio.Reader) string {
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
