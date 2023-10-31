package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func CPeakMain() {
	var n, m int
	fmt.Scan(&n, &m)

	var rdr = bufio.NewReaderSize(os.Stdin, 10000)
	aStrArray := strings.Split(CPeakReadLine(rdr), " ")
	aArray := make([]int, n)

	for i := 0; i < n; i++ {
		aArray[i], _ = strconv.Atoi(aStrArray[i])
	}
	sort.Sort(sort.IntSlice(aArray))

	result := 0
	start := aArray[0]
	var tempSlice []int
	tempLen := 0

	for i := 0; i < n; i++ {
		target := aArray[i]
		sabun := target - start
		if sabun < m {
			tempSlice = append(tempSlice, target)
			tempLen = tempLen + 1
		} else {
			if tempLen == 1 {
				tempSlice = []int{}
				tempSlice = append(tempSlice, target)
				start = target
				tempLen = 1
			} else {
				// 距離が満たすまで取り除く
				for {
					tempSlice = tempSlice[1:]
					tempLen = tempLen - 1
					sabun := target - tempSlice[0]
					if sabun < m {
						start = tempSlice[0]
						tempSlice = append(tempSlice, target)
						tempLen = tempLen + 1
						break
					} else {
						if tempLen == 1 {
							tempSlice = []int{}
							tempSlice = append(tempSlice, target)
							start = target
							tempLen = 1
							break
						}
					}
				}
			}
		}

		if tempLen > result {
			result = tempLen
		}
	}
	fmt.Println(result)
}

func CPeakReadLine(rdr *bufio.Reader) string {
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
