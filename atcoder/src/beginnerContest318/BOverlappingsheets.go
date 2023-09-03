package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func BOverlappingsheetsMain() {
	var n int
	fmt.Scan(&n)

	type SheetRange struct {
		xMin int
		xMax int
		yMin int
		yMax int
	}
	var rdr = bufio.NewReaderSize(os.Stdin, 10000)
	var sheetRangeSlice = make([]SheetRange, n)

	for i := 0; i < n; i++ {
		abcd := strings.Split(BOverlappingsheetsReadLine(rdr), " ")
		xMin, _ := strconv.Atoi(abcd[0])
		xMax, _ := strconv.Atoi(abcd[1])
		yMin, _ := strconv.Atoi(abcd[2])
		yMax, _ := strconv.Atoi(abcd[3])
		sheetRangeSlice[i] = SheetRange{
			xMin: xMin,
			xMax: xMax,
			yMin: yMin,
			yMax: yMax,
		}
	}
	result := 0
	for i := 0; i < 100; i++ {
		xMin := i
		xMax := i + 1
		for j := 0; j < 100; j++ {
			yMin := j
			yMax := j + 1
			for _, sheetRange := range sheetRangeSlice {
				if xMax <= sheetRange.xMax && xMin >= sheetRange.xMin && yMax <= sheetRange.yMax && yMin >= sheetRange.yMin {
					result = result + 1
					break
				}
			}

		}
	}
	fmt.Println(result)
}

func BOverlappingsheetsReadLine(rdr *bufio.Reader) string {
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
