//nilインターフェースの値を呼び出すと存在しないためランタイムエラーとなる。
package main

import "fmt"

type I interface {
	M()
}

func main() {
	//iの値に何も代入していないのでnil
	var i I
	describe(i)
	//ここでnilの値に対して呼び出しをしているのでエラーとなる
	i.M()
}

func describe(i I) {
	fmt.Printf("(%v, %T)\n", i, i)
}
