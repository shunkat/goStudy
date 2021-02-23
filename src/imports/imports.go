package main

import (
	"fmt"
	"math"
)
//mathというパッケージをインポートしているのでmathのSqrtを使える。（平方根を求める)
//このときに外部から参照できる値は大文字から始まっている。
//例えばPrintfとかSqrtとか大文字ですよね
func main() {
	fmt.Printf("Now you have %g problems.\n", math.Sqrt(7))
}

