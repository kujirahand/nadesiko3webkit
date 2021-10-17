//go:build darwin || linux
// +build darwin linux

package main

import (
	"strconv"
	"time"

	"github.com/webview/webview"
)

func main() {
	debug := true
	utime := strconv.FormatInt(time.Now().Unix(), 16)

	go StartServer()
	w := webview.New(debug)
	defer w.Destroy()
	w.SetTitle("nadesiko3")
	w.SetSize(800, 600, webview.HintNone)
	w.Navigate("http://" + LOCAL_SERVER_ADDR + "/webapp/index.html?time=" + utime)
	w.Run()
}
