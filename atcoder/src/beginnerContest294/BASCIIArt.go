package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func BASCIIArtMain() {
	var h, w int
	fmt.Scan(&h, &w)

	var rdr = bufio.NewReaderSize(os.Stdin, 10000)

	var resultSlice []string
	for i := 0; i < h; i++ {
		result := ""
		aStrArray := strings.Split(BASCIIArtReadLine(rdr), " ")
		for j := 0; j < w; j++ {
			a, _ := strconv.Atoi(aStrArray[j])
			if a == 0 {
				result = result + "."
			} else {
				result = result + string(rune('A'+(a-1)))
			}
		}

		resultSlice = append(resultSlice, result)
	}
	fmt.Println(strings.Join(resultSlice, "\n"))

}

func BASCIIArtReadLine(rdr *bufio.Reader) string {
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
