package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func CManyReplacementMain() {
	var n int
	fmt.Scan(&n)
	var rdr = bufio.NewReaderSize(os.Stdin, 10000)
	s := CManyReplacementReadLine(rdr)

	q, _ := strconv.Atoi(CManyReplacementReadLine(rdr))

	changeMap := map[string]string{}
	for i := 0; i < q; i++ {
		qArray := strings.Split(CManyReplacementReadLine(rdr), " ")
		mae := qArray[0]
		ato := qArray[1]

		for k, v := range changeMap {
			if v == mae {
				changeMap[k] = ato
			}
		}
		_, ok := changeMap[mae]
		if !ok {
			changeMap[mae] = ato
		}
	}

	var resultSlice []string
	for _, c := range s {
		sTan := string([]rune{c})
		v, ok := changeMap[sTan]
		if !ok {
			resultSlice = append(resultSlice, sTan)
		} else {
			resultSlice = append(resultSlice, v)
		}
	}
	fmt.Println(strings.Join(resultSlice, ""))
}

func CManyReplacementReadLine(rdr *bufio.Reader) string {
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
