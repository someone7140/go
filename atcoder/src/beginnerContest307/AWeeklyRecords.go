package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func AWeeklyRecordsMain() {
	var n int
	fmt.Scan(&n)

	var rdr = bufio.NewReaderSize(os.Stdin, 10000)
	s := AWeeklyRecords(rdr)
	var resultSlice []string
	var tempResult int64
	tempResult = 0
	for i, sMoji := range strings.Split(s, " ") {
		temp, _ := strconv.Atoi(sMoji)
		tempResult = tempResult + int64(temp)
		if (i+1)%7 == 0 {
			resultSlice = append(resultSlice, strconv.FormatInt(tempResult, 10))
			tempResult = 0
		}
	}
	fmt.Println(strings.Join(resultSlice, " "))

}

func AWeeklyRecords(rdr *bufio.Reader) string {
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
