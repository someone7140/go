package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func BRotateMain() {
	var n int
	fmt.Scan(&n)
	var rdr = bufio.NewReaderSize(os.Stdin, 10000)

	bArrayArray := make([][]string, n)
	bArrayArrayResult := make([][]string, n)

	for i := 0; i < n; i++ {
		bArray := make([]string, n)
		bStrList := BRotateRdr(rdr)
		for i2, c := range bStrList {
			b := string([]rune{c})
			bArray[i2] = b
		}
		bArrayArray[i] = bArray

		bArrayCopy := make([]string, n)
		bArrayArrayResult[i] = bArrayCopy
	}

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if i == 0 {
				if j == n-1 {
					bArrayArrayResult[i+1][j] = bArrayArray[i][j]
				} else {
					bArrayArrayResult[i][j+1] = bArrayArray[i][j]
				}
			} else if i == n-1 {
				if j == 0 {
					bArrayArrayResult[i-1][j] = bArrayArray[i][j]
				} else {
					bArrayArrayResult[i][j-1] = bArrayArray[i][j]
				}
			} else {
				if j == 0 {
					bArrayArrayResult[i-1][j] = bArrayArray[i][j]
				} else if j == n-1 {
					bArrayArrayResult[i+1][j] = bArrayArray[i][j]
				} else {
					bArrayArrayResult[i][j] = bArrayArray[i][j]
				}
			}
		}
	}
	for i := 0; i < n; i++ {
		fmt.Println(strings.Join(bArrayArrayResult[i], ""))
	}

}

func BRotateRdr(rdr *bufio.Reader) string {
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
