package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func DUnicyclicComponentsMain() {
	var n, m int
	fmt.Scan(&n, &m)

	union := DUnicyclicComponentsNewUnionFind(n)

	var rdr = bufio.NewReaderSize(os.Stdin, 10000)

	var uvSlice []string
	for i := 0; i < m; i++ {
		uvStr := DUnicyclicComponentsReadLine(rdr)
		uvSlice = append(uvSlice, uvStr)
		uvStrArray := strings.Split(uvStr, " ")
		u, _ := strconv.Atoi(uvStrArray[0])
		v, _ := strconv.Atoi(uvStrArray[1])
		u = u - 1
		v = v - 1
		if u != v && !union.DUnicyclicComponentsSame(u, v) {
			union.DUnicyclicComponentsUnite(u, v)
		}
	}

	henMap := map[int]int{}
	for _, uv := range uvSlice {
		uvStrArray := strings.Split(uv, " ")
		u, _ := strconv.Atoi(uvStrArray[0])
		u = u - 1
		root := union.DUnicyclicComponentsRoot(u)
		count, ok := henMap[root]
		if ok {
			henMap[root] = count + 1
		} else {
			henMap[root] = 1
		}
	}

	result := "Yes"
	// グループ毎の木の大きさ
	for i := 0; i < n; i++ {
		size := union.DUnicyclicComponentsSize(i)
		hen := henMap[union.DUnicyclicComponentsRoot(i)]
		if size != hen {
			result = "No"
			break
		}
	}
	fmt.Println(result)
}

func DUnicyclicComponentsReadLine(rdr *bufio.Reader) string {
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

type DUnicyclicComponentsUnionFind struct {
	par []int
}

/* コンストラクタ */
func DUnicyclicComponentsNewUnionFind(N int) *DUnicyclicComponentsUnionFind {
	u := new(DUnicyclicComponentsUnionFind)
	u.par = make([]int, N)
	for i := range u.par {
		u.par[i] = -1
	}
	return u
}

/* xの所属するグループを返す */
func (u DUnicyclicComponentsUnionFind) DUnicyclicComponentsRoot(x int) int {
	if u.par[x] < 0 {
		return x
	}
	u.par[x] = u.DUnicyclicComponentsRoot(u.par[x])
	return u.par[x]
}

/* xの所属するグループ と yの所属するグループ を合体する */
func (u DUnicyclicComponentsUnionFind) DUnicyclicComponentsUnite(x, y int) {
	xr := u.DUnicyclicComponentsRoot(x)
	yr := u.DUnicyclicComponentsRoot(y)
	if xr == yr {
		return
	}
	u.par[yr] += u.par[xr]
	u.par[xr] = yr
}

/* xとyが同じグループに所属するかどうかを返す */
func (u DUnicyclicComponentsUnionFind) DUnicyclicComponentsSame(x, y int) bool {
	if u.DUnicyclicComponentsRoot(x) == u.DUnicyclicComponentsRoot(y) {
		return true
	}
	return false
}

/* xの所属するグループの木の大きさを返す */
func (u DUnicyclicComponentsUnionFind) DUnicyclicComponentsSize(x int) int {
	return -u.par[u.DUnicyclicComponentsRoot(x)]
}
