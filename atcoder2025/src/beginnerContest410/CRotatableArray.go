package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func CRotatableArrayMain() {
	var n, q int
	fmt.Scan(&n, &q)
	var rdr = bufio.NewReaderSize(os.Stdin, 10000)

	aList := make([]int, n)
	for i := 1; i <= n; i++ {
		aList[i-1] = i
	}

	var resultSlice []string
	sentou := 0
	for i := 0; i < q; i++ {
		qStrList := strings.Split(CRotatableArrayRdr(rdr), " ")
		if qStrList[0] == "1" {
			aIndex, _ := strconv.Atoi(qStrList[1])
			aIndex = aIndex - 1
			aIndex = (sentou + aIndex) % n
			aVal, _ := strconv.Atoi(qStrList[2])
			aList[aIndex] = aVal
		} else if qStrList[0] == "2" {
			aIndex, _ := strconv.Atoi(qStrList[1])
			aIndex = aIndex - 1
			aIndex = (sentou + aIndex) % n
			resultSlice = append(resultSlice, strconv.FormatInt(int64(aList[aIndex]), 10))
		} else {
			k, _ := strconv.Atoi(qStrList[1])
			newSentou := (k + sentou) % n
			sentou = newSentou
		}
	}

	fmt.Println(strings.Join(resultSlice, "\n"))
}

func CRotatableArrayRdr(rdr *bufio.Reader) string {
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
