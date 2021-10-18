# nadesiko3webkit

これは、[日本語プログラミング言語「なでしこ3(Web版)」](https://nadesi.com/)で作ったプログラムを配布したい場合に便利なパッケージです。
なでしこで作ったゲームやツールを一般配布するのにご利用ください。

ソースからビルドする場合は、Go言語が必要になりますが、なでしこ3(Web版)のプログラムを配布したいだけであれば、[こちら](https://github.com/kujirahand/nadesiko3webkit/releases)から配布用パッケージをダウンロードするだけでOKです。

なでしこ3で作ったプログラムを配布するには、Electronを使う方法もありますが、本プロジェクト(nadesiko3webkit)を使うと、OSにインストールされているWebKitベースのブラウザのコンポーネントを使ってなでしこ3を実行するので、配布サイズが小さく手軽にプログラムを配布できるというメリットがあります。簡単なプログラムであれば、ZIP圧縮して5MB程度の配布サイズになります。


# ライブラリのコンパイル

本ライブラリを構築するには、Go言語が必要です。

 - [Go言語](https://golang.org/)

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

 - macOSでは[webview](https://github.com/webview/webview)に代わって、[lorca](https://github.com/zserge/lorca)を使う
 - Windowsでは[go-webview2](https://github.com/jchv/go-webview2)を使う

