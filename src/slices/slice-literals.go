//リテラルというのは数値や文字列を直接記述した定数のこと。まぁ要はベタ書きした値のことですね
//スライスのリテラルは長さのない配列のようなものです。
//参考までに
//配列の場合は[3]bool{true, true, false}
//スライスの場合は[]bool{true, true, false}
//このようにスライスのときは配列のように長さを指定しなくても書くことができます。
package main

import "fmt"

func main() {
	q := []int{2, 3, 5, 7, 11, 13}
	fmt.Println(q)

	r := []bool{true, false, true, true, false, true}
	fmt.Println(r)
//ここでは構造体を定義して、その中にint型のiとbool型のbが入るよという宣言しています。
	//その後に、中に複数のiとbの値をそのまま代入しています。
	s := []struct {
		i int
		b bool
	}{
		{2, true},
		{3, false},
		{5, true},
		{7, true},
		{11, false},
		{13, true},
	}
	fmt.Println(s)
}

