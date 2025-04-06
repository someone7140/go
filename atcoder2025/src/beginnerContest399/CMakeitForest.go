package main

import (
	"fmt"
	"strconv"
)

func CMakeitForestMain() {
	var n, m int
	fmt.Scan(&n, &m)
	vMap := map[int][]int{}

	for i := 0; i < m; i++ {
		var v1, v2 int
		fmt.Scan(&v1, &v2)

		v1List, ok1 := vMap[v1]
		if ok1 {
			vMap[v1] = append(v1List, v2)
		} else {
			vMap[v1] = []int{v2}
		}

		v2List, ok2 := vMap[v2]
		if ok2 {
			vMap[v2] = append(v2List, v1)
		} else {
			vMap[v2] = []int{v1}
		}
	}

	vSet := make(map[int]struct{})
	resultSet := make(map[string]struct{})
	var loopFunc func(targetV int, from int)
	loopFunc = func(targetV int, from int) {
		vList, ok := vMap[targetV]
		if ok {
			for _, v := range vList {
				_, ok2 := vSet[v]
				if ok2 {
					if from != v && from != -1 {
						temp := ""
						if targetV < v {
							temp = strconv.FormatInt(int64(targetV), 10) + "-" + strconv.FormatInt(int64(v), 10)
						} else {
							temp = strconv.FormatInt(int64(v), 10) + "-" + strconv.FormatInt(int64(targetV), 10)
						}
						resultSet[temp] = struct{}{}
					}
				} else {
					vSet[v] = struct{}{}
					loopFunc(v, targetV)
				}
			}
		}
	}

	for i := 1; i <= n; i++ {
		_, ok := vSet[i]
		if !ok {
			vSet[i] = struct{}{}
			loopFunc(i, -1)
		}
	}

	fmt.Println(len(resultSet))

}
