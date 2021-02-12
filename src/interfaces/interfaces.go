//複数のメソッドを一つの変数として扱える。
package main

import (
	"fmt"
	"math"
)
//ここがインターフェース型の変数です。
type Abser interface {
	Abs() float64
}

func main() {
	//ここでインターフェース型の変数をaに入れています。
	var a Abser
	//fはフロート型、vはストラクト型
	f := MyFloat(-math.Sqrt2)
	v := Vertex{3, 4}

	a = f  // a MyFloat implements Abser
	a = &v // a *Vertex implements Abser

	// In the following line, v is a Vertex (not *Vertex)
	// and does NOT implement Abser.
	//a = v

	fmt.Println(a.Abs())
}

type MyFloat float64

func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

type Vertex struct {
	X, Y float64
}

func (v *Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}
