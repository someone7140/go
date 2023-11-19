package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func CCountxxxMain() {
	var n int
	fmt.Scan(&n)
	var rdr = bufio.NewReaderSize(os.Stdin, 10000)
	s := CCountxxxReadLine(rdr)

	beforeMoji := ""
	renzoku := 0
	resultMapMoji := make(map[string]int)

	for _, sMoji := range strings.Split(s, "") {
		if beforeMoji != sMoji {
			beforeMoji = sMoji
			renzoku = 1
			_, ok := resultMapMoji[sMoji]
			if !ok {
				resultMapMoji[sMoji] = 1
			}
		} else {
			renzoku = renzoku + 1
			v, ok := resultMapMoji[sMoji]
			if ok && v < renzoku {
				resultMapMoji[sMoji] = renzoku
			}
		}
	}

	result := 0
	for _, v := range resultMapMoji {
		result = result + v
	}

	fmt.Println(result)

}

func CCountxxxReadLine(rdr *bufio.Reader) string {
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
