package main

import "fmt"
//ここではinterface型に新しくIという別名を付けています
type I interface {
	M()
}
//ここでは上と同じようにstruct型にTという名前をつけています。
type T struct {
	S string
}

// This method means type T implements the interface I,
// but we don't need to explicitly declare that it does so.
func (t T) M() {
	fmt.Println(t.S)
}

func main() {
	var i I = T{"hello"}
	i.M()
}
//参考サイト"https://qiita.com/Yuuki557/items/e9f5bdfbbfe92973a05e"
