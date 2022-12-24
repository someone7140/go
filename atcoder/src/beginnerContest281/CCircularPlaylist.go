package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func CCircularPlaylistMain() {
	var n, t int64
	fmt.Scan(&n, &t)

	var rdr = bufio.NewReaderSize(os.Stdin, 10000)
	aStrArray := strings.Split(CCircularPlaylistReadLine(rdr), " ")
	aArray := make([]int64, n)

	var i int64
	var sumA int64
	sumA = 0
	for i = 0; i < n; i++ {
		a, _ := strconv.Atoi(aStrArray[i])
		aArray[i] = int64(a)
		sumA = sumA + int64(a)
	}

	amari := t % sumA
	var resultIndex int64
	var resultByou int64

	for i = 0; i < n; i++ {
		a := aArray[i]
		if a > amari {
			resultIndex = i + 1
			resultByou = amari
			break
		} else {
			amari = amari - a
		}
	}
	fmt.Printf("%d %d", resultIndex, resultByou)
}

func CCircularPlaylistReadLine(rdr *bufio.Reader) string {
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
