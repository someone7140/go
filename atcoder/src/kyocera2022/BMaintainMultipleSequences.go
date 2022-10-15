package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func BMaintainMultipleSequencesMain() {
	var n, q int
	fmt.Scan(&n, &q)

	var rdr = bufio.NewReaderSize(os.Stdin, 10000)
	var resultSliceStr []string
	var aSliceSlice [][]int

	for i := 0; i < n; i++ {
		var aSlice []int
		laStrArray := strings.Split(BMaintainMultipleSequencesRdr(rdr), " ")
		aLen, _ := strconv.Atoi(laStrArray[0])
		for j := 1; j <= aLen; j++ {
			a, _ := strconv.Atoi(laStrArray[j])
			aSlice = append(aSlice, a)
		}
		aSliceSlice = append(aSliceSlice, aSlice)
	}

	for i := 0; i < q; i++ {
		stStrArray := strings.Split(BMaintainMultipleSequencesRdr(rdr), " ")
		s, _ := strconv.Atoi(stStrArray[0])
		t, _ := strconv.Atoi(stStrArray[1])
		aSlice := aSliceSlice[s-1]
		a := aSlice[t-1]

		resultSliceStr = append(resultSliceStr, strconv.FormatInt(int64(a), 10))
	}

	fmt.Println(strings.Join(resultSliceStr, "\n"))
}

func BMaintainMultipleSequencesRdr(rdr *bufio.Reader) string {
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
