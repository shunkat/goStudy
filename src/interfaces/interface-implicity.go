//go言語ではクラスと呼ばれるものが無いのでinterfaceとメソッドを使い分けることでクラスと同じように扱えます
package main

import "fmt"
//普通の言語だとinterfaceを作ったあと、それを実装するクラスを定義して、その後にどれを使うかを宣言する必要がありました。
type I interface {
	M()
}

type T struct {
	S string
}
//struct型で作った方にinterface型で定義したメソッドを入れないとその2つが結びつかないのでクラスの用には扱えません。
// This method means type T implements the interface I,
// but we don't need to explicitly declare that it does so.
func (t T) M() {
	fmt.Println(t.S)
}

func main() {
	var i I = T{"hello"}
	i.M()
}
