package main

import "fmt"
//先程のようにVertexという構造体が定義されてその中にXとYというint型の変数が格納されています。
type Vertex struct {
	X int
	Y int
}

func main() {
//上での宣言の方に矛盾が無いような変数をここで代入しています。
	//具体的にはx=1,y=2という形です。ここでも宣言で使われてのと同じような{}が使われていることが注意点です。
	v := Vertex{1, 2}
//v.Xのように.を用いて中身に値を入れたリ、その値を呼び出したりすることができます。
	v.X = 4
	fmt.Print(v.X)
}

