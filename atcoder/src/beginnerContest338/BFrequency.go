package main

import (
	"bufio"
	"fmt"
	"os"
)

func BFrequencyMain() {
	var rdr = bufio.NewReaderSize(os.Stdin, 10000)
	s := BFrequencyReadLine(rdr)

	sMap := map[string]int{}
	// 回数を記録
	for _, c := range s {
		sMoji := string([]rune{c})
		res, ok := sMap[sMoji]

		if ok {
			sMap[sMoji] = res + 1
		} else {
			sMap[sMoji] = 1
		}
	}

	result := ""
	resultCount := 0

	for k, v := range sMap {
		if result == "" {
			result = k
			resultCount = v
		} else if resultCount == v && k < result {
			result = k
			resultCount = v
		} else if resultCount < v {
			result = k
			resultCount = v
		}
	}

	fmt.Println(result)

}

func BFrequencyReadLine(rdr *bufio.Reader) string {
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
