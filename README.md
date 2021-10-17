# nadesiko3webkit

nadesiko3 for webkit

# コンパイルに必要なライブラリ

 - Go言語

## コンパイルの方法

必要なモジュールを取得する。

```bash
go get -u
```

## Windowsの場合

WindowsではChromium版のEdgeに加えて、WebView2のランタイム「[こちら](https://developer.microsoft.com/en-us/microsoft-edge/webview2/)」が必要です。

```bash
batch¥build-win.bat
```

## macOSの場合

```bash
batch/build-mac.sh
```

## (memo) macOSとWindowsで使うライブラリの違い

 - macOSでは[webview](https://github.com/webview/webview)を使う
 - Windowsでは[go-webview2](https://github.com/jchv/go-webview2)を使う

