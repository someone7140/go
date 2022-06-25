package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var DJumpingTakahashiResult = 0

type XypJump struct {
	x int
	y int
	p int
}

func main() {
	var n int
	fmt.Scan(&n)

	var rdr = bufio.NewReaderSize(os.Stdin, 10000)
	xypArray := make([]XypJump, n)
	renbanArray := make([]int, n)

	for i := 0; i < n; i++ {
		xyp := strings.Split(DJumpingTakahashi2Rdr(rdr), " ")
		x, _ := strconv.Atoi(xyp[0])
		y, _ := strconv.Atoi(xyp[1])
		p, _ := strconv.Atoi(xyp[2])
		xypArray[i] = XypJump{
			x: x,
			y: y,
			p: p,
		}
		renbanArray[i] = i
	}
}

func DJumpingTakahashi2Rdr(rdr *bufio.Reader) string {
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
