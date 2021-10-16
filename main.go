package main

import (
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"time"

	webview "github.com/jchv/go-webview2"
	//"github.com/webview/webview"
)

const LOCAL_SERVER_ADDR = "127.0.0.1:17145"

func Exists(name string) bool {
	_, err := os.Stat(name)
	return !os.IsNotExist(err)
}

func getBokanPath() string {
	cur, _ := os.Getwd()
	webapp := filepath.Join(cur, "webapp")
	if Exists(webapp) {
		return cur
	}
	exe, _ := os.Executable()
	return filepath.Dir(exe)
}

func startServer() {
	rootDir := getBokanPath()
	http.HandleFunc("/", IndexErrorHandler)
	http.HandleFunc("/webapp/", func(w http.ResponseWriter, r *http.Request) {
		file := filepath.Join(rootDir, r.URL.Path[1:])
		log.Println("[REQ]" + r.RequestURI)
		log.Println("[FILE]" + file)
		http.ServeFile(w, r, file)
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
	w.SetTitle("nadesiko3")
	w.SetSize(800, 600, webview.HintFixed)
	w.Navigate("http://" + LOCAL_SERVER_ADDR + "/webapp/index.html?time=" + utime)
	w.Run()
}
