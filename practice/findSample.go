package main

import "github.com/thoas/go-funk"

func mainFindSample() {
	numbers := []int{1, 2, 3, 4}

	// Findメソッドだと返り値がinterface型
	// 見つからない場合はnil
	res1 := funk.Find(numbers, func(x int) bool {
		return x%10 == 0
	})
	if res1 == nil {
		println("nil") // nil
	}
	// 見つかった場合はキャストして値を取り出す
	res2 := funk.Find(numbers, func(x int) bool {
		return x%2 == 0
	})
	if res2 != nil {
		println(res2.(int)) // 2
	}

	// Find+型のメソッドだと結果が、二つ目の返り値に設定される
	// 見つからなかった場合は二つ目の返り値がfalse
	res3, success3 := funk.FindInt(numbers, func(x int) bool {
		return x%10 == 0
	})
	println(success3) // false
	println(res3)     // 0が設定される
	// 見つかった場合は二つ目の返り値がtrue
	res4, success4 := funk.FindInt(numbers, func(x int) bool {
		return x%2 == 0
	})
	println(success4) // true
	println(res4)     // 2が設定される
}
