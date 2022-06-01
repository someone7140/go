package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

func BDistanceBetweenTokensMain() {
	var h, w int
	fmt.Scan(&h, &w)
	var rdr = bufio.NewReaderSize(os.Stdin, 10000)

	firstH := 0
	firstW := 0
	secondFlag := false
	secondH := 0
	secondW := 0

	for i := 0; i < h; i++ {
		s := BDistanceBetweenTokensRdr(rdr)
		for i2, c := range s {
			sMoji := string([]rune{c})
			if sMoji == "o" {
				if secondFlag {
					secondH = i
					secondW = i2
				} else {
					firstH = i
					firstW = i2
					secondFlag = true
				}
			}
		}
	}
	wide := math.Abs(float64(firstH - secondH))
	height := math.Abs(float64(firstW - secondW))
	fmt.Println(int(wide + height))
}

func BDistanceBetweenTokensRdr(rdr *bufio.Reader) string {
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
