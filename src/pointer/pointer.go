package main

import "fmt"
//アスタリスク*をつけることでその変数のメモリの値を指すことができます。
func main() {
	i, j := 42, 2701
//&をつけることでint型の変数iのポインタをpの中に代入しています。
//言い方を変えれば&が変数のポインタを引き出していると捉えることもできます。
	p := &i
//逆に*をつけることでそのポインタが指し示す値を引き出すこともできます。
	fmt.Println(*p) // read i through the pointer
//また、ポインタで示した場所に値を代入することでその変数を上書きすることもできます。
	*p = 21         // set i through the pointer
	fmt.Println(i)  // see the new value of i

	p = &j         // point to j
	*p = *p / 37   // divide j through the pointer
//もちろんポインタを通じて値を演算することもできるのでその形で値を上書きすることもできます。
	fmt.Println(j) // see the new value of j
}
