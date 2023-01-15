package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func DHappyNewYear2023Main() {
	var t int
	fmt.Scan(&t)

	resultArray := make([]string, t)

	var rdr = bufio.NewReaderSize(os.Stdin, 10000)

	for i := 0; i < t; i++ {
		n, _ := strconv.ParseInt(DHappyNewYear2023ReadLine(rdr), 10, 64)
		last := int64(-1)
		p := int64(-1)
		soinsuuF := func(n int64) {
			var pfs []int64
			// Get the number of 2s that divide n
			for n%2 == 0 {
				if last == 2 {
					p = 2
					return
				}
				pfs = append(pfs, 2)
				n = n / 2
				last = 2
			}

			// n must be odd at this point. so we can skip one element
			// (note i = i + 2)
			var i int64
			for i = 3; i*i <= n; i = i + 2 {
				// while i divides n, append i and divide n
				for n%i == 0 {
					if last == i {
						p = i
						return
					}
					pfs = append(pfs, i)
					n = n / i
					last = i
				}
			}

			/*
				// This condition is to handle the case when n is a prime number
				// greater than 2
				if n > 2 {
					pfs = append(pfs, n)
				}

				return
			*/
		}
		soinsuuF(n)
		q := n / (p * p)
		resultArray[i] = strconv.FormatInt(p, 10) + " " + strconv.FormatInt(q, 10)
	}

	fmt.Println(strings.Join(resultArray, "\n"))

}

func DHappyNewYear2023ReadLine(rdr *bufio.Reader) string {
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
