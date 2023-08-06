package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func AToBeSaikyoMain() {
	var n int
	fmt.Scan(&n)
	var rdr = bufio.NewReaderSize(os.Stdin, 10000)
	pListStr := AToBeSaikyoReadLine(rdr)
	max := -1
	first := -1
	for i, pMoji := range strings.Split(pListStr, " ") {
		p, _ := strconv.Atoi(pMoji)
		if i == 0 {
			first = p
		} else {
			if p > max {
				max = p
			}
		}
	}

	if max < first {
		fmt.Println(0)
	} else {
		fmt.Println(max - first + 1)
	}

}

func AToBeSaikyoReadLine(rdr *bufio.Reader) string {
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
