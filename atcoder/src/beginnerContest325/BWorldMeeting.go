package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func BWorldMeetingMain() {
	var n int
	fmt.Scan(&n)
	var rdr = bufio.NewReaderSize(os.Stdin, 10000)

	type BWorldMeetingWX struct {
		w int
		x int
	}
	var wxSlice = make([]BWorldMeetingWX, n)
	for i := 0; i < n; i++ {
		wxStrs := strings.Split(BWorldMeetingReadLine(rdr), " ")
		w, _ := strconv.Atoi(wxStrs[0])
		x, _ := strconv.Atoi(wxStrs[1])
		wxSlice[i] = BWorldMeetingWX{
			w: w,
			x: x,
		}
	}

	result := -1
	for i := 0; i < 24; i++ {
		tempResult := 0
		for j := 0; j < n; j++ {
			w := wxSlice[j].w
			x := wxSlice[j].x

			jikan := (i + x) % 24
			if jikan >= 9 && jikan < 18 {
				tempResult = tempResult + w
			}
		}
		if tempResult > result {
			result = tempResult
		}
	}

	fmt.Println(result)

}

func BWorldMeetingReadLine(rdr *bufio.Reader) string {
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
