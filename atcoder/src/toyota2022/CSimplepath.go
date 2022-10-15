package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var n, x, y int
var resultSliceStrSlice [][]string
var resultMinLen int
var xyMap map[int][]int

func main() {
	fmt.Scan(&n, &x, &y)

	var rdr = bufio.NewReaderSize(os.Stdin, 10000)
	resultSliceStrSlice = [][]string{}
	xyMap = map[int][]int{}

	resultMinLen = -1

	for i := 0; i < n-1; i++ {
		xyStrArray := strings.Split(CSimplepathRdr(rdr), " ")
		tempX, _ := strconv.Atoi(xyStrArray[0])
		tempY, _ := strconv.Atoi(xyStrArray[1])

		xArray, okX := xyMap[tempX]
		if okX {
			xyMap[tempX] = append(xArray, tempY)
		} else {
			xyMap[tempX] = []int{tempY}
		}
		yArray, okY := xyMap[tempY]
		if okY {
			xyMap[tempY] = append(yArray, tempX)
		} else {
			xyMap[tempY] = []int{tempX}
		}

	}

	visitMap := map[int]bool{}
	visitMap[x] = true
	LoopCSimplepathRdr(x, []string{})

	resultSliceStr := []string{}
	for _, tempResultSlice := range resultSliceStrSlice {
		if resultMinLen == len(tempResultSlice) {
			resultSliceStr = tempResultSlice
		}
	}
	fmt.Println(strings.Join(resultSliceStr, " "))
}

func LoopCSimplepathRdr(pos int, visitSlice []string) {
	posArray := xyMap[pos]
	visitSliceNext := append(visitSlice, strconv.FormatInt(int64(pos), 10))
	tempLen := len(visitSliceNext)
	if pos == y {
		if resultMinLen < 0 || resultMinLen > tempLen {
			resultSliceStrSlice = append(resultSliceStrSlice, visitSliceNext)
			resultMinLen = tempLen
		}
	} else if (resultMinLen == -1 || tempLen < resultMinLen) && tempLen <= n {
		for _, conterPart := range posArray {
			if conterPart != x {
				LoopCSimplepathRdr(conterPart, visitSliceNext)
			}
		}
	}
}

func CSimplepathRdr(rdr *bufio.Reader) string {
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
