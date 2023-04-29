package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	var h, w int
	fmt.Scan(&h, &w)

	var rdr = bufio.NewReaderSize(os.Stdin, 10000)
	cArrayArray := make([][]string, h)

	resultMap := map[int]int{}

	for i := 0; i < h; i++ {
		cArray := make([]string, w)
		row := CCrossReadLine(rdr)
		for j, c := range row {
			cArray[j] = string([]rune{c})
		}
		cArrayArray[i] = cArray
	}

	var countFunc func(startH int, startW int)
	countFunc = func(startH int, startW int) {
		targetH := startH
		targetW := startW
		countResult := 1
		for {
			if targetH == h-1 || targetW == w-1 {
				break
			}
			targetH = targetH + 1
			targetW = targetW + 1
			if cArrayArray[targetH][targetW] == "#" {
				countResult = countResult + 1
			} else {
				break
			}
		}
		countResult = (countResult - 1) / 2
		if countResult > 0 {
			v, ok := resultMap[countResult]
			if ok {
				resultMap[countResult] = v + 1
			} else {
				resultMap[countResult] = 1
			}
		}

	}

	// 左上を見つけていく
	for i := 0; i < h; i++ {
		if i == 0 {
			for j := 0; j < w; j++ {
				if cArrayArray[i][j] == "#" {
					if j != w-1 {
						if cArrayArray[i+1][j+1] == "#" {
							countFunc(i, j)
						}
					}
				}
			}
		} else {
			if i != h-1 {
				for j := 0; j < w; j++ {
					if cArrayArray[i][j] == "#" {
						if j == 0 {
							if cArrayArray[i+1][j+1] == "#" {
								countFunc(i, j)
							}
						} else if j != w-1 {
							if cArrayArray[i+1][j+1] == "#" && cArrayArray[i-1][j-1] != "#" {
								countFunc(i, j)
							}
						}
					}
				}
			}

		}

	}

	loop := h
	if w < h {
		loop = w
	}

	resultStrArray := make([]string, loop)
	for i := 0; i < loop; i++ {
		v, ok := resultMap[i+1]
		if !ok {
			resultStrArray[i] = strconv.FormatInt(int64(0), 10)
		} else {
			resultStrArray[i] = strconv.FormatInt(int64(v), 10)
		}
	}
	fmt.Println(strings.Join(resultStrArray, " "))

}

func CCrossReadLine(rdr *bufio.Reader) string {
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
