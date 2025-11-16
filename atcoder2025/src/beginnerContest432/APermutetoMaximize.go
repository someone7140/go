package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func APermutetoMaximizeMain() {
	var rdr = bufio.NewReaderSize(os.Stdin, 10000)
	abcStrArray := strings.Split(APermutetoMaximizeRdr(rdr), " ")
	abcArray := make([]int, len(abcStrArray))
	for i := 0; i < len(abcStrArray); i++ {
		a, _ := strconv.Atoi(abcStrArray[i])
		abcArray[i] = a
	}

	sort.Sort(sort.Reverse(sort.IntSlice(abcArray)))
	resultList := make([]string, len(abcStrArray))

	for i := 0; i < len(abcArray); i++ {
		resultList[i] = strconv.FormatInt(int64(abcArray[i]), 10)
	}
	fmt.Println(strings.Join(resultList, ""))
}

func APermutetoMaximizeRdr(rdr *bufio.Reader) string {
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
