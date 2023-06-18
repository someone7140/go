package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func BBase2Main() {
	var rdr = bufio.NewReaderSize(os.Stdin, 10000)
	astrList := strings.Split(BBase2ReadLine(rdr), " ")
	var result uint64
	result = 0
	for i, aStr := range astrList {
		if aStr == "1" {
			var tempResult uint64
			tempResult = 1
			for j := 0; j < i; j++ {
				tempResult = tempResult * uint64(2)
			}
			result = result + tempResult
		}
	}
	fmt.Println(result)

}

func BBase2ReadLine(rdr *bufio.Reader) string {
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
