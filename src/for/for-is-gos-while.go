//条件式以外はとっちゃってしかもその間をつなぐ;も省略することができます。
//コレがwhile文と同じだからgoにはwhile文が無いんですね。
package main

import "fmt"

func main() {
	sum := 1
	for sum < 1000 {
		sum += sum
	}
	fmt.Println(sum)
}

