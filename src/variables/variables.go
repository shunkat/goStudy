package main

import (
	"fmt"
	"math"
)

var k int = 22

func main() {
	//int という数字を合う買う方で変数iを宣言直後に初期値の10を代入
	var i int = 10
	//:=  を使って型宣言や初期化をしないで使うこともできます。
	j := 6
	fmt.Printf("%T %v\n", i, j)
	var s float64 = math.Sqrt(float64(j))
	fmt.Println(s)
}
