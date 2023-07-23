package main

import (
	"bufio"
	"fmt"
	"os"
)

func CReversibleMain() {
	var n int
	fmt.Scan(&n)

	var rdr = bufio.NewReaderSize(os.Stdin, 10000)
	sMap := map[string]bool{}

	var reverseStr func(s string) string
	reverseStr = func(s string) string {
		runes := []rune(s)
		for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
			runes[i], runes[j] = runes[j], runes[i]
		}
		return string(runes)
	}

	for i := 0; i < n; i++ {
		s := CReversibleReadLine(rdr)
		reverseS := reverseStr(s)
		_, ok1 := sMap[s]
		_, ok2 := sMap[reverseS]
		if !ok1 && !ok2 {
			sMap[s] = true
		}
	}
	result := 0
	for _, v := range sMap {
		if v {
			result = result + 1
		}
	}
	fmt.Println(result)

}

func CReversibleReadLine(rdr *bufio.Reader) string {
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
