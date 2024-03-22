package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func BDelimiterMain() {
	var rdr = bufio.NewReaderSize(os.Stdin, 10000)
	var aSlice []int
	for {
		a, _ := strconv.Atoi(BDelimiterReadLine(rdr))
		aSlice = append(aSlice, a)
		if a == 0 {
			break
		}
	}

	var resultSlice []string
	len := len(aSlice)
	for i := len - 1; i >= 0; i-- {
		resultSlice = append(resultSlice, strconv.FormatInt(int64(aSlice[i]), 10))
	}

	fmt.Println(strings.Join(resultSlice, "\n"))
}

func BDelimiterReadLine(rdr *bufio.Reader) string {
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
