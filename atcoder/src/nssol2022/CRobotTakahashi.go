package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func CRobotTakahashiMain() {
	var n int
	fmt.Scan(&n)

	var rdr = bufio.NewReaderSize(os.Stdin, 10000)

	sAdultMap := map[int]int{}
	sChildMap := map[int]int{}
	adultCount := 0
	childCount := 0

	s := CRobotTakahashiRdr(rdr)

	wStrArray := strings.Split(CRobotTakahashiRdr(rdr), " ")
	wArray := make([]int, n)
	for i := 0; i < n; i++ {
		wArray[i], _ = strconv.Atoi(wStrArray[i])
	}

	for i, c := range s {
		w := wArray[i]
		sMoji := string([]rune{c})
		if sMoji == "1" {
			v, ok := sAdultMap[w]
			if ok {
				sAdultMap[w] = v + 1
			} else {
				sAdultMap[w] = 1
			}
			adultCount = adultCount + 1
		} else {
			v, ok := sChildMap[w]
			if ok {
				sChildMap[w] = v + 1
			} else {
				sChildMap[w] = 1
			}
			childCount = childCount + 1
		}
	}
	uniqueArray := CRobotTakahashiUnique(wArray)
	sort.Ints(uniqueArray)
	result := adultCount
	tempResult := adultCount

	for _, w := range uniqueArray {

		wAdultCount := 0
		getAdultMap, ok := sAdultMap[w]
		if ok {
			wAdultCount = getAdultMap
		}

		wChildCount := 0
		geChildMap, ok := sChildMap[w]
		if ok {
			wChildCount = geChildMap
		}
		tempResult = tempResult - wAdultCount + wChildCount
		if result < tempResult {
			result = tempResult
		}
	}

	fmt.Println(result)
}

func CRobotTakahashiRdr(rdr *bufio.Reader) string {
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

func CRobotTakahashiUnique(intSlice []int) []int {
	keys := make(map[int]bool)
	list := []int{}
	for _, entry := range intSlice {
		if _, value := keys[entry]; !value {
			keys[entry] = true
			list = append(list, entry)
		}
	}
	return list
}
