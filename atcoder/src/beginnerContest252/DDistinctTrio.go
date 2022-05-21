package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type DDistinctTrioKv struct {
	Key   int
	Value int
}

func DDistinctTrioMain() {
	var n int
	fmt.Scan(&n)
	var rdr = bufio.NewReaderSize(os.Stdin, 10000)
	aLine := DDistinctTrioRdr(rdr)
	aStrArray := strings.Split(aLine, " ")
	aMap := map[int]int{}

	for _, aStr := range aStrArray {
		aInt, _ := strconv.Atoi(aStr)
		v, ok := aMap[aInt]
		if ok {
			aMap[aInt] = v + 1
		} else {
			aMap[aInt] = 1
		}
	}

	var ss []DDistinctTrioKv
	for k, v := range aMap {
		ss = append(ss, DDistinctTrioKv{k, v})
	}
	sort.Slice(ss, func(i, j int) bool {
		return ss[i].Value > ss[j].Value
	})

	var result int64 = 0
	ssLength := len(ss)

	for i := 0; i < ssLength-2; i++ {
		topValue := ss[i].Value
		if topValue == 1 {
			result = result + int64((ssLength-i-1)*(ssLength-i-2)/2)
		} else {
			for j := i + 1; j < ssLength-1; j++ {
				secondValue := ss[j].Value
				if secondValue == 1 {
					result = result + int64((ssLength-j)*(ssLength-j-1)*topValue/2)
					break
				} else {
					temp12 := topValue * secondValue
					for l := j + 1; l < n; l++ {
						thirdValue := ss[l].Value
						if thirdValue == 1 {
							result = result + int64(temp12*(ssLength-l))
							break
						} else {
							result = result + int64(temp12*thirdValue)
						}
					}
				}
			}
		}
	}

	fmt.Println(result)
}

func DDistinctTrioRdr(rdr *bufio.Reader) string {
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
