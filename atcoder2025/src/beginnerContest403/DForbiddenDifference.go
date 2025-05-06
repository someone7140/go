package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func DForbiddenDifferenceMain() {
	var n, d int
	fmt.Scan(&n, &d)

	var rdr = bufio.NewReaderSize(os.Stdin, 10000)
	qSlice := make([]int, n)
	qStrs := strings.Split(DForbiddenDifferenceRdr(rdr), " ")
	countMap := make(map[int]int)
	for i := 0; i < n; i++ {
		q, _ := strconv.Atoi(qStrs[i])
		qSlice[i] = q
		count, ok := countMap[q]
		if ok {
			countMap[q] = count + 1
		} else {
			countMap[q] = 1
		}
	}

	keys := make([]int, 0, len(countMap))
	for k := range countMap {
		keys = append(keys, k)
	}
	sort.Ints(keys)
	deletedSet := make(map[int]struct{})
	result := 0
	for _, key := range keys {
		count := countMap[key]
		_, targetDelete := deletedSet[key]
		if !targetDelete {
			before, ok1 := countMap[key-d]
			_, beforeDelete := deletedSet[key-d]
			after, ok2 := countMap[key+d]
			_, afterDelete := deletedSet[key+d]

			if ok1 && !beforeDelete && ok2 && !afterDelete {
				if before+after > count {
					deletedSet[key] = struct{}{}
					result = result + count
				} else {
					deletedSet[key-d] = struct{}{}
					deletedSet[key+d] = struct{}{}
					result = result + before + after
				}
				continue
			}
			if ok1 && !beforeDelete {
				if before > count {
					deletedSet[key] = struct{}{}
					result = result + count
				} else {
					deletedSet[key-d] = struct{}{}
					result = result + before
				}
				continue
			}
			if ok2 && !afterDelete {
				if after < count {
					deletedSet[key+d] = struct{}{}
					result = result + after
				} else {
					deletedSet[key] = struct{}{}
					result = result + count
				}
				continue
			}
		}

	}
	fmt.Println(result)
}

func DForbiddenDifferenceRdr(rdr *bufio.Reader) string {
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
