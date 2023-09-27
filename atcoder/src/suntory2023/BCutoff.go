package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func BCutoffMain() {
	var n, x int
	fmt.Scan(&n, &x)

	var rdr = bufio.NewReaderSize(os.Stdin, 10000)
	aStrArray := strings.Split(BCutoffReadLine(rdr), " ")

	var aArray []int
	for i := 0; i < n-1; i++ {
		a, _ := strconv.Atoi(aStrArray[i])
		aArray = append(aArray, a)
	}
	sort.Ints(aArray)

	if len(aArray) > 2 {
		sum := 0
		for i := 0; i < n-1; i++ {
			if i != 0 && i != n-2 {
				sum += aArray[i]
			}
		}
		if sum >= x {
			fmt.Println(0)
		} else {
			result := x - sum
			if result <= 100 {
				max := aArray[n-2]
				if max < result {
					fmt.Println(-1)
				} else {
					min := aArray[0]
					if result <= min {
						fmt.Println(0)
					} else {
						fmt.Println(result)
					}
				}
			} else {
				fmt.Println(-1)
			}
		}
	} else {
		if x > aArray[1] {
			fmt.Println(-1)
		} else if x <= aArray[0] {
			fmt.Println(0)
		} else {
			fmt.Println(x)
		}
	}
}

func BCutoffReadLine(rdr *bufio.Reader) string {
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
