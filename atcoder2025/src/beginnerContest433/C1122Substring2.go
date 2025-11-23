package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	var rdr = bufio.NewReaderSize(os.Stdin, 10000)
	sStrArray := strings.Split(C1122Substring2Rdr(rdr), "")
	sLen := len(sStrArray)
	var result int64 = 0

	var calcFunc func(leftCountParam int, rightCountParam int)
	calcFunc = func(leftCountParam int, rightCountParam int) {
		if leftCountParam > 0 && rightCountParam > 0 {
			if leftCountParam < rightCountParam {
				result = result + int64(leftCountParam)
			} else {
				result = result + int64(rightCountParam)
			}
		}
	}

	left := -100
	leftCount := 0
	right := -100
	rightCount := 0
	for i := 0; i < sLen; i++ {
		s, _ := strconv.Atoi(sStrArray[i])
		if left < 0 {
			left = s
			leftCount = 1
		} else if left == s && right < 0 {
			leftCount = leftCount + 1
		} else if s == left+1 && right < 0 {
			right = s
			rightCount = 1
		} else if right == s {
			rightCount = rightCount + 1
		} else {
			calcFunc(leftCount, rightCount)
			if s == right+1 {
				left = right
				leftCount = rightCount
				right = s
				rightCount = 1
			} else {
				left = s
				leftCount = 1
				right = -100
				rightCount = 0
			}
		}

		if i == sLen-1 {
			calcFunc(leftCount, rightCount)
		}
	}

	fmt.Println(result)
}

func C1122Substring2Rdr(rdr *bufio.Reader) string {
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
