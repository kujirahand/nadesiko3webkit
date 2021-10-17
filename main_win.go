//go:build windows
// +build windows

package main

import (
	"strconv"
	"time"

	webview "github.com/jchv/go-webview2"
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
