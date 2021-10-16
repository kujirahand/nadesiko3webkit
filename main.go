package main

import (
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/webview/webview"
)

const LOCAL_SERVER_ADDR = "127.0.0.1:17145"

func startServer() {
	http.HandleFunc("/", IndexErrorHandler)
	http.HandleFunc("/webapp/", func(w http.ResponseWriter, r *http.Request) {
		log.Println("[REQ]" + r.RequestURI)
		http.ServeFile(w, r, r.URL.Path[1:])
	})
	err := http.ListenAndServe(LOCAL_SERVER_ADDR, nil)
	if err != nil {
		log.Fatal(err)
	}
}

func IndexErrorHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(404)
	w.Header().Set("Content-Type", "text/html; charset=utf8")
	w.Write([]byte("<html><body><h1>Loading Error</h1></body></html>"))
}

func main() {
	debug := true
	utime := strconv.FormatInt(time.Now().Unix(), 16)

	go startServer()
	w := webview.New(debug)
	defer w.Destroy()
	w.SetTitle("なでしこ3")
	w.SetSize(800, 600, webview.HintNone)
	w.Navigate("http://" + LOCAL_SERVER_ADDR + "/webapp/index.html?time=" + utime)
	w.Run()
}
