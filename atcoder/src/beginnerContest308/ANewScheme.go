package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func ANewSchemeMain() {
	var rdr = bufio.NewReaderSize(os.Stdin, 10000)
	sList := ANewSchemeReadLine(rdr)
	result := "Yes"
	beforeS := -1
	for _, sMoji := range strings.Split(sList, " ") {
		s, _ := strconv.Atoi(sMoji)
		if s < 100 || s > 675 {
			result = "No"
			break
		}
		if s%25 != 0 {
			result = "No"
			break
		}
		if s < beforeS {
			result = "No"
			break
		}
		beforeS = s
	}
	fmt.Println(result)

}

func ANewSchemeReadLine(rdr *bufio.Reader) string {
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
