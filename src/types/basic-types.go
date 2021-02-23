//基本的な型はbool(真偽値),string(文字列),int(数字だがその値の取れる範囲によって色々な型が存在する),byte(uint8のこと),rune,float64(小数点が存在する数値),complex
package main

import (
	"fmt"
	"math/cmplx"
)

var (
	//変数名　型　＝　値　という書き方をする。
	ToBe   bool       = false
	MaxInt uint64     = 1<<64 - 1
	z      complex128 = cmplx.Sqrt(-5 + 12i)
)

func main() {
	//％を使うことでその変数の型を見ることができる。
	//%Tでその型、%v\nでその値を取れる。
	fmt.Printf("Type: %T Value: %v\n", ToBe, ToBe)
	fmt.Printf("Type: %T Value: %v\n", MaxInt, MaxInt)
	fmt.Printf("Type: %T Value: %v\n", z, z)
}

//ここでちょっとわからなくなったので補足
//fmtパッケージの中にあるprint系の関数はいくつか種類がある。
//Print("Hello World")だとそのまま出力
//Fprint,Fprintf,Fprintlnだと書き込み先を毎時的に指定することができる。fmt.Fprint(os.Stdout,"Hello world!")
//Sprint()だと出力ではなくフォーマットした結果を文字列で返します。変数への代入で使うそうです。え、普通に代入するのではだめなのか？hello:=fmt.Sprint("Hello World")<br>fmt.Print(hello)
//Printfはフォーマットを自分で指定することができます。fmt.Printf("%s\n",hello)やfmt.Printf("%T,hello)のような形です。コレで上のやつが理解できましたね。
//Printlnはオペランドの間に半角スペースが入り、文字列の最後に改行が追加されます。fmt.Println("Hello","world") -> Hello world<br>のような感じです。
//以上

