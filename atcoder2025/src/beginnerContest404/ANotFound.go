package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func ANotFoundMain() {
	var rdr = bufio.NewReaderSize(os.Stdin, 10000)
	sList := strings.Split(ANotFoundRdr(rdr), "")

	result := ""
	for i := 0; i < 26; i++ {
		alpha := string('a' + i)
		findFlag := false
		for _, s := range sList {
			if s == alpha {
				findFlag = true
				break
			}
		}
		if !findFlag {
			result = alpha
			break
		}
	}
	fmt.Println(result)
}

func ANotFoundRdr(rdr *bufio.Reader) string {
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
