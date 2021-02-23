//スライスは長さと容量を両方持っています。
//スライスの容量は、スライスのはじめの要素から数えて、元の配列があとどれだけあるかという数字です。わかりにくいので下でもう一度説明します。
//長さはそのスライスに含まれる要素数のことで
package main

import "fmt"

func main() {
	s := []int{2, 3, 5, 7, 11, 13}
	//ここで元の配列が長さ6で宣言されたと同時にsの長さも容量も６になりました
	printSlice(s)

	// Slice the slice to give it zero length.
	//ここでは長さは０の配列ですが、始まりが１番目からなので容量は６となりました。
	s = s[:0]
	printSlice(s)

	// Extend its length.
	//ここも始まりは一個目なので容量は６ですが、長さは４ですね
	s = s[:4]
	printSlice(s)

	// Drop its first two values.
	//始まりは３個めなので６番目までに４つの要素が入れるので容量は４です。長さも４ですね。
	s = s[2:]
	printSlice(s)
}
//ここでまとめとして、スライスの中で配列の中での最後の値を指定しなければスライスの容量と長さは同じものになるということがわかりました。
func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}

