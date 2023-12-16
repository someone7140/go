package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func CTshirtsMain() {
	var n, m int
	fmt.Scan(&n, &m)
	var rdr = bufio.NewReaderSize(os.Stdin, 10000)
	s := CTshirtsReadLine(rdr)

	zanMaisuu1 := m
	tempMaisuu2 := 0
	result := 0
	for i, sMoji := range strings.Split(s, "") {
		if sMoji == "0" {
			zanMaisuu1 = m
			if result < tempMaisuu2 {
				result = tempMaisuu2
			}
			tempMaisuu2 = 0
		} else if sMoji == "1" {
			if zanMaisuu1 > 0 {
				zanMaisuu1 = zanMaisuu1 - 1
			} else {
				tempMaisuu2 = tempMaisuu2 + 1
			}
		} else {
			tempMaisuu2 = tempMaisuu2 + 1
		}

		if i == n-1 {
			if result < tempMaisuu2 {
				result = tempMaisuu2
			}
		}
	}

	fmt.Println(result)

}

func CTshirtsReadLine(rdr *bufio.Reader) string {
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
