package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var h, w int
	fmt.Scan(&h, &w)

	var rdr = bufio.NewReaderSize(os.Stdin, 10000)
	rowMaps := make([]map[string][]int, w)
	columnMaps := make([]map[string][]int, h)

	deleteRowIndexs := make(map[int]struct{})
	deleteColumnIndexs := make(map[int]struct{})

	for i := 0; i < h; i++ {
		cStr := DMagicalCookiesReadLine(rdr)
		rowMap := rowMaps[i]
		for j, c := range cStr {
			cMoji := string([]rune{c})
			indexList, ok := rowMap[cMoji]
			if ok {
				rowMap[cMoji] = append(indexList, j)
			} else {
				rowMap[cMoji] = []int{j}
			}

			columnMap := columnMaps[j]
			indexListC, ok2 := columnMap[cMoji]
			if ok2 {
				columnMap[cMoji] = append(indexListC, i)
			} else {
				columnMap[cMoji] = []int{i}
			}
		}

	}

	loopCount := 0
	for {
		deleteRowCount := 0
		deleteColumnCount := 0

		if deleteRowCount == 0 && deleteColumnCount == 0 {
			break
		}
		loopCount = loopCount + 1
	}

	total := h * w
	deleteRowsCount := len(deleteRowIndexs) * w
	deleteColumnsCount := len(deleteColumnIndexs) * (h - deleteRowsCount)
	fmt.Println(total - deleteRowsCount - deleteColumnsCount)
}

func DMagicalCookiesReadLine(rdr *bufio.Reader) string {
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
