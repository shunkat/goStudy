package main

import (
	"fmt"
	"math"
)

type Vertex struct {
	X, Y float64
}

func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}
//ここでポインター越しにアクセスして値を変えています。
//ポインターのアスタリスクを消すと値が５に変わります。
//ポインターの場合はそのメモリの場所の値を更新するような形で変わります。
//しかしポインター出ない場合は新しくvertex{3,4}が生まれてもとの値が変更されないです。
func (v *Vertex) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func main() {
	v := Vertex{3, 4}
	v.Scale(10)
	fmt.Println(v.Abs())
}
