package main

import (
	"bufio"
	"fmt"
	"os"
)

func AHammingDistanceMain() {
	var n int
	fmt.Scan(&n)
	var rdr = bufio.NewReaderSize(os.Stdin, 10000)
	s := AHammingDistanceRdr(rdr)
	t := AHammingDistanceRdr(rdr)

	sRunes := []rune(s)
	tRunes := []rune(t)
	result := 0
	for i := 0; i < n; i++ {
		sMoji := string([]rune{sRunes[i]})
		tMoji := string([]rune{tRunes[i]})
		if sMoji != tMoji {
			result = result + 1
		}
	}

	fmt.Println(result)

}

func AHammingDistanceRdr(rdr *bufio.Reader) string {
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
