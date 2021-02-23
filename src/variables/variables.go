//変数は宣言してから使います。複数同時に宣言することもできますし、その複数の変数が同じ型なら最後に一回だけ型を宣言すれば良いです。
package main

import "fmt"

//関数func main以外でも使うものは外側で宣言する
var c, python, java bool

func main() {
	//関数内でしか使わないiはfuncの中で宣言している。
	var i int
	i = 2
	c = true
	fmt.Println(i, c, python, java)
}
