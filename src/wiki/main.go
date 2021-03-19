package main

import (
	"errors" //新しく自分でエラー文を作るためのpackage
	"html/template"
	"io/ioutil" //コレはgo言語で他のファイル形式を読み取るときに使うパッケージです。
	"log"
	"net/http"
	"regexp" //これは正規表現用のpackeage
	"strings"
)

// Page はwikiのデータ構造,Pageという構造体型の別名にTitleというstringとBodyというバイト型の変数が格納されて、それが一つのページとなっています
type Page struct {
	Title string //タイトル
	Body  []byte //タイトルの中身

}

//パスのアドレスを設定して文字の長さを定数として持つ
const lenPath = len("/view/")

// テンプレートファイルの配列を用意,keyを持つ配列templatesを用意した。
var templates = make(map[string]*template.Template)

// 正規表現でURLを生成できる大文字小文字の英字と数字を判別するため
//^は一文字目のチェック、[]はその中の文字のチェック＄は最後の文字にマッチ+1文字以上
//これにより/などが使えなくなっている
var titleValidator = regexp.MustCompile("^[a-zA-Z0-9]+$")

//.txt
const expend_string = ".txt"

// 初期化関数を用意する、これはmain関数より先に呼び出される
func init() {
	for _, tmpl := range []string{"edit", "view"} {
		//templateMustではOSのエラーのときにはpanicが起こって操作が終了するというエラー処理をしてくれる
		t := template.Must(template.ParseFiles(tmpl + ".html"))
		templates[tmpl] = t
	}
}

// タイトルのチェックを行う
func getTitle(w http.ResponseWriter, r *http.Request) (title string, err error) {
	title = r.URL.Path[lenPath:]
	if !titleValidator.MatchString(title) {
		//これにより意図しないURLに対しては４０４notFoundが表示されるようになった。
		http.NotFound(w, r)
		err = errors.New("Invalid Page Title")
		//自由に作ったエラーメッセージをlog出力
		log.Print(err)
	}
	return
}

//前回のハンドラーをviewHandlerとして関数化します
func viewHandler(w http.ResponseWriter, r *http.Request, title string) {
	p, err := loadPage(title)
	if err != nil {
		//新規ページのURLが入力されたときはeditHanderのURLに飛ばすことで編集ページに飛ばすことができます。
		http.Redirect(w, r, "/edit/"+title, http.StatusFound)
		return
	}
	//ここでは読み込んだページのtitleをh1タグに、bodyをdivタグに入れています。
	renderTemplate(w, "view", p)
}

//ページの編集用のページを作成します。
func editHandler(w http.ResponseWriter, r *http.Request, title string) {
	p, err := loadPage(title)
	if err != nil {
		p = &Page{Title: title}
	}
	renderTemplate(w, "edit", p)
}

//引数は他と同様
func saveHandler(w http.ResponseWriter, r *http.Request, title string) {
	body := r.FormValue("body")
	p := &Page{Title: title, Body: []byte(body)}
	// ここにエラーハンドリングを追加
	err := p.save()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	http.Redirect(w, r, "/view/"+title, http.StatusFound)
	// 以上のコードによって、存在しないページのeditボタンが押されたときに、入力されたbodyとtitleのページが新たに作られるような関数ができた。
}
func topHandler(w http.ResponseWriter, r *http.Request) {
	//今のmain.goがいるディレクトリの階層にある.txtデータを取得する
	files, err := ioutil.ReadDir("./")
	if err != nil {
		err = errors.New("所定のディレクトリ内にテキストファイルがありません")
		log.Print(err)
		return
	}

	var paths []string    //textdataの名前
	var fileName []string //testdataのファイル名
	for _, file := range files {
		//対象となる.txtデータのみを取得
		if strings.HasSuffix(file.Name(), expend_string) {
			//このままでは.txtも含めたすべてが表示されてしまうので、テキストデータの.txtで名前をスライスしたものをfileNameにいれる
			fileName = strings.Split(string(file.Name()), expend_string)
			paths = append(paths, fileName[0])
		}
	}
	// テキストファイルがなくてpathsがnilだった場合にエラーを表示させる。
	if paths == nil {
		err = errors.New("テキストファイルが存在しません")
		log.Print(err)
	}

	t := template.Must(template.ParseFiles("top.html"))
	err = t.Execute(w, paths)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

//
func makeHandler(fn func(http.ResponseWriter, *http.Request, string)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Requestからページタイトルを取り出して、fnを呼び出す。
		title := r.URL.Path[lenPath:]
		if !titleValidator.MatchString(title) {
			//これにより意図しないURLに対しては４０４notFoundが表示されるようになった。
			http.NotFound(w, r)
			err := errors.New("Invalid Page Title")
			//自由に作ったエラーメッセージをlog出力
			log.Print(err)
			return
		}
		fn(w, r, title)
	}
}

// // 以下は自分のポートフォリオ用のページの作成メソッドです。
// func portHandler(w http.ResponseWriter, r *http.Request) {
// 	//このページは他と違い単一のページなのでタイトルも直打ちしちゃいます。
// 	title := port
// 	p, err := loadPage(title)
// 	if err != nil {
// 		//もし正しいURL出なかったときはエラーメッセージを返します。
// 		http.Error(w, err.Error(), http.StatusInternalServerError)
// 		return
// 	}
// 	//ここでは読み込んだページのtitleをh1タグに、bodyをdivタグに入れています。
// 	renderTemplate(w, "view", p)
// }

//viewとeditで共通する部分の切り出し
func renderTemplate(w http.ResponseWriter, tmpl string, p *Page) {

	//tmol.htmlのなかにTitleやBodyを入れれるようにする
	err := templates[tmpl].Execute(w, p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
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
	//makeHandlerのなかで関数のvalidationを行うように変更
	http.HandleFunc("/view/", makeHandler(viewHandler))
	http.HandleFunc("/edit/", makeHandler(editHandler))
	// saveというURLに飛ばされたときにsaveHandlerというメソッドが動くようにしている。
	http.HandleFunc("/save/", makeHandler(saveHandler))
	http.HandleFunc("/top/", topHandler)
	//p2という変数に先程のページを読み出して代入します。そのときに使う関数は上で定義しているloadpageです
	http.ListenAndServe(":8080", nil)
}
