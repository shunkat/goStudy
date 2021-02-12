//goにはクラスが無いので変わりに型にメソッドを定義することができます。
package main

import (
	"fmt"
	"math"
)

type Vertex struct {
	X, Y float64
}
//以下のv vertexというものが特殊なレシーバというもので
func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func main() {
	v := Vertex{3, 4}
	fmt.Println(v.Abs())
}
