package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func BFindPermutationMain() {
	var n int
	fmt.Scan(&n)
	var rdr = bufio.NewReaderSize(os.Stdin, 10000)
	aStrArray := strings.Split(BFindPermutationRdr(rdr), " ")
	aArray := make([]int, n)
	for i := 0; i < n; i++ {
		a, _ := strconv.Atoi(aStrArray[i])
		aArray[i] = a
	}

	result := "Yes"
	aSet := make(map[int]struct{})

	for i := 0; i < n; i++ {
		if aArray[i] > 0 {
			_, ok := aSet[aArray[i]]
			if ok {
				result = "No"
				break
			} else {
				aSet[aArray[i]] = struct{}{}
			}
		}
	}

	if result == "No" {
		fmt.Println(result)
	} else {
		resultList := make([]string, n)
		aSet2 := make(map[int]struct{})
		for i := 0; i < n; i++ {
			if aArray[i] > 0 {
				resultList[i] = strconv.FormatInt(int64(aArray[i]), 10)
			} else {
				temp := -1
				for j := 1; j <= n; j++ {
					_, ok := aSet[j]
					_, ok2 := aSet2[j]
					if !ok && !ok2 {
						aSet2[j] = struct{}{}
						temp = j
						break
					}
				}
				resultList[i] = strconv.FormatInt(int64(temp), 10)
			}
		}
		fmt.Println(result)
		fmt.Println(strings.Join(resultList, " "))
	}
}

func BFindPermutationRdr(rdr *bufio.Reader) string {
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
