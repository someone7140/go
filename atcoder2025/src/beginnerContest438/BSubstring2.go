package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func BSubstring2Main() {
	var n, m int
	fmt.Scan(&n, &m)
	var rdr = bufio.NewReaderSize(os.Stdin, 10000)
	s := strings.Split(BSubstring2Rdr(rdr), "")
	t := strings.Split(BSubstring2Rdr(rdr), "")

	result := 999999999999999999
	for i := range s {
		if i+m > n {
			break
		}
		tempResult := 0
		for j := range t {
			sNum, _ := strconv.Atoi(s[i+j])
			jNum, _ := strconv.Atoi(t[j])
			if sNum < jNum {
				tempResult = tempResult + (10 + sNum - jNum)
			} else {
				tempResult = tempResult + (sNum - jNum)
			}
		}

		if tempResult < result {
			result = tempResult
		}
	}

	fmt.Println(result)
}

func BSubstring2Rdr(rdr *bufio.Reader) string {
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
