package main

import "fmt"

const (
	// Create a huge number by shifting a 1 bit left 100 places.
	// In other words, the binary number that is 1 followed by 100 zeroes.
	//<<みたいな記号は２進数に直してから桁をずらすときに使います。
	//この下の例だと、1を２進数にして右に１００桁ずらしているので100~000的な大きめの２進数になります。
	Big = 1 << 100
	//こんなんちゃんと説明しないとわからないよね、、、
	// Shift it right again 99 places, so we end up with 1<<1, or 2.
	//次に左に99bitずらしているので最終的には右に1桁ずらしたのと同じものです。
	Small = Big >> 99
)

func needInt(x int) int { return x*10 + 1 }
func needFloat(x float64) float64 {
	return x * 0.1
}

func main() {
	fmt.Println(needInt(Small))
	fmt.Println(needFloat(Small))
	fmt.Println(needFloat(Big))
}

