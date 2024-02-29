package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	var n int
	fmt.Scan(&n)
	var rdr = bufio.NewReaderSize(os.Stdin, 10000)
	aStrArray := strings.Split(DSquarePairReadLine(rdr), " ")
	var heihouCount int64
	var result int64
	sosuuMap := map[int]int{}

	var primeFactors func(nInput int) map[int]int
	primeFactors = func(nInput int) map[int]int {
		primeFactorMap := map[int]int{}
		// Get the number of 2s that divide n
		for nInput%2 == 0 {
			v, ok := primeFactorMap[2]
			if !ok {
				primeFactorMap[2] = 1
			} else {
				primeFactorMap[2] = v + 1
			}
			nInput = nInput / 2
		}

		// n must be odd at this point. so we can skip one element
		// (note i = i + 2)
		for i := 3; i*i <= nInput; i = i + 2 {
			// while i divides n, append i and divide n
			for nInput%i == 0 {
				v, ok := primeFactorMap[i]
				if !ok {
					primeFactorMap[i] = 1
				} else {
					primeFactorMap[i] = v + 1
				}
				nInput = nInput / i
			}
		}

		// This condition is to handle the case when n is a prime number
		// greater than 2
		if nInput > 2 {
			primeFactorMap[nInput] = 1
		}

		return primeFactorMap
	}

	for _, aStr := range aStrArray {
		a, _ := strconv.Atoi(aStr)

		if a == 0 || a == 1 {
			v, ok := sosuuMap[a]
			if !ok {
				sosuuMap[a] = 1
			} else {
				sosuuMap[a] = v + 1
			}
		} else {
			// 素因数分解のmap
			aMap := primeFactors(a)
			// 奇数個のみの配列
			var kisuuSlice []int
			for k, v := range aMap {
				if v%2 != 0 {
					kisuuSlice = append(kisuuSlice, k)
				}
			}
			if len(kisuuSlice) == 0 {
				heihouCount = heihouCount + 1
			} else {
				value := 1
				for _, kisuu := range kisuuSlice {
					value = value * kisuu
				}
				v, ok := sosuuMap[value]
				if !ok {
					sosuuMap[value] = 1
				} else {
					sosuuMap[value] = v + 1
				}
			}
		}
	}

	// まずは平方数同士の数
	result += heihouCount * (heihouCount - 1) / 2
	for k, v := range sosuuMap {
		if k == 0 {
			result += int64(v * (v - 1) / 2)
			result += int64(v * (n - v))
		} else if k == 1 {
			result += int64(v * (v - 1) / 2)
			result += int64(v) * heihouCount
		} else {
			result += int64(v * (v - 1) / 2)
		}
	}
	fmt.Println(result)
}

func DSquarePairReadLine(rdr *bufio.Reader) string {
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
