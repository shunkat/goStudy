//mainのfuncが最終的に実行される、色々funcがあるが結局最後にはコレが実行されるファイルだよという意味になっている。
//いろいろな関数をmainが呼び出していればあっている
//package mainがディレクトリの中に一つだけで、それが一つのfunc mainを呼び出すというイメージ
package main

import (
	"fmt"
	"math/rand"
)

func main() {
	fmt.Println("My favorite number is", rand.Intn(10))
}
