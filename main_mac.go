//go:build darwin || linux
// +build darwin linux

package main

import (
	"github.com/zserge/lorca"
	// "github.com/webview/webview"
)

func main() {
	info := ReadIndexJson()

	// ローカルサーバーを起動
	go StartServer(&info)

	// ブラウザを起動
	ui, _ := lorca.New(GetIndexURI(&info), "", info.Width, info.Height)
	defer ui.Close()

	<-ui.Done()
}

/*
// webview main
func main() {
	debug := true
	info := ReadIndexJson()

	// ローカルサーバーを起動
	go StartServer(&info)

	// ブラウザを起動
	w := webview.New(debug)
	defer w.Destroy()
	w.SetTitle(info.Title)
	hint := webview.HintNone
	if !info.Resize {
		hint = webview.HintFixed
	}

	w.Bind("alert", func(text string) bool {
		return messagebox(text, `"Yes"`)
	})
	w.Bind("confirm", func(text string) bool {
		return messagebox(text, `"No", "Yes"`)
	})

	w.SetSize(info.Width, info.Height, webview.Hint(hint))
	w.Navigate(GetIndexURI(&info))
	w.Run()
}

func messagebox(text string, buttons string) bool {
	title := GlobalInfo.Title
	script := `set T to button returned of ` +
		`(display dialog "%s" with title "%s" buttons {%s} default button "Yes")`
	out, err := exec.Command("osascript", "-e", fmt.Sprintf(script, text, title, buttons)).Output()
	if err != nil {
		if exitError, ok := err.(*exec.ExitError); ok {
			return exitError.Sys().(syscall.WaitStatus).ExitStatus() == 0
		}
	}
	return strings.TrimSpace(string(out)) == "Yes"
}

*/
