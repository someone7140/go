package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func BPermutetoMinimizeMain() {
	var rdr = bufio.NewReaderSize(os.Stdin, 10000)
	xStrArray := strings.Split(BPermutetoMinimizeRdr(rdr), "")
	xLen := len(xStrArray)
	xArray := make([]int, xLen)
	for i := 0; i < xLen; i++ {
		a, _ := strconv.Atoi(xStrArray[i])
		xArray[i] = a
	}

	sort.Ints(xArray)
	resultList := make([]string, xLen)
	irekaeFlag := true
	for i := 0; i < xLen; i++ {
		if i == 0 {
			irekaeFlag = xArray[i] == 0
		} else {
			if irekaeFlag {
				if xArray[i] != 0 {
					temp := xArray[i]
					xArray[0] = temp
					xArray[i] = 0
					break
				}
			} else {
				break
			}
		}
	}

	for i := 0; i < xLen; i++ {
		resultList[i] = strconv.FormatInt(int64(xArray[i]), 10)
	}

	fmt.Println(strings.Join(resultList, ""))
}

func BPermutetoMinimizeRdr(rdr *bufio.Reader) string {
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
