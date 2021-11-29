package persona

import (
	"bufio"
	"fmt"
	"net/http"
	"os"
	"path/filepath"

	svg "github.com/ajstarks/svgo"
)

type CardInfo struct {
	Name string
	Job  string
	Cost int
}

func DeckMake(s string) CardInfo {
	var ci CardInfo

	ci.Name = s
	ci.Job = "ALL"
	createSVG(ci)
	return ci
}

// Handler is Vercelにデプロイした時に「/api」でここが呼ばれる
func Handler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Handler-start.")
	myURL := r.URL.Path

	q := r.URL.Query()
	v := q.Get("actor")

	fv := r.FormValue("actorItem")

	// item情報をセットする
	item := fv
	if len(v) > 0 {
		item = v
	}

	noDataPage := ""
	// 検索文字がなければエラーとする
	if len(item) == 0 {
		noDataPage = "<h1> 検索対象データがありません. </h1>" + myURL
		fmt.Fprintf(w, noDataPage)
		return
	}
	mainD := DeckMake(item)

	startPage := "<h1>Hey from Go!</h1>" + mainD.Name + v + fv + "<br>" + myURL
	imgTag := fmt.Sprintf(`<img src="/data/%s.svg" />`, v)
	startPage = startPage + imgTag
	// HTMLを描画
	fmt.Fprintf(w, startPage)

	if myURL != "/api" {
		errorHandler(w, r, http.StatusNotFound)
		return
	}
}

// SvgHandler is Vercelにデプロイした時に「/api」でここが呼ばれる
func SvgHandler(w http.ResponseWriter, r *http.Request) {

	fmt.Println("SvgHandler-start.")

	myURL := r.URL.Path

	q := r.URL.Query()
	v := q.Get("actor")

	fv := r.FormValue("actorItem")

	mainD := DeckMake(v)

	startPage := "<h1>Hey!! from Go!</h1>" + mainD.Name + v + fv

	// HTMLを描画
	fmt.Fprintf(w, startPage)

	if myURL != "/api" {
		errorHandler(w, r, http.StatusNotFound)
		return
	}
}

func errorHandler(w http.ResponseWriter, r *http.Request, status int) {
	w.WriteHeader(status)
	if status == http.StatusNotFound {
		fmt.Fprint(w, "custom 404")
	}
}

func createSVG(rec CardInfo) {
	fmt.Printf("createSVG-start.%v \n", rec.Name)
	svgStyle := map[string]string{}

	// useAndroid()
	svgfilename := filepath.Join("storage", "%s.svg")
	// svgfilename := "%s.svg"
	f, err := os.Create(fmt.Sprintf(svgfilename, rec.Name))
	if err != nil {
		panic(err)
	}
	defer f.Close()

	// キャンバスサイズ
	width := 2400
	height := 100

	w := bufio.NewWriter(f)
	defer w.Flush()

	// キャンバスを作る
	canvas := svg.New(w)
	canvas.Start(width, height)
	canvas.Title(rec.Name)
	baseURL := "https://go-persona.vercel.app/"
	canvas.Link(baseURL, "HOMEに戻る")
	canvas.LinkEnd()
	idx := 0
	basePos := 0

	idx = 0

	fontSize := 20

	// 配置
	svgStyle["text-anchor"] = "start"
	svgStyle["font-size"] = fmt.Sprintf("%dpx", fontSize)
	// 文字の配置（lr=横書き）
	svgStyle["writing-mode"] = "lr"
	// 色
	svgStyle["fill"] = "black"
	const itemStyleBase = "text-anchor:start;font-size:%dpx; writing-mode:%s;fill:black"
	itemStyle := fmt.Sprintf(itemStyleBase, fontSize, "lr")
	itemText := fmt.Sprintf("%s【%s】(%d)", rec.Name, rec.Job, rec.Cost)
	// 横位置
	posX := 10
	// 縦位置
	posY := basePos + fontSize // (idx * fontSize)
	canvas.Link(baseURL, "★")
	canvas.Text(posX, posY, itemText, itemStyle)
	fmt.Printf("Twitter//Y:%d,base:%d[%d] font:%d \n", posY, basePos, idx, fontSize)
	canvas.LinkEnd()
	idx++
	basePos = posY

	canvas.End()

}
