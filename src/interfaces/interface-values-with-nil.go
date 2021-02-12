//多くの他の言語だと値がnilの変数がレシーバに渡されたときにエラーが起きるので例外処理を書く必要がある、しかしgoでは基本nilが渡されてもそれを利用して動いてしまう。
package main

import "fmt"

type I interface {
	M()
}

type T struct {
	S string
}

func (t *T) M() {
	if t == nil {
		fmt.Println("<nil>")
		return
	}
	fmt.Println(t.S)
}

func main() {
	var i I
//ここでｔはnilなのだがそれを代入したIはnilではないよ
	var t *T
	i = t
	describe(i)
	i.M()

	i = &T{"hello"}
	describe(i)
	i.M()
}

func describe(i I) {
	fmt.Printf("(%v, %T)\n", i, i)
}
