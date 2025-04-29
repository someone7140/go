package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func C403ForbiddenMain() {
	var n, m, q int
	fmt.Scan(&n, &m, &q)

	var rdr = bufio.NewReaderSize(os.Stdin, 10000)
	var resultSlice []string

	allOkSet := make(map[int]struct{})
	authSetMap := map[int]map[int]struct{}{}

	for i := 0; i < q; i++ {
		qStrs := strings.Split(C403ForbiddenRdr(rdr), " ")
		if qStrs[0] == "1" {
			userID, _ := strconv.Atoi(qStrs[1])
			page, _ := strconv.Atoi(qStrs[2])
			authSet, ok := authSetMap[userID]
			if ok {
				authSet[page] = struct{}{}
				authSetMap[userID] = authSet
			} else {
				newPageSet := make(map[int]struct{})
				newPageSet[page] = struct{}{}
				authSetMap[userID] = newPageSet
			}
		} else if qStrs[0] == "2" {
			userID, _ := strconv.Atoi(qStrs[1])
			allOkSet[userID] = struct{}{}
		} else {
			userID, _ := strconv.Atoi(qStrs[1])
			page, _ := strconv.Atoi(qStrs[2])

			_, ok1 := allOkSet[userID]
			if ok1 {
				resultSlice = append(resultSlice, "Yes")
			} else {
				authSet, ok2 := authSetMap[userID]
				if ok2 {
					_, ok3 := authSet[page]
					if ok3 {
						resultSlice = append(resultSlice, "Yes")
					} else {
						resultSlice = append(resultSlice, "No")
					}
				} else {
					resultSlice = append(resultSlice, "No")
				}
			}
		}
	}
	fmt.Println(strings.Join(resultSlice, "\n"))
}

func C403ForbiddenRdr(rdr *bufio.Reader) string {
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
