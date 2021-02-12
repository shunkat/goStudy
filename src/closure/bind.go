package main

import "fmt"

func adder() func(int) int {
	sum := 0
	//以下は戻り値として使う関数なのだが、sumはこのスコープの外にあるsum変数を使ってしかもそれを返している
	return func(x int) int {
		sum += x
		return sum
	}
}

func main() {
	pos, neg := adder(), adder()
	for i := 0; i < 10; i++ {
		fmt.Println(
			i,
			//posは今までのiを足し合わせる関数
			pos(i),
			//negはposの返り値のsumに対して−２をかけている
			neg(-2*i),
		)
	}
}
