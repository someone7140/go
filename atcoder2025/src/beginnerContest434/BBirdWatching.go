package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type BBirdWatchingBird struct {
	id    int
	count int
	sum   int
}

func BBirdWatchingMain() {
	var n, m int
	fmt.Scan(&n, &m)
	var rdr = bufio.NewReaderSize(os.Stdin, 10000)
	birdMap := map[int]BBirdWatchingBird{}
	for i := 0; i < n; i++ {
		abStrArray := strings.Split(BBirdWatchingRdr(rdr), " ")
		a, _ := strconv.Atoi(abStrArray[0])
		b, _ := strconv.Atoi(abStrArray[1])
		bird, ok := birdMap[a]
		if ok {
			birdMap[a] = BBirdWatchingBird{
				id:    a,
				count: bird.count + 1,
				sum:   bird.sum + b,
			}
		} else {
			birdMap[a] = BBirdWatchingBird{
				id:    a,
				count: 1,
				sum:   b,
			}
		}
	}
	for i := 1; i <= m; i++ {
		bird := birdMap[i]
		fmt.Println(float64(bird.sum) / float64(bird.count))
	}
}

func BBirdWatchingRdr(rdr *bufio.Reader) string {
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
