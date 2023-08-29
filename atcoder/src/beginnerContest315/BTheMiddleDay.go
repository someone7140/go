package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func BTheMiddleDayMain() {
	var m int
	fmt.Scan(&m)

	var rdr = bufio.NewReaderSize(os.Stdin, 10000)
	var dSlice = make([]int, m)
	var dRuisekiSlice = make([]int, m)
	var sum = 0

	for i, dMoji := range strings.Split(BTheMiddleDayReadLine(rdr), " ") {
		d, _ := strconv.Atoi(dMoji)
		dSlice[i] = d
		dRuisekiSlice[i] = sum + d
		sum = sum + d
	}

	targetDay := (sum + 1) / 2
	result1 := 0
	result2 := 0

	for i, dRui := range dRuisekiSlice {
		if dRui >= targetDay {
			temp := dRui - targetDay
			d := dSlice[i]
			result1 = i + 1
			result2 = d - temp
			break
		}
	}

	fmt.Println(strconv.FormatInt(int64(result1), 10) + " " + strconv.FormatInt(int64(result2), 10))
}

func BTheMiddleDayReadLine(rdr *bufio.Reader) string {
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
