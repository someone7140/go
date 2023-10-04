package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	var rdr = bufio.NewReaderSize(os.Stdin, 10000)
	pSliceSlice := make([][]string, 12)

	for i := 0; i < 12; i++ {
		pSliceSlice[i] = strings.Split(DPolyominoReadLine(rdr), "")
	}

	var poriomino1 [][]int
	var poriomino2 [][]int
	var poriomino3 [][]int
	poriominoMap := map[string]bool{}
	poriominoIndex := 1

	var setFunc func(row int, column int, key string)
	setFunc = func(row int, column int, key string) {
		poriominoMap[key] = true
		if poriominoIndex == 1 {
			poriomino1 = append(poriomino1, []int{row, column})
		}
		if poriominoIndex == 2 {
			poriomino2 = append(poriomino2, []int{row, column})

		}
		if poriominoIndex == 3 {
			poriomino3 = append(poriomino3, []int{row, column})
		}
		// 上に移動
		if row > 0 {
			newKey := strconv.FormatInt(int64(row-1), 10) + "-" + strconv.FormatInt(int64(column), 10)
			flag, ok := poriominoMap[newKey]
			if (!ok || !flag) && pSliceSlice[row-1][column] == "#" {
				setFunc(row-1, column, newKey)
			}
		}
		// 下に移動
		if row < 11 {
			newKey := strconv.FormatInt(int64(row+1), 10) + "-" + strconv.FormatInt(int64(column), 10)
			flag, ok := poriominoMap[newKey]
			if (!ok || !flag) && pSliceSlice[row+1][column] == "#" {
				setFunc(row+1, column, newKey)
			}
		}
		// 左に移動
		if column > 0 {
			newKey := strconv.FormatInt(int64(row), 10) + "-" + strconv.FormatInt(int64(column-1), 10)
			flag, ok := poriominoMap[newKey]
			if (!ok || !flag) && pSliceSlice[row][column-1] == "#" {
				setFunc(row, column-1, newKey)
			}
		}
		// 右に移動
		if column < 3 {
			newKey := strconv.FormatInt(int64(row), 10) + "-" + strconv.FormatInt(int64(column+1), 10)
			flag, ok := poriominoMap[newKey]
			if (!ok || !flag) && pSliceSlice[row][column+1] == "#" {
				setFunc(row, column+1, newKey)
			}
		}
	}

	for i := 0; i < 12; i++ {
		for j := 0; j < 4; j++ {
			p := pSliceSlice[i][j]
			if p == "#" && poriominoIndex < 4 {
				mapKey := strconv.FormatInt(int64(i), 10) + "-" + strconv.FormatInt(int64(j), 10)
				flag, ok := poriominoMap[mapKey]
				if !ok || !flag {
					setFunc(i, j, mapKey)
					poriominoIndex = poriominoIndex + 1
				}
			}

		}
	}

	fmt.Println("dummy")
}

func DPolyominoReadLine(rdr *bufio.Reader) string {
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
