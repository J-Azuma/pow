//xのn乗を求めるプログラム
//v = 1にxをn回掛ける処理
package main

import "fmt"

func pow(x int, n int) int {
	v := 1
	for i := 0; i < n; i++ {
		v *= x
	}
	return v
}

func main() {
	fmt.Println(pow(2, 8))
	fmt.Println(pow(2, 16))
}
