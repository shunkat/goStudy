package main

import (
	"fmt"
	"math"
)

func sqrt(x float64) string {
	if x < 0 {
		//xがマイナスだった場合にプラスにしてから平方根にして虚数記号のiをつけています
		return sqrt(-x) + "i"
	}

	return fmt.Sprint(math.Sqrt(x))
}

func main() {
	fmt.Println(sqrt(2), sqrt(-4))
}
