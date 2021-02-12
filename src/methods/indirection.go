package main

import "fmt"

type Vertex struct {
	X, Y float64
}

func (v *Vertex) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func ScaleFunc(v *Vertex, f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func main() {

	v := Vertex{3, 4}
	v.Scale(2)
	//scalefuncはポインターを引数に取ると上で宣言されているので、pointerとして渡してあげなくてはいけない。それはしたのScaleFuncも同じ。
	ScaleFunc(&v, 10)

	p := &Vertex{4, 3}
	//しかしここにあるScale関数はポインタレシーバなので変数でもポインタでもどちらでもレシーバとして受け取ることできる。
	p.Scale(3)
  //goが勝手にポインタとして変換してくれてるのです。
	ScaleFunc(p, 8)

	fmt.Println(v, p)
}
