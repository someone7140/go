package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	var n int
	fmt.Scan(&n)
	var rdr = bufio.NewReaderSize(os.Stdin, 10000)
	sMap := map[int64]int64{}
	sArray := make([]int64, n)

	// SCの配列を読み込み
	for i := 0; i < n; i++ {
		scStrs := strings.Split(DMergeSlimesReadLine(rdr), " ")
		s, _ := strconv.Atoi(scStrs[0])
		c, _ := strconv.Atoi(scStrs[1])
		sArray[i] = int64(s)
		sMap[int64(s)] = int64(c)
	}
	sort.Slice(sArray, func(i, j int) bool {
		return sArray[i] < sArray[j]
	})
	arrayLen := n

	var funcInsertNibunArray func(value int64)
	funcInsertNibunArray = func(value int64) {
		start := 0
		end := arrayLen - 1
		half := (start + end) / 2
		for {
			if start >= half || end <= half {
				break
			}
			if sArray[half] < value {
				start = half
				half = (start + end) / 2
			} else if sArray[half] >= value {
				end = half
				half = (start + end) / 2
			}
		}
		sArray = append(sArray[:half+1], sArray[half:]...)
		sArray[half+1] = value
	}

	for {
		if arrayLen == 0 {
			break
		}
		// 最初
		key := sArray[0]
		value := sMap[key]
		if value > 1 {
			keyDouble := key * int64(2)
			countDouble := value / int64(2)
			if value%2 == 1 {
				sMap[key] = 1
			} else {
				delete(sMap, key)
			}
			// 2倍したものが存在
			_, ok := sMap[keyDouble]
			if ok {
				sMap[keyDouble] = sMap[keyDouble] + countDouble
			} else {
				sMap[keyDouble] = countDouble
				// 最大値より大きい
				if sArray[arrayLen-1] < keyDouble {
					sArray = append(sArray, keyDouble)
				} else {
					funcInsertNibunArray(keyDouble)
				}
				arrayLen = arrayLen + 1
			}
		}
		if arrayLen == 1 {
			break
		} else {
			// 配列から削除
			sArray = sArray[1:]
			arrayLen = arrayLen - 1
		}
	}

	fmt.Println(len(sMap))

}

func DMergeSlimesReadLine(rdr *bufio.Reader) string {
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
