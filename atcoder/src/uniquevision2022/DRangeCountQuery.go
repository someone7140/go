package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	var n int
	fmt.Scan(&n)

	var rdr = bufio.NewReaderSize(os.Stdin, 10000)
	aStrSlice := strings.Split(readLine(rdr), " ")
	q, _ := strconv.Atoi(readLine(rdr))
	var resultSlice = make([]string, q)

	aMap := map[int][]int{}
	for i := 0; i < n; i++ {
		a, _ := strconv.Atoi(aStrSlice[i])

		aMapValue, ok := aMap[a]
		if ok {
			aMap[a] = append(aMapValue, i+1)
		} else {
			aMap[a] = []int{i + 1}
		}
	}

	for i := 0; i < q; i++ {
		qStrSlice := strings.Split(readLine(rdr), " ")
		l, _ := strconv.Atoi(qStrSlice[0])
		r, _ := strconv.Atoi(qStrSlice[1])
		x, _ := strconv.Atoi(qStrSlice[2])
		xArray, ok := aMap[x]
		if ok {
			xArrayLen := len(xArray)
			startIndex := -1
			lastIndex := -1
			for j := 0; j < xArrayLen; j++ {
				if startIndex != -1 && lastIndex != -1 {
					break
				}
				if startIndex == -1 {
					if l <= xArray[j] {
						startIndex = j
					}
				}
				if lastIndex == -1 {
					if r >= xArray[xArrayLen-1-j] {
						lastIndex = xArrayLen - 1 - j
					}
				}
			}
			if startIndex == -1 || lastIndex == -1 {
				resultSlice[i] = "0"
			} else {
				tempResult := xArrayLen - (startIndex) - (xArrayLen - (lastIndex + 1))
				resultSlice[i] = strconv.FormatInt(int64(tempResult), 10)
			}

		} else {
			resultSlice[i] = "0"
		}
	}

	fmt.Println(strings.Join(resultSlice, "\n"))

}

func readLine(rdr *bufio.Reader) string {
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
