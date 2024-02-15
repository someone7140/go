package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func EBadJuiceMain() {
	var n int
	fmt.Scan(&n)

	var juiceList = make([]string, n-1)
	for i := 0; i < n-1; i++ {
		juiceList[i] = strconv.FormatInt(int64(i+1), 10)
	}
	juiceListLen := len(juiceList)
	var kaList = make([]string, n-1)
	for i := 0; i < n-1; i++ {
		honsuu := strconv.FormatInt(int64(juiceListLen-i), 10)
		kaList[i] = honsuu + " " + strings.Join(juiceList[i:], " ")
	}

	fmt.Println(n - 1)
	for i := 0; i < n-1; i++ {
		fmt.Println(kaList[i])
	}

	var rdr = bufio.NewReaderSize(os.Stdin, 10000)
	s := EBadJuiceReadLine(rdr)

	result := n
	dokuFlag := false
	for i, c := range s {
		sMoji := string([]rune{c})
		if sMoji == "1" {
			dokuFlag = true
		} else {
			if dokuFlag {
				result = i
				break
			}
		}
	}
	if dokuFlag && result == n {
		result = result - 1
	}
	fmt.Println(result)
}

func EBadJuiceReadLine(rdr *bufio.Reader) string {
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
