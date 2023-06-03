package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type DAPieceofCakeXy struct {
	x int
	y int
}

func main() {
	var w, h int
	fmt.Scan(&w, &h)

	var n int
	fmt.Scan(&n)
	var rdr = bufio.NewReaderSize(os.Stdin, 10000)
	var xySlice = make([]DAPieceofCakeXy, n)

	for i := 0; i < n; i++ {
		xyString := strings.Split(DAPieceofCakeReadLine(rdr), " ")
		x, _ := strconv.Atoi(xyString[0])
		y, _ := strconv.Atoi(xyString[1])
		xySlice[i] = DAPieceofCakeXy{
			x: x,
			y: y,
		}
	}

	a, _ := strconv.Atoi(DAPieceofCakeReadLine(rdr))
	aStrSlice := strings.Split(DAPieceofCakeReadLine(rdr), " ")
	var aSlice = make([]int, a)
	for i := 0; i < a; i++ {
		aInt, _ := strconv.Atoi(aStrSlice[i])
		aSlice[i] = aInt
	}

	b, _ := strconv.Atoi(DAPieceofCakeReadLine(rdr))
	bStrSlice := strings.Split(DAPieceofCakeReadLine(rdr), " ")
	var bSlice = make([]int, b)
	for i := 0; i < b; i++ {
		bInt, _ := strconv.Atoi(bStrSlice[i])
		bSlice[i] = bInt
	}

	wharray := make([][]string, w)
	for i := 0; i < w; i++ {
		wharray[i] = make([]string, h)
	}

	aIndex := 0
	bIndex := 0
	countMap := map[string]int{}
	for i := 0; i < w; i++ {
		aKey := ""
		if aSlice[aIndex] <= i {
			if aIndex < a-1 {
				aIndex = aIndex + 1
				aKey = strconv.FormatInt(int64(aSlice[aIndex]), 10) + "before"
			} else {
				aKey = strconv.FormatInt(int64(aSlice[aIndex]), 10) + "after"
			}
		} else {
			aKey = strconv.FormatInt(int64(aSlice[aIndex]), 10) + "before"
		}
		for j := 0; j < h; j++ {
			bKey := ""
			if bSlice[bIndex] <= j {
				if bIndex < b-1 {
					bIndex = bIndex + 1
					bKey = strconv.FormatInt(int64(bSlice[bIndex]), 10) + "before"
				} else {
					bKey = strconv.FormatInt(int64(bSlice[bIndex]), 10) + "after"
				}
			} else {
				bKey = strconv.FormatInt(int64(bSlice[bIndex]), 10) + "before"
			}
			key := aKey + "-" + bKey
			wharray[i][j] = aKey + "-" + bKey
			countMap[key] = 0
		}
	}
	for i := 0; i < n; i++ {
		xy := xySlice[i]
		key := wharray[xy.x-1][xy.y-1]
		v, _ := countMap[key]
		countMap[key] = v + 1
	}

	min := -1
	max := -1

	for _, v := range countMap {
		if min == -1 || v < min {
			min = v
		}
		if max == -1 || v > max {
			max = v
		}
	}

	fmt.Println(strconv.FormatInt(int64(min), 10) + " " + strconv.FormatInt(int64(max), 10))
}

func DAPieceofCakeReadLine(rdr *bufio.Reader) string {
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
