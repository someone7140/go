package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	var n int
	fmt.Scan(&n)

	var rdr = bufio.NewReaderSize(os.Stdin, 10000)
	s := DGomamayoSequenceReadLine(rdr)

	cList := make([]int64, n)
	cStrArray := strings.Split(DGomamayoSequenceReadLine(rdr), " ")
	minC := int64(-1)
	for i := 0; i < n; i++ {
		c, _ := strconv.Atoi(cStrArray[i])
		c64 := int64(c)
		if minC == -1 || minC > c64 {
			minC = c64
		}
		cList[i] = int64(c)
	}

	nowDuplicateCount := 0
	for i := 0; i < n-1; i++ {
		s1 := string(s[i])
		s2 := string(s[i+1])
		if s1 == s2 {
			nowDuplicateCount = nowDuplicateCount + 1
		}
	}

	var result int64
	result = -1
	if nowDuplicateCount == 1 {
		result = 0
	} else {
		sMap := map[string]int64{}
		sMap[s] = 0

		result = -1

		var replaceAtIndex2 func(str string, replacement string, index int) string
		replaceAtIndex2 = func(str string, replacement string, index int) string {
			return str[:index] + replacement + str[index+1:]
		}

		var funcDp func(duplicateCount int, cost int64, index int, nowS string)
		funcDp = func(duplicateCount int, cost int64, index int, nowS string) {
			if index == n {
				return
			}

			// 変えない
			funcDp(duplicateCount, cost, index+1, nowS)
			// 変える
			targetString := string(nowS[index])
			if targetString == "1" {
				targetString = "0"
			} else {
				targetString = "1"
			}
			newStr := replaceAtIndex2(nowS, targetString, index)
			newCost := cost + cList[index]

			v, ok := sMap[newStr]

			if !ok {
				sMap[newStr] = newCost
			} else {
				if v < newCost {
					return
				} else {
					sMap[newStr] = newCost
				}
			}

			zougen := 0
			if index < n-1 {
				afterString := string(nowS[index+1])
				if targetString == afterString {
					zougen = zougen + 1
				} else {
					zougen = zougen - 1
				}
			}

			if index > 0 {
				beforeString := string(nowS[index-1])
				if targetString == beforeString {
					zougen = zougen + 1
				} else {
					zougen = zougen - 1
				}
			}

			newDuplicateCount := duplicateCount + zougen

			if newDuplicateCount == 1 {
				if result == -1 || result > newCost {
					result = newCost
				}
			} else {
				funcDp(newDuplicateCount, newCost, index+1, newStr)
			}
		}
		funcDp(nowDuplicateCount, 0, 0, s)
	}

	fmt.Println(result)
}

func DGomamayoSequenceReadLine(rdr *bufio.Reader) string {
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
