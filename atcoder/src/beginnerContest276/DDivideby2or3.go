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
	aStrArray := strings.Split(DDivideby2or3ReadLine(rdr), " ")
	setA := make(map[int]struct{})

	for i := 0; i < n; i++ {
		a, _ := strconv.Atoi(aStrArray[i])
		setA[a] = struct{}{}
	}

	result := 0
	for {
		if len(setA) > 1 {
			baisuu2 := 0
			baisuu3 := 0
			baisuuIgai := 0
			newSetA := make(map[int]struct{})
			for k := range setA {
				if k%2 == 0 {
					baisuu2 = baisuu2 + 1
				} else if k%3 == 0 {
					baisuu3 = baisuu3 + 1
				} else {
					baisuuIgai = baisuuIgai + 1
				}
			}

			if baisuuIgai > 1 {
				result = -1
				break
			} else {
				if baisuu3 < baisuu2 {
					for k := range setA {
						if k%2 == 0 {
							newSetA[k/2] = struct{}{}
						} else {
							newSetA[k] = struct{}{}
						}
					}
				} else {
					for k := range setA {
						if k%3 == 0 {
							newSetA[k/3] = struct{}{}
						} else {
							newSetA[k] = struct{}{}
						}
					}
				}
				result = result + 1
				setA = newSetA
			}
		} else {
			break
		}
	}

	fmt.Println(result)
}

func DDivideby2or3ReadLine(rdr *bufio.Reader) string {
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
