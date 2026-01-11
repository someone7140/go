package main

import (
	"bufio"
	"container/list"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type C1DpuyopuyoCount struct {
	val   int
	count int
}

func C1DpuyopuyoMain() {
	var n int
	fmt.Scan(&n)
	var rdr = bufio.NewReaderSize(os.Stdin, 10000)
	aStrs := strings.Split(C1Dpuyopuyo2Rdr(rdr), " ")
	var before *C1DpuyopuyoCount
	deque := list.New()
	for i, a := range aStrs {
		aNum, _ := strconv.Atoi(a)
		if before == nil {
			before = &C1DpuyopuyoCount{
				val:   aNum,
				count: 1,
			}
		} else {
			if aNum == before.val {
				if before.count == 3 {
					before = nil
				} else {
					before = &C1DpuyopuyoCount{
						val:   aNum,
						count: before.count + 1,
					}
				}
			} else {
				if deque.Len() > 0 {
					front := deque.Front().Value.(C1DpuyopuyoCount)
					if front.val != before.val {
						deque.PushFront(*before)
					} else {
						deque.Remove(deque.Front())
						if front.count+before.count > 4 {
							deque.PushFront(C1DpuyopuyoCount{
								val:   before.val,
								count: front.count + before.count - 4,
							})
						} else if front.count+before.count < 4 {
							deque.PushFront(C1DpuyopuyoCount{
								val:   before.val,
								count: front.count + before.count,
							})
						}
					}
				} else {
					deque.PushFront(*before)
				}
				before = &C1DpuyopuyoCount{
					val:   aNum,
					count: 1,
				}
			}
		}

		if i == n-1 && before != nil {
			if deque.Len() > 0 {
				front := deque.Front().Value.(C1DpuyopuyoCount)
				if front.val != before.val {
					deque.PushFront(*before)
				} else {
					deque.Remove(deque.Front())
					if front.count+before.count > 4 {
						deque.PushFront(C1DpuyopuyoCount{
							val:   before.val,
							count: front.count + before.count - 4,
						})
					} else if front.count+before.count < 4 {
						deque.PushFront(C1DpuyopuyoCount{
							val:   before.val,
							count: front.count + before.count,
						})
					}
				}
			} else {
				deque.PushFront(*before)
			}
		}
	}

	result := 0
	for e := deque.Front(); e != nil; e = e.Next() {
		count := e.Value.(C1DpuyopuyoCount)
		result = result + count.count
	}
	fmt.Println(result)
}

func C1Dpuyopuyo2Rdr(rdr *bufio.Reader) string {
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
