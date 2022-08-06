package main

import (
	"bufio"
	"fmt"
	"os"
)

func BTournamentResultMain() {
	var n int
	fmt.Scan(&n)

	var rdr = bufio.NewReaderSize(os.Stdin, 10000)

	aArrayArray := make([][]string, n)

	result := "correct"
	for i := 0; i < n; i++ {
		aArray := make([]string, n)
		aStr := BTournamentResultRdr(rdr)
		for i2, c := range aStr {
			aMoji := string([]rune{c})
			aArray[i2] = aMoji
		}
		aArrayArray[i] = aArray
	}

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			a := aArrayArray[i][j]
			a2 := aArrayArray[j][i]

			if a == "W" {
				if a2 != "L" {
					result = "incorrect"
				}
			} else if a == "L" {
				if a2 != "W" {
					result = "incorrect"
				}
			} else if a == "D" {
				if a2 != "D" {
					result = "incorrect"
				}
			}
		}
	}

	fmt.Println(result)
}

func BTournamentResultRdr(rdr *bufio.Reader) string {
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
