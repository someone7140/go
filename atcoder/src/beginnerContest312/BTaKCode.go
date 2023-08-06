package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func BTaKCodeMain() {
	var n, m int
	fmt.Scan(&n, &m)

	var rdr = bufio.NewReaderSize(os.Stdin, 10000)

	var sSliceSlice [][]string
	var resultSlice []string
	for i := 0; i < n; i++ {
		var sSlice []string
		s := BTaKCodeReadLine(rdr)
		for _, c := range s {
			sMoji := string([]rune{c})
			sSlice = append(sSlice, sMoji)

		}
		sSliceSlice = append(sSliceSlice, sSlice)
	}

	for i := 0; i < n; i++ {
		if i > n-9 {
			break
		}
		for j := 0; j < m; j++ {
			if j > m-9 {
				break
			}
			if sSliceSlice[i][j] == "#" {
				if sSliceSlice[i][j+1] == "#" &&
					sSliceSlice[i][j+2] == "#" &&
					sSliceSlice[i+1][j] == "#" &&
					sSliceSlice[i+1][j+1] == "#" &&
					sSliceSlice[i+1][j+2] == "#" &&
					sSliceSlice[i+2][j] == "#" &&
					sSliceSlice[i+2][j+1] == "#" &&
					sSliceSlice[i+2][j+2] == "#" &&
					sSliceSlice[i][j+3] == "." &&
					sSliceSlice[i+1][j+3] == "." &&
					sSliceSlice[i+2][j+3] == "." &&
					sSliceSlice[i+3][j] == "." &&
					sSliceSlice[i+3][j+1] == "." &&
					sSliceSlice[i+3][j+2] == "." &&
					sSliceSlice[i+3][j+3] == "." {
					i2 := i + 5
					j2 := j + 5
					if sSliceSlice[i2][j2] == "." &&
						sSliceSlice[i2][j2+1] == "." &&
						sSliceSlice[i2][j2+2] == "." &&
						sSliceSlice[i2][j2+3] == "." &&
						sSliceSlice[i2+1][j2] == "." &&
						sSliceSlice[i2+2][j2] == "." &&
						sSliceSlice[i2+3][j2] == "." &&
						sSliceSlice[i2+1][j2+1] == "#" &&
						sSliceSlice[i2+1][j2+2] == "#" &&
						sSliceSlice[i2+1][j2+3] == "#" &&
						sSliceSlice[i2+2][j2+1] == "#" &&
						sSliceSlice[i2+2][j2+2] == "#" &&
						sSliceSlice[i2+2][j2+3] == "#" &&
						sSliceSlice[i2+3][j2+1] == "#" &&
						sSliceSlice[i2+3][j2+2] == "#" &&
						sSliceSlice[i2+3][j2+3] == "#" {
						resultSlice = append(resultSlice, strconv.FormatInt(int64(i+1), 10)+" "+strconv.FormatInt(int64(j+1), 10))
					}

				}
			}
		}
	}

	fmt.Println(strings.Join(resultSlice, "\n"))

}

func BTaKCodeReadLine(rdr *bufio.Reader) string {
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
