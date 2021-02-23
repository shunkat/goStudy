package main

import "fmt"

func main() {
	names := [4]string{
		"John",
		"Paul",
		"George",
		"Ringo",
	}
	fmt.Println(names)

	a := names[0:2]
	b := names[1:3]
	fmt.Println(a, b)
//スライスも中の値を再代入すると参照元の配列の値も切り替わります。
	b[0] = "XXX"
	fmt.Println(a, b)
	fmt.Println(names)
}

