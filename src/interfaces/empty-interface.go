//ゼロ個のメソッドを指定したinterface型の変数は値も方も持たない
package main

import "fmt"

func main() {
	var i interface{}
	describe(i)
//柔軟に値を入れることで勝手に判断してくれる
	i = 42
	describe(i)

	i = "hello"
	describe(i)
}

func describe(i interface{}) {
	fmt.Printf("(%v, %T)\n", i, i)
}
