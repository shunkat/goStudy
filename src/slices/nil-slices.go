package main

import "fmt"
//値を何も代入しないスライスはnilとなり、長さと容量がともに０です
func main() {
	var s []int
	fmt.Println(s, len(s), cap(s))
	if s == nil {
		fmt.Println("nil!")
	}
//0個長さの配列を宣言するのと同じだが、それだとnilにはならない。
	var a [0] int
	fmt.Println(a, len(a), cap(a))
}

