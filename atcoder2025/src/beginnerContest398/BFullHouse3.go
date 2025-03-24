package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func BFullHouse3Main() {
	var rdr = bufio.NewReaderSize(os.Stdin, 10000)
	aStrs := strings.Split(BFullHouse3Rdr(rdr), " ")
	aArray := make([]int, 7)

	for i, aMoji := range aStrs {
		a, _ := strconv.Atoi(aMoji)
		aArray[i] = a
	}
	sort.Ints(aArray)

	tempCount := 0
	before := -1
	twoCountFlag := false
	threeCountFlag := false
	threeCount2Flag := false
	fourCountFlag := false
	fiveCountFlag := false
	for i := 0; i < 7; i++ {
		if before != aArray[i] {
			if tempCount == 2 {
				twoCountFlag = true
			}
			if tempCount == 3 {
				if threeCountFlag {
					threeCount2Flag = true
				} else {
					threeCountFlag = true
				}
			}
			if tempCount == 4 {
				fourCountFlag = true
			}
			if tempCount == 5 {
				fiveCountFlag = true
			}
			tempCount = 1
		} else {
			tempCount = tempCount + 1
		}
		before = aArray[i]
	}
	if tempCount == 2 {
		twoCountFlag = true
	}
	if tempCount == 3 {
		if threeCountFlag {
			threeCount2Flag = true
		} else {
			threeCountFlag = true
		}
	}
	if tempCount == 4 {
		fourCountFlag = true
	}
	if tempCount == 5 {
		fiveCountFlag = true
	}

	if (twoCountFlag || fourCountFlag || threeCount2Flag) && threeCountFlag {
		fmt.Println("Yes")
	} else if twoCountFlag && fiveCountFlag {
		fmt.Println("Yes")
	} else if twoCountFlag && fourCountFlag {
		fmt.Println("Yes")
	} else {
		fmt.Println("No")
	}

}

func BFullHouse3Rdr(rdr *bufio.Reader) string {
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
