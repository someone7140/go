package main

import (
	"fmt"
)

type St struct {
	s string
	t string
}

func BUniqueNicknamesMain() {
	var n int
	fmt.Scan(&n)
	var stSlice = make([]St, n)
	stMap := map[string]int{}

	for i := 0; i < n; i++ {
		var s, t string
		fmt.Scan(&s, &t)
		stSlice[i] = St{
			s: s, t: t,
		}
		vs, ok := stMap[s]
		if ok {
			stMap[s] = 1 + vs
		} else {
			stMap[s] = 1
		}
		if s != t {
			vt, ok := stMap[t]
			if ok {
				stMap[t] = 1 + vt
			} else {
				stMap[t] = 1
			}
		}
	}
	result := "Yes"
	for i := 0; i < n; i++ {
		st := stSlice[i]
		vs, ok := stMap[st.s]
		if ok && vs > 1 {
			vt, ok := stMap[st.t]
			if ok && vt > 1 {
				result = "No"
			}
		}
	}
	fmt.Println(result)
}
