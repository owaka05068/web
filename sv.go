package trial

import (
	"html/template"
	"net/http"
)

type Page struct {
	Title string
	Count int
}

func viewHandler(w http.ResponseWriter, r *http.Request) {
	//テンプレートをパースする
	t := template.Must(template.ParseFiles("test.html"))

	str := "メッセージあああああ"

	//テンプレート描画する
	err := t.ExecuteTemplate(w, "test.html", str)
}
func main() {
	http.HandleFunc("/", viewHandler)
	http.ListenAndServe(":8080", nil)
}
