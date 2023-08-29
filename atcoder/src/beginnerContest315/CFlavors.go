package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func CFlavorsMain() {
	var n int
	fmt.Scan(&n)

	type Ice struct {
		aji     int
		oishisa int
	}

	var rdr = bufio.NewReaderSize(os.Stdin, 10000)
	var iceSlice = make([]Ice, n)
	for i := 0; i < n; i++ {
		fs := strings.Split(CFlavorsReadLine(rdr), " ")
		f, _ := strconv.Atoi(fs[0])
		s, _ := strconv.Atoi(fs[1])
		iceSlice[i] = Ice{
			aji:     f,
			oishisa: s,
		}
	}
	sort.Slice(iceSlice, func(i, j int) bool { return iceSlice[i].oishisa > iceSlice[j].oishisa })
	ice1 := iceSlice[0]
	result := 0
	for i := 1; i < n; i++ {
		ice2 := iceSlice[i]
		tempResult := ice1.oishisa + ice2.oishisa
		if ice1.aji == ice2.aji {
			tempResult = ice1.oishisa + ice2.oishisa/2
		}
		if tempResult > result {
			result = tempResult
		} else {
			if ice1.aji != ice2.aji {
				break
			}
		}
	}

	fmt.Println(result)
}

func CFlavorsReadLine(rdr *bufio.Reader) string {
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
