package main

import "fmt"
//再三にはなりますがここでstruct型の変数vertexを宣言しています。
type Vertex struct {
	X, Y int
}
//こうやってまとめて変数を宣言することもできるんですね。
var (
	//v1には値を全部入れて、v2には全然値を入れていないのですが、それでもv2の出力のときにはy=0が勝手に代入されています。コレは初期値が代入されているからなんですね。
	//v3も同じように何の値を入れていないので初期値の０が２つの変数の両方に代入されています。
	v1 = Vertex{1, 2}  // has type Vertex
	v2 = Vertex{X: 1}  // Y:0 is implicit
	v3 = Vertex{}      // X:0 and Y:0
	//ここは少し理解しづらいです。
	//*でポインタの値を呼び出しますが、それをしなくても&をいきなり書くことでその変数へのポインタをいきなり呼び出すことができます。
	p  = &Vertex{1, 2} // has type *Vertex
)

func main() {
	fmt.Println(v1, p, v2, v3)
}

