package main

import "fmt"
//スライスは組み込みのmake関数を使用して、作成することができます。
//make関数はゼロ化された配列を割り当ててその配列を指すスライスを変えします。
//具体的に見ていきましょう
func main() {
	//１つ目の引数はどの型のスライスを作成するか、２つ目の引数はそのスライスの長さです
	//この例だと値が全部ゼロだけど５この要素を持つスライスを割り当てています
	a := make([]int, 5)
	printSlice("a", a)
	//３つ目の引数は容量です。このときには容量を５で長さが０のスライスを作ります。
	b := make([]int, 0, 5)
	printSlice("b", b)
	//このときにはcはbのなかの２つ目までを取り出しているので０が２つはいったスライスを宣言しているのと同じです。
	c := b[:2]
	printSlice("c", c)

	d := c[2:5]
	printSlice("d", d)
	//このように動的な長さを持つスライスを作ることもmake関数なら可能です。
}

func printSlice(s string, x []int) {
	fmt.Printf("%s len=%d cap=%d %v\n",
		s, len(x), cap(x), x)
}

