package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func BWhichisaheadMain() {
	var n int
	fmt.Scan(&n)
	var rdr = bufio.NewReaderSize(os.Stdin, 10000)
	sArray := strings.Split(BWhichisaheadReadLine(rdr), " ")

	posMap := map[string]int{}
	for i, s := range sArray {
		posMap[s] = i
	}
	q, _ := strconv.Atoi(BWhichisaheadReadLine(rdr))

	var resultSlice []string
	for i := 0; i < q; i++ {
		qArray := strings.Split(BWhichisaheadReadLine(rdr), " ")
		hito0Pos := posMap[qArray[0]]
		hito1Pos := posMap[qArray[1]]
		if hito0Pos < hito1Pos {
			resultSlice = append(resultSlice, qArray[0])
		} else {
			resultSlice = append(resultSlice, qArray[1])
		}
	}

	fmt.Println(strings.Join(resultSlice, "\n"))
}

func BWhichisaheadReadLine(rdr *bufio.Reader) string {
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
