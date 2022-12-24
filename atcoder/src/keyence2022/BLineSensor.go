package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func BLineSensorMain() {
	var h, w int
	fmt.Scan(&h, &w)

	var result = make([]int, w)

	var rdr = bufio.NewReaderSize(os.Stdin, 10000)

	for i := 0; i < h; i++ {
		cLine := BLineSensorRdr(rdr)
		for i2, c := range cLine {
			cMoji := string([]rune{c})
			if cMoji == "#" {
				result[i2] = result[i2] + 1
			}
		}

	}

	var resultSlice []string
	for i := 0; i < w; i++ {
		resultSlice = append(resultSlice, strconv.FormatInt(int64(result[i]), 10))
	}
	fmt.Println(strings.Join(resultSlice, " "))
}

func BLineSensorRdr(rdr *bufio.Reader) string {
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
