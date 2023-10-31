package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func DPrintingMachineMain() {
	var n int
	fmt.Scan(&n)
	var rdr = bufio.NewReaderSize(os.Stdin, 10000)

	type DPrintingMachineTD struct {
		start int64
		end   int64
	}

	var nowTime int64
	nowTime = -1
	var tdSlice = make([]DPrintingMachineTD, n)
	for i := 0; i < n; i++ {
		tdStrs := strings.Split(DPrintingMachineReadLine(rdr), " ")
		t, _ := strconv.ParseInt(tdStrs[0], 10, 64)
		d, _ := strconv.ParseInt(tdStrs[1], 10, 64)
		tdSlice[i] = DPrintingMachineTD{
			start: t,
			end:   t + d,
		}
		if nowTime < 0 || nowTime > t {
			nowTime = t
		}
	}

	sort.Slice(tdSlice, func(i, j int) bool {
		if tdSlice[i].end != tdSlice[j].end {
			return tdSlice[i].end < tdSlice[j].end
		} else {
			return tdSlice[i].start < tdSlice[j].start
		}
	})

	len := n
	result := 0
	for {
		if len < 1 {
			break
		}
		// 最初の要素を取得
		first := tdSlice[0]
		if first.start > nowTime {
			nowTime = first.start
		} else {
			if first.end >= nowTime {
				nowTime = nowTime + 1
				result = result + 1
			}

			len = len - 1
			if len == 0 {
				break
			} else {
				tdSlice = tdSlice[1:]
			}
		}

	}

	fmt.Println(result)

}

func DPrintingMachineReadLine(rdr *bufio.Reader) string {
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
