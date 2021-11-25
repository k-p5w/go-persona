package persona

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"

	svg "github.com/ajstarks/svgo"
)

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
