package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func DBankMain() {
	var n, q int
	fmt.Scan(&n, &q)

	var rdr = bufio.NewReaderSize(os.Stdin, 10000)
	nextCallNumber := 1
	var uketsukeArray []int
	deleteMap := map[int]int{}
	var resultSlice []string
	for i := 0; i < q; i++ {
		qStr := DBankReadLine(rdr)
		if qStr == "1" {
			uketsukeArray = append(uketsukeArray, nextCallNumber)
			nextCallNumber = nextCallNumber + 1
		} else if qStr == "3" {
			resultSlice = append(resultSlice, strconv.FormatInt(int64(uketsukeArray[0]), 10))
		} else {
			qStrArray := strings.Split(qStr, " ")
			hitoNum, _ := strconv.Atoi(qStrArray[1])
			deleteMap[hitoNum] = 1
			if uketsukeArray[0] == hitoNum {
				uketsukeTempArray := []int{}
				for _, v := range uketsukeArray {
					_, ok := deleteMap[v]
					if !ok {
						uketsukeTempArray = append(uketsukeTempArray, v)
					}
				}
				uketsukeArray = uketsukeTempArray
				deleteMap = map[int]int{}
			}
		}
	}

	fmt.Println(strings.Join(resultSlice, "\n"))
}

func DBankReadLine(rdr *bufio.Reader) string {
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
