//変数に初期値を与えずに宣言するとゼロ値が与えられます。
//ゼロの値はintとかだと０，bool型だとfalse,string型だと""が与えられます。
package main

import "fmt"

func main() {
	var i int
	var f float64
	var b bool
	var s string
	//ちなみにここで%q\nという形でフォーマットしているのは空文字がそのままだとわかりにくいから""で囲む形で表現したかったからです。
	fmt.Printf("%v %v %v %q\n", i, f, b, s)
}

