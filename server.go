package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"runtime"
	"strconv"
	"strings"
)

type ApiResult struct {
	Result bool   `json:"result"`
	Tag    string `json:"tag"`
	Value  string `json:"value"`
}

func Exists(name string) bool {
	_, err := os.Stat(name)
	return !os.IsNotExist(err)
}

func GetBokanPath() string {
	cur, _ := os.Getwd()
	webapp := filepath.Join(cur, DIR_WEBAPP)
	if Exists(webapp) {
		return cur
	}
	exe, _ := os.Executable()
	return filepath.Dir(exe)
}

// アプリ専用の保存フォルダを得る
func getUserDir() string {
	// get home path
	home := os.Getenv("HOMEPATH")
	if runtime.GOOS != "windows" {
		home = os.Getenv("HOME")
	}
	appid := url.PathEscape(GlobalInfo.AppId)
	return filepath.Join(home, ".nadesiko3", appid)
}

func getUserFilename(name string) string {
	name = url.PathEscape(name)
	path := filepath.Join(getUserDir(), name)
	return path
}

func SaveData(name string, value string) error {
	// 保存フォルダの確認
	dir := getUserDir()
	if !Exists(dir) {
		os.MkdirAll(dir, os.ModePerm)
	}
	// 保存
	path := getUserFilename(name)
	err := ioutil.WriteFile(path, []byte(value), os.ModePerm)
	if err != nil {
		return err
	}
	return nil
}

func LoadData(name string) (string, error) {
	path := getUserFilename(name)
	body, err := ioutil.ReadFile(path)
	if err != nil {
		return "", err
	}
	return string(body), nil
}

func GetFiles() ([]string, error) {
	path := getUserDir()
	files, err := ioutil.ReadDir(path)
	if err != nil {
		return nil, err
	}
	var paths []string = []string{}
	for _, f := range files {
		name, _ := url.PathUnescape(f.Name())
		paths = append(paths, name)
	}
	return paths, nil
}

// API
func apiHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	funcName := r.FormValue("func")
	switch funcName {
	case "save":
		name := r.FormValue("name")
		value := r.FormValue("value")
		fmt.Printf("name=%s, value=%s", name, value)
		err := SaveData(name, value)
		if err != nil {
			setResponse(w, &ApiResult{Result: false, Value: err.Error(), Tag: name})
		} else {
			setResponse(w, &ApiResult{Result: true, Value: "success", Tag: name})
		}
	case "load":
		name := r.FormValue("name")
		value, err := LoadData(name)
		if err != nil {
			setResponse(w, &ApiResult{Result: false, Value: err.Error(), Tag: name})
		} else {
			setResponse(w, &ApiResult{Result: true, Value: value, Tag: name})
		}
	case "exists":
		name := r.FormValue("name")
		if Exists(getUserFilename(name)) {
			setResponse(w, &ApiResult{Result: true, Value: "1", Tag: name})
		} else {
			setResponse(w, &ApiResult{Result: true, Value: "0", Tag: name})
		}
	case "files":
		files, err := GetFiles()
		if err != nil {
			setResponse(w, &ApiResult{Result: false, Value: err.Error(), Tag: "error"})
		} else {
			setResponse(w, &ApiResult{Result: true, Value: strings.Join(files, ","), Tag: "files"})
		}
	default:
		setResponse(w, &ApiResult{Result: false, Value: "no func", Tag: "error"})
	}
}

func StartServer(info *IndexInfo) {
	rootDir := GetBokanPath()
	// URI
	http.HandleFunc("/", indexErrorHandler)
	http.HandleFunc("/webapp/", func(w http.ResponseWriter, r *http.Request) {
		file := filepath.Join(rootDir, r.URL.Path[1:])
		log.Println("[REQ]" + r.RequestURI)
		log.Println("[FILE]" + file)
		http.ServeFile(w, r, file)
	})
	http.HandleFunc("/api", apiHandler)
	// start server
	addr := "127.0.0.1:" + strconv.Itoa(info.Port)
	fmt.Printf("[Server] http://%s\n", addr)
	err := http.ListenAndServe(addr, nil)
	if err != nil {
		log.Fatal(err)
	}
}

func setResponse(w http.ResponseWriter, result *ApiResult) {
	w.WriteHeader(200)
	w.Header().Set("Content-Type", "application/json; charset=utf8")
	jsonData, err := json.Marshal(result)
	if err != nil {
		w.Write([]byte("{\"result\": false, \"unknown error\"}"))
		return
	}
	w.Write([]byte(jsonData))
}

func indexErrorHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(404)
	w.Header().Set("Content-Type", "text/html; charset=utf8")
	w.Write([]byte("<html><body><h1>Loading Error</h1></body></html>"))
}

func checkPort(info *IndexInfo) {
	if info.Port <= 0 {
		// 適当に空いているポートを探す
		l, err2 := net.Listen("tcp", "127.0.0.1:0")
		if err2 != nil {
			// 空きポートの検索に失敗
			log.Fatal(err2)
		}
		// ポート番号を得る
		addr := l.Addr().String()
		a := strings.Split(addr, ":")
		pno, err3 := strconv.Atoi(a[1])
		if err3 != nil {
			log.Fatal(err3)
		}
		info.Port = pno
		l.Close() // HTTPサーバーの起動前にソケットを閉じておく
		return
	}
	// ポートが利用可能か調べる
	ll, err := net.Listen("tcp", "127.0.0.1:"+strconv.Itoa(info.Port))
	if err != nil {
		info.Port = 0
		checkPort(info)
		return
	} else {
		fmt.Printf("Server Addr: %s\n", ll.Addr().String())
		ll.Close()
	}

}
