package main

import (
	"fmt"
	"io/ioutil" //コレはgo言語で他のファイル形式を読み取るときに使うパッケージです。
	"net/http"
)

//wikiのデータ構造,Pageという構造体型の別名にTitleというstringとBodyというバイト型の変数が格納されて、それが一つのページとなっています。
type Page struct {
	Title string //タイトル
	Body  []byte //タイトルの中身

}

//パスのアドレスを設定して文字の長さを定数として持つ
const lenPath = len("/view/")

//前回のハンドラーをviewHandlerとして関数化します
func viewHandler(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Path[lenPath:]
	p, _ := loadPage(title)
	fmt.Fprintf(w, "<h1>%s</h1><div>%s</div>", p.Title, p.Body)
}

//まずはページの保存とそれを読み取れるようなものを目指します、
//ここではPage型の変数のポインターpに対してメソッドとしてsaveを定義しています。
//その返り値はerror型です。
func (p *Page) save() error {
	//タイトルの名前でテキストファイルを作成して保存します。
	filename := p.Title + ".txt"
	//０６００はテキストデータを書き込んだり読み込んだりする権限を設定しており今回は自分のみです。
	return ioutil.WriteFile(filename, p.Body, 0600)
}

//titleからファイル名を読み込んで、新しいPage型のポインター変数を返す。
//string型のtitleという引数をとって*Page型の返り値を返します。
//やっている作業的には与えられたtitleのテキストファイル名をfilenameという変数に代入して、ioutil.Readfileでそのファイル名のファイルをよ見出しています。
//もしエラーが入っていたらエラーには読み出し時のエラーの値が、*Pageにはnilが入ります。
//逆にエラーが何もなかった場合はPage型のpointer変数に引数で与えられたtitleと読みだしたbodyが格納されたものとerr値がnilの返り値が帰ってきます。
func loadPage(title string) (*Page, error) {
	filename := title + ".txt"

	body, err := ioutil.ReadFile(filename)
	//errに値が入ったらエラーとしてbodyの値をnilにして返します。
	if err != nil {
		return nil, err
	}
	return &Page{Title: title, Body: body}, nil
}

func main() {
	//ここでp1という仮のページを格納する変数を用意して、そこにtest用のタイトルとbodyを格納します。
	p1 := &Page{Title: "TestPage", Body: []byte("This is a smple page")}
	//そのあとにそれを保存して
	p1.save()
	//p2という変数に先程のページを読み出して代入します。そのときに使う関数は上で定義しているloadpageです
	p2, _ := loadPage("TestPage")
	fmt.Println(string(p2.Body))
}