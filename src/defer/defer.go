//deferで指定した関数をその関数の呼び出し元の関数が終わったあとに実行することができます。
//すぐにその引数は評価されますが、実行されるのは最後になります。
package main

import "fmt"

func main() {
	defer fmt.Println("world")

	fmt.Println("hello")

}
