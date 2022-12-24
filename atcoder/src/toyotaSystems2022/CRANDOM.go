package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func CRANDOMMain() {
	var h, w int
	fmt.Scan(&h, &w)

	var rdr = bufio.NewReaderSize(os.Stdin, 10000)

	sArrayArray := make([][]string, w)
	for i := 0; i < w; i++ {
		sArrayArray[i] = make([]string, h)
	}

	for i := 0; i < h; i++ {
		s := CRANDOMReadLine(rdr)
		for j, c := range s {
			sMoji := string([]rune{c})
			sArrayArray[j][i] = sMoji
		}
	}
	sMap := map[string]int{}
	for i := 0; i < w; i++ {
		sArray := sArrayArray[i]
		sArrayJoin := strings.Join(sArray, "")
		count, ok := sMap[sArrayJoin]
		if ok {
			sMap[sArrayJoin] = count + 1
		} else {
			sMap[sArrayJoin] = 1
		}
	}

	tArrayArray := make([][]string, w)
	for i := 0; i < w; i++ {
		tArrayArray[i] = make([]string, h)
	}
	for i := 0; i < h; i++ {
		t := CRANDOMReadLine(rdr)
		for j, c := range t {
			tMoji := string([]rune{c})
			tArrayArray[j][i] = tMoji
		}
	}

	result := "Yes"
	for i := 0; i < w; i++ {
		tArray := tArrayArray[i]
		tArrayJoin := strings.Join(tArray, "")
		count, ok := sMap[tArrayJoin]
		if !ok {
			result = "No"
			break
		} else {
			if count == 0 {
				result = "No"
				break
			} else {
				sMap[tArrayJoin] = count - 1
			}
		}
	}
	fmt.Println(result)
}

func CRANDOMReadLine(rdr *bufio.Reader) string {
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
