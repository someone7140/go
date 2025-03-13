package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func CSnakeQueueMain() {
	var q int
	fmt.Scan(&q)
	var resultSlice []string
	var snakeSlice []int
	var snakeSumSlice []int
	startIndex := -1

	var rdr = bufio.NewReaderSize(os.Stdin, 10000)
	for i := 0; i < q; i++ {
		query := CSnakeQueueRdr(rdr)
		queryStrArray := strings.Split(query, " ")

		if queryStrArray[0] == "1" {
			snakeLen, _ := strconv.Atoi(queryStrArray[1])
			snakeSlice = append(snakeSlice, snakeLen)
			lenSum := len(snakeSumSlice)
			if lenSum == 0 {
				snakeSumSlice = append(snakeSumSlice, snakeLen)
			} else {
				snakeSumSlice = append(snakeSumSlice, snakeSumSlice[lenSum-1]+snakeLen)
			}
		} else if queryStrArray[0] == "2" {
			startIndex = startIndex + 1
		} else {
			target, _ := strconv.Atoi(queryStrArray[1])
			if target == 1 {
				resultSlice = append(resultSlice, "0")
			} else {
				target = target - 2
				if startIndex == -1 {
					resultSlice = append(resultSlice, strconv.FormatInt(int64(snakeSumSlice[target]), 10))
				} else {
					startVal := snakeSumSlice[startIndex]
					endVal := snakeSumSlice[startIndex+target+1]
					resultSlice = append(resultSlice, strconv.FormatInt(int64(endVal-startVal), 10))
				}
			}
		}
	}

	fmt.Println(strings.Join(resultSlice, "\n"))

}

func CSnakeQueueRdr(rdr *bufio.Reader) string {
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
