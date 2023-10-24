package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

func DSquarePermutationMain() {
	var n int
	fmt.Scan(&n)
	var rdr = bufio.NewReaderSize(os.Stdin, 10000)
	inputStrs := strings.Split(DSquarePermutationReadLine(rdr), "")
	sArray := make([]int, n)

	sMap := map[int]int{}
	for i := 0; i < n; i++ {
		s, _ := strconv.Atoi(inputStrs[i])
		sArray[i] = s
		v, ok := sMap[s]
		if ok {
			sMap[s] = 1 + v
		} else {
			sMap[s] = 1
		}
	}
	sArrayAsc := make([]int, n)
	copy(sArrayAsc, sArray)
	sort.Sort(sort.IntSlice(sArrayAsc))

	sArrayDesc := make([]int, n)
	copy(sArrayDesc, sArray)
	sort.Sort(sort.Reverse(sort.IntSlice(sArrayDesc)))

	sArrayAscStr := ""
	for _, s := range sArrayAsc {
		sArrayAscStr = sArrayAscStr + strconv.FormatInt(int64(s), 10)
	}
	sArrayAscInt64, _ := strconv.ParseInt(sArrayAscStr, 10, 64)

	sArrayDescStr := ""
	for _, s := range sArrayDesc {
		sArrayDescStr = sArrayDescStr + strconv.FormatInt(int64(s), 10)
	}
	sArrayDescInt64, _ := strconv.ParseInt(sArrayDescStr, 10, 64)

	sArrayAscSqrt := int64(math.Sqrt(float64(sArrayAscInt64)))
	sArrayDescSqrt := int64(math.Sqrt(float64(sArrayDescInt64))) + 1
	result := 0
	for i := sArrayAscSqrt; i <= sArrayDescSqrt; i++ {
		jijouMap := map[int]int{}
		jijou := i * i
		jijouStr := strconv.FormatInt(jijou, 10)
		jijouMap[0] = n - len(jijouStr)
		for _, c := range jijouStr {
			sMoji := string([]rune{c})
			sTemp, _ := strconv.Atoi(sMoji)
			v, ok := jijouMap[sTemp]
			if ok {
				jijouMap[sTemp] = 1 + v
			} else {
				jijouMap[sTemp] = 1
			}
		}
		tempResult := true
		for j := 0; j < 10; j++ {
			moto, ok := sMap[j]
			if !ok {
				moto = 0
			}
			saki, ok2 := jijouMap[j]
			if !ok2 {
				saki = 0
			}
			if moto != saki {
				tempResult = false
				break
			}
		}
		if tempResult {
			result = result + 1
		}
	}
	fmt.Println(result)

}

func DSquarePermutationReadLine(rdr *bufio.Reader) string {
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
