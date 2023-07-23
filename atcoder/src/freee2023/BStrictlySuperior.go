package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type BStrictlySuperiorStruct struct {
	id    int
	price int
	funcs []int
}

func BStrictlySuperiorMain() {
	var n, m int
	fmt.Scan(&n, &m)

	var rdr = bufio.NewReaderSize(os.Stdin, 10000)
	var products []BStrictlySuperiorStruct

	for i := 0; i < n; i++ {
		pcfListStr := strings.Split(BStrictlySuperiorReadLine(rdr), " ")
		price, _ := strconv.Atoi(pcfListStr[0])
		listLen := len(pcfListStr)
		var funcs []int
		for j := 2; j < listLen; j++ {
			f, _ := strconv.Atoi(pcfListStr[j])
			funcs = append(funcs, f)
		}
		products = append(products, BStrictlySuperiorStruct{
			id:    i,
			price: price,
			funcs: funcs,
		})
	}

	result := "No"
	for i, p := range products {
		for j, p2 := range products {
			if i != j {
				okFlag := false
				if p.price > p2.price {
					okFlag = true
				} else if p.price == p2.price {
					okTemp := false
					for _, f2 := range p2.funcs {
						okTemp2 := true
						for _, f1 := range p.funcs {
							if f1 == f2 {
								okTemp2 = false
								break
							}
						}
						if okTemp2 {
							okTemp = true
							break
						}
					}
					if okTemp {
						okFlag = true
					}
				}
				if okFlag {
					for _, f1 := range p.funcs {
						okTemp := false
						for _, f2 := range p2.funcs {
							if f1 == f2 {
								okTemp = true
								break
							}
						}
						if !okTemp {
							okFlag = false
							break
						}
					}
				}
				if okFlag {
					result = "Yes"
				}
			}
			if result == "Yes" {
				break
			}
		}
		if result == "Yes" {
			break
		}
	}
	fmt.Println(result)

}

func BStrictlySuperiorReadLine(rdr *bufio.Reader) string {
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
