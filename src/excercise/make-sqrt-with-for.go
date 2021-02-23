package main

import (
	"fmt"
	//平方根を使えるようにインポート
	"math"
)

func Sqrt(x float64) float64 {
	z := 1.0
	//ここで平方根に近づける式を繰り返す
	for math.Abs((z*z) -x) > 0.000000000001 {
		z -= (z*z - x) / (2*z)
	}
	return z
}

func main() {
	fmt.Println(Sqrt(2))
	fmt.Println(math.Sqrt(2))
}

