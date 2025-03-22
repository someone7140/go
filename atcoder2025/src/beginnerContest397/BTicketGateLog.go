package main

import (
	"bufio"
	"fmt"
	"os"
)

func BTicketGateLogMain() {
	var rdr = bufio.NewReaderSize(os.Stdin, 10000)
	s := BTicketGateLogRdr(rdr)

	tempCount := 0
	result := 0
	for _, c := range s {
		sMoji := string([]rune{c})
		if (tempCount % 2) == 0 {
			if sMoji == "i" {
				tempCount = tempCount + 1
			} else {
				tempCount = tempCount + 2
				result = result + 1
			}
		} else {
			if sMoji == "o" {
				tempCount = tempCount + 1
			} else {
				tempCount = tempCount + 2
				result = result + 1
			}
		}
	}

	if (tempCount % 2) != 0 {
		result = result + 1
	}
	fmt.Println(result)
}

func BTicketGateLogRdr(rdr *bufio.Reader) string {
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
