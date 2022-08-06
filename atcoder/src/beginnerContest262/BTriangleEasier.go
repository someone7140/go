package main

import (
	"fmt"
)

type Hen struct {
	start int
	end   int
}

func BTriangleEasierMain() {
	var n, m int
	fmt.Scan(&n, &m)

	var henArray = make([]Hen, m)

	for i := 0; i < m; i++ {
		var u, v int
		fmt.Scan(&u, &v)
		henArray[i] = Hen{
			start: u,
			end:   v,
		}
	}

	result := 0
	for i := 0; i < m-2; i++ {
		hen1 := henArray[i]
		for j := i + 1; j < m-1; j++ {
			hen2 := henArray[j]
			if hen2.start == hen1.start ||
				hen2.end == hen1.end ||
				hen2.end == hen1.start ||
				hen2.start == hen1.end {
				var targetHen Hen
				if hen2.start == hen1.start {
					if hen1.end > hen2.end {
						targetHen = Hen{
							start: hen2.end,
							end:   hen1.end,
						}
					} else {
						targetHen = Hen{
							start: hen1.end,
							end:   hen2.end,
						}
					}
				} else if hen2.end == hen1.end {
					if hen1.start > hen2.start {
						targetHen = Hen{
							start: hen2.start,
							end:   hen1.start,
						}
					} else {
						targetHen = Hen{
							start: hen1.start,
							end:   hen2.start,
						}
					}
				} else if hen2.end == hen1.start {
					targetHen = Hen{
						start: hen2.start,
						end:   hen1.end,
					}
				} else if hen2.start == hen1.end {
					targetHen = Hen{
						start: hen1.start,
						end:   hen2.end,
					}
				}

				for l := j + 1; l < m; l++ {
					hen3 := henArray[l]
					if targetHen == hen3 {
						result = result + 1
					}
				}
			}
		}
	}

	fmt.Println(result)

}
