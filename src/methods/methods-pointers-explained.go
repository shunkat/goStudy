package main


import (
	"fmt"
	"math"
)

//
type Vertex struct {
	X, Y float64
}

func Abs(v Vertex) float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}
//ここのアスタリスクを消して
func Scale(v *Vertex, f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func main() {
	v := Vertex{3, 4}
  //ここの＆を消しても先ほどと同じで値が5になります。理由も同じです。
	Scale(&v, 10)
	fmt.Println(Abs(v))
}
