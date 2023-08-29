package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type DLOWERQuery struct {
	sousa int
	index int
	moji  string
}

func DLOWERMain() {
	var n int
	fmt.Scan(&n)

	var rdr = bufio.NewReaderSize(os.Stdin, 10000)
	s := DLOWERReadLine(rdr)
	sArray := make([]string, n)
	for i, c := range s {
		sMoji := string([]rune{c})
		sArray[i] = sMoji
	}

	q, _ := strconv.Atoi(DLOWERReadLine(rdr))
	queryArray := make([]DLOWERQuery, q)

	lastOkikaeIndex := -1

	for i := 0; i < q; i++ {
		queies := strings.Split(DLOWERReadLine(rdr), " ")
		sousa, _ := strconv.Atoi(queies[0])
		index, _ := strconv.Atoi(queies[1])
		moji := queies[2]
		queryArray[i] = DLOWERQuery{
			sousa: sousa,
			index: index - 1,
			moji:  moji,
		}
		if sousa == 2 || sousa == 3 {
			lastOkikaeIndex = i
		}
	}

	for i := 0; i < q; i++ {
		query := queryArray[i]
		if query.sousa == 1 {
			sArray[query.index] = query.moji
		} else {
			if i == lastOkikaeIndex {
				if query.sousa == 2 {
					for j := 0; j < n; j++ {
						sArray[j] = strings.ToLower(sArray[j])
					}
				} else {
					for j := 0; j < n; j++ {
						sArray[j] = strings.ToUpper(sArray[j])
					}
				}
			}
		}
	}

	fmt.Println(strings.Join(sArray, ""))
}

func DLOWERReadLine(rdr *bufio.Reader) string {
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
