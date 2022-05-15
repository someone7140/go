package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func BAtMostMain() {
	var n, w int
	fmt.Scan(&n, &w)

	var rdr = bufio.NewReaderSize(os.Stdin, 10000)

	mapA := map[int]int{}

	aLine := BAtMostReadLine(rdr)
	aStrArray := strings.Split(aLine, " ")
	var aArray []int
	for i := 0; i < n; i++ {
		a, _ := strconv.Atoi(aStrArray[i])
		if a <= w {
			aArray = append(aArray, a)
		}

	}
	sort.Slice(aArray, func(i, j int) bool { return aArray[i] < aArray[j] })
	aSize := len(aArray)

	if aSize == 1 {
		a := aArray[0]
		if a <= w {
			_, ok := mapA[a]
			if !ok {
				mapA[a] = 1
			}
		}
	} else if aSize == 2 {
		for i := 0; i < 2; i++ {
			a1 := aArray[i]
			if a1 <= w {
				_, ok := mapA[a1]
				if !ok {
					mapA[a1] = 1
				}
			}
			if i == 0 {
				a2 := aArray[1]
				sum := a1 + a2
				if sum <= w {
					_, ok := mapA[sum]
					if !ok {
						mapA[sum] = 1
					}
				}
			}
		}
	} else {
		for i := 0; i < aSize; i++ {
			a1 := aArray[i]
			if a1 <= w {
				_, ok := mapA[a1]
				if !ok {
					mapA[a1] = 1
				}
			} else {
				break
			}
			if i != n-1 && a1 < w {
				for j := i + 1; j < n; j++ {
					a2 := aArray[j]
					sum12 := a1 + a2
					if sum12 <= w {
						_, ok := mapA[sum12]
						if !ok {
							mapA[sum12] = 1
						}
					} else {
						break
					}
					if j != n-1 && sum12 < w {
						for l := j + 1; l < n; l++ {
							a3 := aArray[l]
							sum123 := sum12 + a3
							if sum123 <= w {
								_, ok := mapA[sum123]
								if !ok {
									mapA[sum123] = 1
								}
							} else {
								break
							}
						}
					}
				}
			}
		}
	}

	fmt.Println(len(mapA))
}

func BAtMostReadLine(rdr *bufio.Reader) string {
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
