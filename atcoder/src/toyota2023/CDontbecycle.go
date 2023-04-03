package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	var n, m int
	fmt.Scan(&n, &m)

	var rdr = bufio.NewReaderSize(os.Stdin, 10000)
	abMap := map[int][]int{}
	visitedMap := map[int]bool{}
	duplicateRoute := map[string]struct{}{}

	for i := 0; i < m; i++ {
		abArr := strings.Split(CDontbecycleReadLine(rdr), " ")
		a, _ := strconv.Atoi(abArr[0])
		b, _ := strconv.Atoi(abArr[1])

		aRes, aOk := abMap[a]
		if aOk {
			abMap[a] = append(aRes, b)
		} else {
			abMap[a] = []int{b}
		}

		bRes, bOk := abMap[b]
		if bOk {
			abMap[b] = append(bRes, a)
		} else {
			abMap[b] = []int{a}
		}
	}

	var loopFunc func(nextNum int, beforeNum int)
	loopFunc = func(nextNum int, beforeNum int) {
		nextArray := abMap[nextNum]
		for _, v := range nextArray {
			if beforeNum != v {
				res, ok := visitedMap[v]
				if res && ok {
					min := 0
					max := 0
					if nextNum < v {
						min = nextNum
						max = v
					} else {
						min = v
						max = nextNum
					}
					duplicateRoute[strconv.FormatInt(int64(min), 10)+"-"+strconv.FormatInt(int64(max), 10)] = struct{}{}
				} else {
					visitedMap[v] = true
					loopFunc(v, nextNum)
				}
			}
		}
	}
	// 1からスタート
	visitedMap[1] = true
	loopFunc(1, -99999)
	fmt.Println(len(duplicateRoute))
}

func CDontbecycleReadLine(rdr *bufio.Reader) string {
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
