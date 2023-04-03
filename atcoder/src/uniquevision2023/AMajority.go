package main

import (
	"bufio"
	"fmt"
	"os"
)

func AMajorityMain() {
	var n int
	fmt.Scan(&n)

	var rdr = bufio.NewReaderSize(os.Stdin, 10000)
	sansei := 0
	for i := 0; i < n; i++ {
		s := AMajorityReadLine(rdr)
		if s == "For" {
			sansei = sansei + 1
		}
	}

	result := "No"
	hantai := n - sansei
	if sansei > hantai {
		result = "Yes"
	}
	fmt.Println(result)
}

func AMajorityReadLine(rdr *bufio.Reader) string {
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
