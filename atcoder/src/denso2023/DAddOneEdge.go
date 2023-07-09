package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	var n1, n2, m int
	fmt.Scan(&n1, &n2, &m)
	var rdr = bufio.NewReaderSize(os.Stdin, 10000)

	chouten1Map := map[int][]int{}
	chouten2Map := map[int][]int{}
	for i := 0; i < m; i++ {
		abList := strings.Split(DAddOneEdgeRdr(rdr), " ")
		aInt, _ := strconv.Atoi(abList[0])
		bInt, _ := strconv.Atoi(abList[1])

		if aInt <= n1 {
			aChoutens, okA := chouten1Map[aInt]
			if !okA {
				chouten1Map[aInt] = []int{bInt}
			} else {
				chouten1Map[aInt] = append(aChoutens, bInt)
			}
			bChoutens, okB := chouten1Map[bInt]
			if !okB {
				chouten1Map[bInt] = []int{aInt}
			} else {
				chouten1Map[bInt] = append(bChoutens, aInt)
			}
		} else {
			aChoutens, okA := chouten2Map[aInt]
			if !okA {
				chouten2Map[aInt] = []int{bInt}
			} else {
				chouten2Map[aInt] = append(aChoutens, bInt)
			}
			bChoutens, okB := chouten2Map[bInt]
			if !okB {
				chouten2Map[bInt] = []int{aInt}
			} else {
				chouten2Map[bInt] = append(bChoutens, aInt)
			}
		}
	}
	chouten1LongLen := 0
	chouten2LongLen := 0
	chouten1 := 1
	chouten2 := n1 + n2
	chouten1LenMap := map[int]int{}
	chouten1LenMap[chouten1] = 0
	chouten2LenMap := map[int]int{}
	chouten2LenMap[chouten2] = 0

	// 連結1の長さを求める
	var update1Func func(target int, len int)
	update1Func = func(target int, len int) {
		choutens := chouten1Map[target]
		for _, v := range choutens {
			if v != chouten1 {
				lenFromMap, ok := chouten1LenMap[v]
				newLen := len + 1
				if !ok || lenFromMap > newLen {
					chouten1LenMap[v] = newLen
					update1Func(v, newLen)
				}
			}

		}
	}

	// 連結2の長さを求める
	var update2Func func(target int, len int)
	update2Func = func(target int, len int) {
		choutens := chouten2Map[target]
		for _, v := range choutens {
			if v != chouten2 {
				lenFromMap, ok := chouten2LenMap[v]
				newLen := len + 1
				if !ok || lenFromMap > newLen {
					chouten2LenMap[v] = newLen
					update2Func(v, newLen)
				}
			}
		}
	}

	update1Func(chouten1, 0)
	update2Func(chouten2, 0)

	for _, v := range chouten1LenMap {
		if v > chouten1LongLen {
			chouten1LongLen = v
		}
	}
	for _, v := range chouten2LenMap {
		if v > chouten2LongLen {
			chouten2LongLen = v
		}
	}

	result := int64(chouten1LongLen) + int64(chouten2LongLen) + 1
	fmt.Println(result)
}

func DAddOneEdgeRdr(rdr *bufio.Reader) string {
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
