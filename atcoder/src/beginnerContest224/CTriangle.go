package main

import (
	"fmt"
)

type Xy struct {
	x int
	y int
}

func CTriangleMain() {
	var n int
	fmt.Scan(&n)

	var xySlice = make([]Xy, n)

	var result int64
	result = 0

	for i := 0; i < n; i++ {
		var x, y int
		fmt.Scan(&x, &y)
		xySlice[i] = Xy{
			x: x, y: y,
		}
	}
	for i := 0; i < n-2; i++ {
		iChouten := xySlice[i]
		for j := i + 1; j < n-1; j++ {
			jChouten := xySlice[j]
			zeroJosan := false
			yZero := false
			ijSyou := 0
			ijAmari := 0
			if (jChouten.x - iChouten.x) == 0 {
				zeroJosan = true
			} else if (jChouten.y - iChouten.y) == 0 {
				yZero = true
			} else {
				ijSyou = (jChouten.y - iChouten.y) / (jChouten.x - iChouten.x)
				ijAmari = (jChouten.y - iChouten.y) % (jChouten.x - iChouten.x)
			}
			for k := j + 1; k < n; k++ {
				kChouten := xySlice[k]
				if (iChouten.x-kChouten.x) == 0 && !zeroJosan {
					result = result + 1
				} else if (iChouten.y-kChouten.y) == 0 && !yZero {
					result = result + 1
				} else {
					kiSyou := (iChouten.y - kChouten.y) / (iChouten.x - kChouten.x)
					kiAmari := (iChouten.y - kChouten.y) % (iChouten.x - kChouten.x)
					if ijSyou != kiSyou || ijAmari != kiAmari {
						result = result + 1
					}
				}
			}
		}

	}
	fmt.Println(result)

}
