package main

import (
	"bufio"
	"fmt"
	"os"
)

func BracecarMain() {
	var n int
	fmt.Scan(&n)

	var rdr = bufio.NewReaderSize(os.Stdin, 10000)
	var sSlice []string
	for i := 0; i < n; i++ {
		sSlice = append(sSlice, BracecarReadLine(rdr))
	}
	result := "No"

	var funcReverse func(s string) string
	funcReverse = func(s string) string {
		runes := []rune(s)
		for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
			runes[i], runes[j] = runes[j], runes[i]
		}
		return string(runes)
	}

	for i := 0; i < n; i++ {
		temp1 := sSlice[i]
		for j := 0; j < n; j++ {
			if i != j {
				temp2 := temp1 + sSlice[j]
				temp2Reverse := funcReverse(temp2)
				if temp2 == temp2Reverse {
					result = "Yes"
					break
				}
			}
		}
		if result == "Yes" {
			break
		}
	}
	fmt.Println(result)

}

func BracecarReadLine(rdr *bufio.Reader) string {
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
