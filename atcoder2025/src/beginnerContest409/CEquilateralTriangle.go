package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func CEquilateralTriangleMain() {
	var n, l int
	fmt.Scan(&n, &l)
	var rdr = bufio.NewReaderSize(os.Stdin, 10000)
	dStrs := strings.Split(CEquilateralTriangleRdr(rdr), " ")
	result := 0

	// 3で割り切れなかったら0でreturn
	if l%3 != 0 {
		fmt.Println(result)
		return
	}

	amariCountMap := map[int][]int{}
	nowCount := 0
	for i := 0; i < n; i++ {
		if i == 0 {
			amariCountMap[0] = []int{0}
		} else {
			d, _ := strconv.Atoi(dStrs[i-1])
			nowCount = (nowCount + d) % l
			points, ok := amariCountMap[nowCount]
			if ok {
				amariCountMap[nowCount] = append(points, i)
			} else {
				amariCountMap[nowCount] = []int{i}
			}
		}
	}

	checkSet := map[int]struct{}{}
	targetDistance := l / 3
	for i := 0; i < l; i++ {
		// すでにチェックしている点ならスキップ
		_, checked := checkSet[i]
		if !checked {
			point2 := (i + targetDistance) % l
			point3 := (point2 + targetDistance) % l
			_, checked2 := checkSet[point2]
			_, checked3 := checkSet[point3]

			//3点はチェック済みとする
			checkSet[i] = struct{}{}
			checkSet[point2] = struct{}{}
			checkSet[point3] = struct{}{}
			if !checked2 && !checked3 {
				// 3点の個数の掛け算
				points1, ok1 := amariCountMap[i]
				points2, ok2 := amariCountMap[point2]
				points3, ok3 := amariCountMap[point3]
				if ok1 && ok2 && ok3 {
					result = result + len(points1)*len(points2)*len(points3)
				}
			}
		}
	}
	fmt.Println(result)
}

func CEquilateralTriangleRdr(rdr *bufio.Reader) string {
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
