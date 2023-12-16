package main

import (
	"bufio"
	"fmt"
	"os"
)

func BPentagonMain() {
	var rdr = bufio.NewReaderSize(os.Stdin, 10000)
	s := BPentagonReadLine(rdr)
	t := BPentagonReadLine(rdr)

	getNagasa := func(pair string) int {
		if pair == "AB" || pair == "BA" ||
			pair == "BC" || pair == "CB" ||
			pair == "CD" || pair == "DC" ||
			pair == "DE" || pair == "ED" ||
			pair == "EA" || pair == "AE" {
			return 1
		}
		return 2
	}

	if getNagasa(s) == getNagasa(t) {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}

}

func BPentagonReadLine(rdr *bufio.Reader) string {
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
