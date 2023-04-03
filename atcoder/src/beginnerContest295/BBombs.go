package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func BBombsMain() {
	var r, c int
	fmt.Scan(&r, &c)

	var rdr = bufio.NewReaderSize(os.Stdin, 10000)
	rcArrayArray := make([][]string, r)

	for i := 0; i < r; i++ {
		cArray := make([]string, c)
		cStr := BBombsReadLine(rdr)
		for j, c := range cStr {
			cMoji := string([]rune{c})
			cArray[j] = cMoji
		}
		rcArrayArray[i] = cArray
	}

	resultArray := make([][]string, r)
	for i := 0; i < r; i++ {
		cArray := make([]string, c)
		copy(cArray, rcArrayArray[i])
		resultArray[i] = cArray
	}

	var loopFunc func(nowNum int, bombCount int, nowR int, nowC int)
	loopFunc = func(nowNum int, bombCount int, nowR int, nowC int) {
		resultArray[nowR][nowC] = "."
		if nowNum < bombCount {
			if nowR != 0 {
				loopFunc(nowNum+1, bombCount, nowR-1, nowC)
			}
			if nowR != r-1 {
				loopFunc(nowNum+1, bombCount, nowR+1, nowC)
			}
			if nowC != 0 {
				loopFunc(nowNum+1, bombCount, nowR, nowC-1)
			}
			if nowC != c-1 {
				loopFunc(nowNum+1, bombCount, nowR, nowC+1)
			}
		}

	}

	for ri, cArray := range rcArrayArray {
		for ci, masu := range cArray {
			if masu != "." && masu != "#" {
				bomb, _ := strconv.Atoi(masu)
				loopFunc(0, bomb, ri, ci)
			}
		}
	}

	for _, cArray := range resultArray {
		fmt.Println(strings.Join(cArray, ""))
	}

}

func BBombsReadLine(rdr *bufio.Reader) string {
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
