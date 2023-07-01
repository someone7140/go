package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func BDefaultPriceMain() {
	var n, m int
	fmt.Scan(&n, &m)

	var rdr = bufio.NewReaderSize(os.Stdin, 10000)
	cList := strings.Split(BDefaultPriceReadLine(rdr), " ")
	dList := strings.Split(BDefaultPriceReadLine(rdr), " ")

	var pList []int64
	pListStr := BDefaultPriceReadLine(rdr)
	for _, pMoji := range strings.Split(pListStr, " ") {
		p, _ := strconv.Atoi(pMoji)
		pList = append(pList, int64(p))
	}

	var result int64
	result = 0

	for _, c := range cList {
		findIndex := -1
		for j, d := range dList {
			if d == c {
				findIndex = j
				break
			}
		}

		if findIndex == -1 {
			result = result + pList[0]
		} else {
			result = result + pList[findIndex+1]
		}
	}

	fmt.Println(result)

}

func BDefaultPriceReadLine(rdr *bufio.Reader) string {
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
