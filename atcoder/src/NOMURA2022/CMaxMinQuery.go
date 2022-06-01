package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func CMaxMinQueryMain() {
	var q int
	fmt.Scan(&q)
	var rdr = bufio.NewReaderSize(os.Stdin, 10000)

	sMap := map[int]int{}

	sMax := -1
	sMin := -1

	var resultSlice []string
	for i := 0; i < q; i++ {
		queryStr := CMaxMinQueryRdr(rdr)
		queryArray := strings.Split(queryStr, " ")
		queryType := queryArray[0]

		if queryType == "1" {
			insertS, _ := strconv.Atoi(queryArray[1])
			if sMax == -1 || sMax < insertS {
				sMax = insertS
			}
			if sMin == -1 || sMin > insertS {
				sMin = insertS
			}
			v, ok := sMap[insertS]
			if ok {
				sMap[insertS] = v + 1
			} else {
				sMap[insertS] = 1
			}
		} else if queryType == "2" {
			deleteS, _ := strconv.Atoi(queryArray[1])
			deleteCount, _ := strconv.Atoi(queryArray[2])
			v, ok := sMap[deleteS]
			if ok {
				if v <= deleteCount {
					delete(sMap, deleteS)
					if deleteS == sMax {
						tempSMax := -1
						for k := range sMap {
							if tempSMax == -1 || k > tempSMax {
								tempSMax = k
							}
						}
						sMax = tempSMax
					}
					if deleteS == sMin {
						tempSMin := -1
						for k := range sMap {
							if tempSMin == -1 || k < tempSMin {
								tempSMin = k
							}
						}
						sMin = tempSMin
					}
				} else {
					sMap[deleteS] = v - deleteCount
				}
			}
		} else {
			resultSlice = append(resultSlice, strconv.FormatInt(int64(sMax-sMin), 10))
		}
	}

	fmt.Println(strings.Join(resultSlice, "\n"))
}

func CMaxMinQueryRdr(rdr *bufio.Reader) string {
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
