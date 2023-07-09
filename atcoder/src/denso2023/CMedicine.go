package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type CMedicineStruct struct {
	a int64
	b int64
}

func CMedicineMain() {
	var n, k int64
	fmt.Scan(&n, &k)
	var rdr = bufio.NewReaderSize(os.Stdin, 10000)

	var sum int64
	sum = 0
	var result int64
	result = 1
	sum = 0

	abArray := make([]CMedicineStruct, n)

	var i int64
	for i = 0; i < n; i++ {
		abList := strings.Split(CMedicineRdr(rdr), " ")
		aInt, _ := strconv.Atoi(abList[0])
		bInt, _ := strconv.Atoi(abList[1])
		a := int64(aInt)
		b := int64(bInt)
		sum = sum + b
		abArray[i] = CMedicineStruct{
			a: a, b: b,
		}
	}
	sort.Slice(abArray, func(i, j int) bool { return abArray[i].a < abArray[j].a })
	for i = 0; i < n; i++ {
		if sum <= k {
			break
		} else {
			ab := abArray[i]
			result = ab.a + 1
			sum = sum - ab.b
		}
	}
	fmt.Println(result)
}

func CMedicineRdr(rdr *bufio.Reader) string {
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
