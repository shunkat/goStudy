//型アサーションはインターフェースの値のもとになる具体的な値を利用する手段を提供します。
//まずt := i.(T)という書き方なのですが、これの意味は変数tの中にiの変数を入れる、そして大文字のT型として入れるという意味
//このときにiがTの方を保持できないとエラーが起きる。
//そいつを防ぐために値を渡すときにokというやつも含めて2つの引数を渡します。
//それをすると
package main

import "fmt"

func main() {
	var i interface{} = "hello"

	s := i.(string)
	fmt.Println(s)
//okという値も渡すことで、値が保持できているときにはokの中にはtrueがそうでなければ
	s, ok := i.(string)
	fmt.Println(s, ok)

	f, ok := i.(float64)
	fmt.Println(f, ok)

	f = i.(float64) // panic
	fmt.Println(f)
}
