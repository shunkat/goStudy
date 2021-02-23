package main

import "fmt"
//make関数のようにappendという関数をつかうことでsliceに要素を追加することができます。
func main() {
	var s []int
	printSlice(s)

	// append works on nil slices.
	//自分自身を更新する形で配列に要素を追加することができます。
	s = append(s, 0)
	printSlice(s)
	//このときにもともとのスライスの容量が足りなくなっていたら、それを増やした配列を割り当て直してその中に値を格納します。（倍々と増えていく）
	// The slice grows as needed.
	s = append(s, 1)
	printSlice(s)
	//ここでは２つ以上の値をまとめてappendしています。こういうやり方もあるんですね。
	// We can add more than one element at a time.
	s = append(s, 2, 3)
	printSlice(s)
}

func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}

