package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"

	"github.com/google/gxui"
	"github.com/google/gxui/drivers/gl"
	"github.com/google/gxui/math"
	"github.com/google/gxui/samples/flags"
)

func slurp(url string) (string, error) {
	res, err := http.Get(url)
	if err != nil {
		return "", err
	}
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return "", err
	}
	res.Body.Close()
	return string(body), nil
}

func list_naver() []string {
	s, err := slurp("http://www.naver.com")
	if err != nil {
		return nil
	}
	return re_groups(regexp.MustCompile("<option value=\".+\">.+: (.+)</option>"),
		// <option value=".+">.+: (.+)</option>
		s,
		1)
}

// return all specific groups
func re_groups(re *regexp.Regexp, text string, group int) []string {
	result := []string{}
	found := re.FindAllStringSubmatch(text, -1)
	for _, v := range found {
		result = append(result, v[group])
	}
	return result
}

func naverPicker(theme gxui.Theme, overlay gxui.BubbleOverlay) gxui.Control {
	// items 에 list_naver []string 반환
	items := list_naver()

	adapter := gxui.CreateDefaultAdapter()
	adapter.SetItems(items)

	layout := theme.CreateLinearLayout()
	layout.SetDirection(gxui.TopToBottom)

	label0 := theme.CreateLabel()
	label0.SetText(fmt.Sprintln("네이버:"))
	layout.AddChild(label0) // layout 에 label0 붙이기

	dropList := theme.CreateDropDownList()
	dropList.SetAdapter(adapter)
	dropList.SetBubbleOverlay(overlay)
	layout.AddChild(dropList)

	list := theme.CreateList()
	list.SetAdapter(adapter)
	list.SetOrientation(gxui.Vertical)
	layout.AddChild(list)

	return layout
}

type customAdapter struct {
	gxui.AdapterBase
}

func (a *customAdapter) Count() int {
	return 1000
}

func (a *customAdapter) ItemAt(index int) gxui.AdapterItem {
	return index
}

func (a *customAdapter) ItemIndex(item gxui.AdapterItem) int {
	return item.(int)
}

func (a *customAdapter) Size(theme gxui.Theme) math.Size {
	return math.Size{W: 100, H: 100}
}

func (a *customAdapter) Create(theme gxui.Theme, index int) gxui.Control {
	phase := float32(index) / 1000
	c := gxui.Color{
		R: 0.5 + 0.5*math.Sinf(math.TwoPi*(phase+0.000)),
		G: 0.5 + 0.5*math.Sinf(math.TwoPi*(phase+0.333)),
		B: 0.5 + 0.5*math.Sinf(math.TwoPi*(phase+0.666)),
		A: 1.0,
	}
	i := theme.CreateImage()
	i.SetBackgroundBrush(gxui.CreateBrush(c))
	i.SetMargin(math.Spacing{L: 3, T: 3, R: 3, B: 3})
	i.OnMouseEnter(func(ev gxui.MouseEvent) {
		i.SetBorderPen(gxui.CreatePen(2, gxui.Gray80))
	})
	i.OnMouseExit(func(ev gxui.MouseEvent) {
		i.SetBorderPen(gxui.TransparentPen)
	})
	i.OnMouseDown(func(ev gxui.MouseEvent) {
		i.SetBackgroundBrush(gxui.CreateBrush(c.MulRGB(0.7)))
	})
	i.OnMouseUp(func(ev gxui.MouseEvent) {
		i.SetBackgroundBrush(gxui.CreateBrush(c))
	})
	return i
}

func appMain(driver gxui.Driver) {
	theme := flags.CreateTheme(driver)

	// fontdata, err := ioutil.ReadFile("malgun.ttf")
	// if err != nil {
	// 	log.Fatalln(err)
	// }
	// font, err := driver.CreateFont(fontdata, 12)
	// if err != nil {
	// 	log.Fatalln(err)
	// }
	// theme.SetDefaultFont(font)

	overlay := theme.CreateBubbleOverlay()

	holder := theme.CreatePanelHolder()
	holder.AddPanel(naverPicker(theme, overlay), "naver")

	window := theme.CreateWindow(500, 600, "인기 검색어")
	window.SetScale(flags.DefaultScaleFactor)
	window.AddChild(holder)
	window.AddChild(overlay)
	window.OnClose(driver.Terminate)
	window.SetPadding(math.Spacing{L: 10, T: 10, R: 10, B: 10})
}

func main() {
	gl.StartDriver(appMain)
}
