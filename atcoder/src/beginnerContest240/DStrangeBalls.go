package main

import (
	"bufio"
	"container/list"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func DStrangeBallsMain() {
	var n int
	fmt.Scan(&n)

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()

	aStr := scanner.Text()
	aStrArray := strings.Split(aStr, " ")

	resultSlice := make([]int, n)
	deque := list.New()
	lastNum := -9999
	renzoku := 0
	for i, v := range aStrArray {
		a, _ := strconv.Atoi(v)
		if deque.Len() != 0 {
			if lastNum == a {
				if (renzoku + 1) == a {
					for i := 0; i < renzoku; i++ {
						deque.Remove(deque.Back())
					}
					if deque.Len() != 0 {
						lastNum = deque.Back().Value.(int)
						renzoku = 1
					} else {
						lastNum = -9999
						renzoku = 0
					}
				} else {
					deque.PushBack(a)
					renzoku = renzoku + 1
				}
			} else {
				deque.PushBack(a)
				lastNum = a
				renzoku = 1
			}
		} else {
			deque.PushBack(a)
			lastNum = a
			renzoku = 1
		}
		resultSlice[i] = deque.Len()
	}

	for _, v := range resultSlice {
		fmt.Println(v)
	}
}
