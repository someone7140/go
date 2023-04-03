package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func BYellowandRedCardMain() {
	var n, q int
	fmt.Scan(&n, &q)

	memberArray := make([]int, n)
	var rdr = bufio.NewReaderSize(os.Stdin, 10000)
	var resultSlice []string

	for i := 0; i < q; i++ {
		eventStrArray := strings.Split(BYellowandRedCardReadLine(rdr), " ")
		event, _ := strconv.Atoi(eventStrArray[0])
		member, _ := strconv.Atoi(eventStrArray[1])
		if event == 1 {
			memberArray[member-1] = memberArray[member-1] + 1
		} else if event == 2 {
			memberArray[member-1] = memberArray[member-1] + 2
		} else {
			status := memberArray[member-1]
			if status == 0 || status == 1 {
				resultSlice = append(resultSlice, "No")
			} else {
				resultSlice = append(resultSlice, "Yes")
			}
		}
	}

	fmt.Println(strings.Join(resultSlice, "\n"))

}

func BYellowandRedCardReadLine(rdr *bufio.Reader) string {
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
