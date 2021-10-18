package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"
)

type IndexInfo struct {
	Title  string `json:"title"`
	Width  int    `json:"width"`
	Height int    `json:"height"`
	Resize bool   `json:"resize"`
	Port   int    `json:"port"`
}

func Exists(name string) bool {
	_, err := os.Stat(name)
	return !os.IsNotExist(err)
}

func GetBokanPath() string {
	cur, _ := os.Getwd()
	webapp := filepath.Join(cur, "webapp")
	if Exists(webapp) {
		return cur
	}
	exe, _ := os.Executable()
	return filepath.Dir(exe)
}

func StartServer(info *IndexInfo) {
	rootDir := GetBokanPath()
	http.HandleFunc("/", indexErrorHandler)
	http.HandleFunc("/webapp/", func(w http.ResponseWriter, r *http.Request) {
		file := filepath.Join(rootDir, r.URL.Path[1:])
		log.Println("[REQ]" + r.RequestURI)
		log.Println("[FILE]" + file)
		http.ServeFile(w, r, file)
	})
	addr := "127.0.0.1:" + strconv.Itoa(info.Port)
	fmt.Printf("[Server] http://%s\n", addr)
	err := http.ListenAndServe(addr, nil)
	if err != nil {
		log.Fatal(err)
	}
}

func indexErrorHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(404)
	w.Header().Set("Content-Type", "text/html; charset=utf8")
	w.Write([]byte("<html><body><h1>Loading Error</h1></body></html>"))
}

func ReadIndexJson() IndexInfo {
	// index.json
	var info IndexInfo = IndexInfo{Title: "なでしこ3", Width: 640, Height: 400, Port: 8888}
	indexJson := filepath.Join(GetBokanPath(), "webapp", "index.json")
	raw, err := ioutil.ReadFile(indexJson)
	if err == nil { // 読み込めた時
		json.Unmarshal(raw, &info)
		fmt.Printf("size=%d,%d\n", info.Width, info.Height)
	}

	// ポートが利用可能か調べる
	ll, err := net.Listen("tcp", "127.0.0.1:"+strconv.Itoa(info.Port))
	if err != nil {
		// ポートが使用中だったので、適当に空いているポートを探す
		l, err2 := net.Listen("tcp", "127.0.0.1:0")
		if err2 != nil {
			log.Fatal(err2)
		}
		addr := l.Addr().String()
		a := strings.Split(addr, ":")
		pno, err3 := strconv.Atoi(a[1])
		if err3 != nil {
			log.Fatal(err3)
		}
		info.Port = pno
		l.Close() // HTTPサーバーの起動前にソケットを閉じておく
	} else {
		fmt.Printf("Server Addr: %s\n", ll.Addr().String())
	}

	return info
}

func GetIndexURI(info *IndexInfo) string {
	utime := strconv.FormatInt(time.Now().Unix(), 16)
	return "http://127.0.0.1:" + strconv.Itoa(info.Port) + "/webapp/index.html?time=" + utime
}
