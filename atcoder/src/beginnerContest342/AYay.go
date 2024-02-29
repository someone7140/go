package main

import (
	"bufio"
	"fmt"
	"os"
)

func AYayMain() {
	var rdr = bufio.NewReaderSize(os.Stdin, 10000)
	s := AYayReadLine(rdr)

	countMap := map[string]int{}
	firstIndexMap := map[string]int{}
	for i, c := range s {
		sTan := string([]rune{c})
		v, ok := countMap[sTan]
		if !ok {
			countMap[sTan] = 1
			firstIndexMap[sTan] = i + 1
		} else {
			countMap[sTan] = v + 1
		}
	}
	var result int
	for k, v := range countMap {
		if v == 1 {
			result = firstIndexMap[k]
			break
		}
	}
	fmt.Println(result)
}

func AYayReadLine(rdr *bufio.Reader) string {
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
