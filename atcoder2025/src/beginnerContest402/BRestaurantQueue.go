package main

import (
	"bufio"
	"container/list"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func BRestaurantQueueMain() {
	var q int
	fmt.Scan(&q)
	var resultSlice []string
	var rdr = bufio.NewReaderSize(os.Stdin, 10000)
	deque := list.New()

	for i := 0; i < q; i++ {
		queries := strings.Split(BRestaurantQueueRdr(rdr), " ")
		if queries[0] == "1" {
			val, _ := strconv.Atoi(queries[1])
			deque.PushBack(val)

		} else {
			front := deque.Front().Value.(int)
			resultSlice = append(resultSlice, strconv.FormatInt(int64(front), 10))
			deque.Remove(deque.Front())
		}
	}

	fmt.Println(strings.Join(resultSlice, "\n"))

}

func BRestaurantQueueRdr(rdr *bufio.Reader) string {
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
