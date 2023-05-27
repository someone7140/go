package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func BDiscordMain() {
	var n, m int
	fmt.Scan(&n, &m)
	var rdr = bufio.NewReaderSize(os.Stdin, 10000)
	aMap := map[string]bool{}
	result := 0

	for i := 0; i < m; i++ {
		mArray := strings.Split(BDiscordReadLine(rdr), " ")
		for j := 0; j < n-1; j++ {
			a1, _ := strconv.Atoi(mArray[j])
			a2, _ := strconv.Atoi(mArray[j+1])
			aBefore := 0
			aAfter := 0
			if a1 < a2 {
				aBefore = a1
				aAfter = a2
			} else {
				aBefore = a2
				aAfter = a1
			}
			key := strconv.FormatInt(int64(aBefore), 10) + "-" + strconv.FormatInt(int64(aAfter), 10)
			aMap[key] = true
		}
	}
	for i := 1; i < n; i++ {
		for j := i + 1; j <= n; j++ {
			key := strconv.FormatInt(int64(i), 10) + "-" + strconv.FormatInt(int64(j), 10)
			v, ok := aMap[key]
			if !v || !ok {
				result = result + 1
			}
		}
	}
	fmt.Println(result)

}

func BDiscordReadLine(rdr *bufio.Reader) string {
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
