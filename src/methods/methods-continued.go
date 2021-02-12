package main

import (
	"fmt"
	"math"
)
//ここではまずMyFloatをfloat型で定義して
type MyFloat float64
//ここでそのMyFloatにメソッドを定義しています。
func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

func main() {
	f := MyFloat(-math.Sqrt2)
	fmt.Println(f.Abs())
}
