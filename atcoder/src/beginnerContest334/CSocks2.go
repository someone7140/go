package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func CSocks2Main() {
	var n, k int
	fmt.Scan(&n, &k)

	var rdr = bufio.NewReaderSize(os.Stdin, 10000)
	aStrs := strings.Split(CSocks2ReadLine(rdr), " ")
	lenA := len(aStrs)

	if lenA < 2 {
		fmt.Println(0)
	} else {
		// 差分値を記録
		var diffSlice []int64
		var diffTotal int64
		for i := 0; i < k; i++ {
			target, _ := strconv.Atoi(aStrs[i])
			if i < k-1 {
				after, _ := strconv.Atoi(aStrs[i+1])
				diffAfter := int64(after - target)
				diffSlice = append(diffSlice, diffAfter)
				diffTotal = diffTotal + int64(diffAfter)
			}
		}
		result := diffTotal
		diffLen := len(diffSlice)
		if lenA > 2 {
			if lenA%2 == 0 {
				// どこで切り離すか
				for i := 1; i < diffLen-1; i = i + 2 {
					tempResult := diffTotal - diffSlice[i]
					if result > tempResult {
						result = tempResult
					}
				}
			} else {
				// 奇数の時はどれを外すか判定
				for i := 0; i < k; i++ {
					tempResult := diffTotal
					target, _ := strconv.Atoi(aStrs[i])
					if i == 0 {
						after, _ := strconv.Atoi(aStrs[i+1])
						diffAfter := int64(after - target)
						tempResult1 := diffTotal - diffAfter
						// どこで切り離すか
						for j := 2; j < diffLen-1; j = j + 2 {
							tempResult2 := tempResult1 - diffSlice[j]
							if tempResult1 > tempResult2 {
								tempResult1 = tempResult2
							}
						}
						if result > tempResult1 {
							result = tempResult1
						}
					} else if i == k-1 {
						before, _ := strconv.Atoi(aStrs[i-1])
						diffBefore := int64(target - before)
						tempResult1 := diffTotal - diffBefore
						// どこで切り離すか
						for j := 1; j < diffLen-2; j = j + 2 {
							tempResult2 := tempResult1 - diffSlice[j]
							if tempResult1 > tempResult2 {
								tempResult1 = tempResult2
							}
						}
						if result > tempResult1 {
							result = tempResult1
						}
					} else if i%2 == 0 {
						after, _ := strconv.Atoi(aStrs[i+1])
						diffAfter := int64(after - target)
						tempResult = tempResult - diffAfter
						before, _ := strconv.Atoi(aStrs[i-1])
						diffBefore := int64(target - before)
						tempResult = tempResult - diffBefore
					}
					if result > tempResult {
						result = tempResult
					}
				}
			}
		}
		fmt.Println(result)
	}
}

func CSocks2ReadLine(rdr *bufio.Reader) string {
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
