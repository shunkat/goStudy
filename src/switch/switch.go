package main

import (
	"fmt"
	"runtime"
)

//switchに指定した値がcaseのいずれかに合致したら自動でbreakしてくれます。
//この下の式はosという変数に自分のosの情報を格納してそれに応じて条件分岐するのですが当てはまらなかった場合はdefaultが実行されます。
func main() {
	fmt.Print("Go runs on ")
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("OS X.")
	case "linux":
		fmt.Println("Linux.")
	default:
		// freebsd, openbsd,
		// plan9, windows...
		//ここはどうやら最後に.をつけるフォーマットで当てはまらなかった他のOS名を出力しているっぽい
		fmt.Printf("%s.\n", os)
	}
}
