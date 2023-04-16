package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	var n, q int
	fmt.Scan(&n, &q)
	var rdr = bufio.NewReaderSize(os.Stdin, 10000)

	type SortArray struct {
		sorted bool
		values []int
	}

	numMap := map[int]SortArray{}
	boxMap := map[int]SortArray{}

	numResultMap := map[int]string{}
	boxResultMap := map[int]string{}

	qArray := make([]string, q)

	for i := 0; i < q; i++ {
		qArray[i] = CCardsQueryProblemReadLine(rdr)
	}

	for i := 0; i < q; i++ {
		qStrArray := strings.Split(qArray[i], " ")
		if qStrArray[0] == "1" {
			num, _ := strconv.Atoi(qStrArray[1])
			box, _ := strconv.Atoi(qStrArray[2])

			numArr, ok := numMap[num]
			if ok {
				numMap[num] = SortArray{
					sorted: false,
					values: append(numArr.values, box),
				}
			} else {
				numMap[num] = SortArray{
					sorted: false,
					values: []int{box},
				}
			}

			boxArr, ok := boxMap[box]
			if ok {
				boxMap[box] = SortArray{
					sorted: false,
					values: append(boxArr.values, num),
				}
			} else {
				boxMap[box] = SortArray{
					sorted: false,
					values: []int{num},
				}
			}
		} else if qStrArray[0] == "2" {
			box, _ := strconv.Atoi(qStrArray[1])
			boxArr, _ := boxMap[box]
			if !boxArr.sorted {
				sort.Ints(boxArr.values)
				boxMap[box] = SortArray{
					sorted: true,
					values: boxArr.values,
				}
				var resultSlice []string

				for _, v := range boxArr.values {
					resultSlice = append(resultSlice, strconv.FormatInt(int64(v), 10))
				}
				boxResultMap[box] = strings.Join(resultSlice, " ")
				fmt.Println(boxResultMap[box])
			} else {
				fmt.Println(boxResultMap[box])
			}

		} else {
			num, _ := strconv.Atoi(qStrArray[1])
			numArr, _ := numMap[num]
			if !numArr.sorted {
				sort.Ints(numArr.values)
				var resultSlice []string
				var newSlice []int
				numSetMap := map[int]bool{}

				for _, v := range numArr.values {
					already, ok := numSetMap[v]
					if !already && !ok {
						newSlice = append(newSlice, v)
						resultSlice = append(resultSlice, strconv.FormatInt(int64(v), 10))
						numSetMap[v] = true
					}
				}
				numMap[num] = SortArray{
					sorted: true,
					values: newSlice,
				}
				numResultMap[num] = strings.Join(resultSlice, " ")
				fmt.Println(numResultMap[num])
			} else {
				fmt.Println(numResultMap[num])
			}
		}
	}
}

func CCardsQueryProblemReadLine(rdr *bufio.Reader) string {
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
