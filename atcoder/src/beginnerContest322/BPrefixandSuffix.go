package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func BPrefixandSuffixMain() {
	var n, m int
	fmt.Scan(&n, &m)
	var rdr = bufio.NewReaderSize(os.Stdin, 10000)
	s := BPrefixandSuffixReadLine(rdr)
	t := BPrefixandSuffixReadLine(rdr)

	prefixFlag := strings.HasPrefix(t, s)
	suffixFlag := strings.HasSuffix(t, s)

	if prefixFlag && suffixFlag {
		fmt.Println(0)
	} else if prefixFlag && !suffixFlag {
		fmt.Println(1)
	} else if !prefixFlag && suffixFlag {
		fmt.Println(2)
	} else {
		fmt.Println(3)
	}
}

func BPrefixandSuffixReadLine(rdr *bufio.Reader) string {
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
