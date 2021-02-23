//型変換はgoではどのように実現されているのでしょうか？
//
package main

import (
	"fmt"
	"math"
)

func main() {
	var x, y int = 3, 4
	//ここではint型のxとyを一旦float型に直してからｆに入れています。そうしないとエラーになるから
	var f float64 = math.Sqrt(float64(x*x + y*y))
	//その後でfを再びuint型に入れ直しています。まぁ型名(元の変数)という書き方なのでわかりやすいですね
	var z uint = uint(f)
	fmt.Println(x, y,f, z)
}

