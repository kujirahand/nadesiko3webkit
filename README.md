# なでしこ3配布キット(nadesiko3webkit)

これは、[日本語プログラミング言語「なでしこ3(Web版)」](https://nadesi.com/)で作ったプログラムを配布したい場合に便利なパッケージです。
なでしこで作ったゲームやツールを一般配布するのにご利用ください。

## なでしこ3配布キットのダウンロード

Windows/macOSでは、配布パッケージが用意されています。[こちら](https://github.com/kujirahand/nadesiko3webkit/releases)から対象OSのパッケージをダウンロードするだけでOKです。

## 配布キットの資料

- [配布キットの解説ページ](https://nadesi.com/doc3/go.php?16421)
- 配布キットの使い方や配布キットのみで使える命令についての使い方は、[こちら](batch/res/README.md)をご覧ください。

## Linux用の配布パッケージの作り方

Linuxの場合、簡単な手順でアプリのビルドが可能です。Ubuntu/Debianではコマンドラインで以下のようにしてパッケージを作成できます。

```
# Go言語のインストール
$ sudo apt install golang git webkit2gtk-4.0 chromium
# リポジトリを取得
$ git clone https://github.com/kujirahand/nadesiko3webkit.git
# ビルド
$ cd nadesiko3webkit
$ go get -u
$ bash batch/build-linux-chrome.sh
$ bash batch/build-linux-webview.sh
```

ソースからビルドする場合は、Go言語が必要になりますが、なでしこ3(Web版)のプログラムを配布したいだけであれば、

なでしこ3で作ったプログラムを配布するには、Electronを使う方法もありますが、本プロジェクト(nadesiko3webkit)を使うと、OSにインストールしたChromeのコンポーネントを使ってなでしこ3を実行するので、配布サイズが小さく手軽にプログラムを配布できるというメリットがあります。簡単なプログラムであれば、ZIP圧縮して5MB程度の配布サイズになります。

# (参考) なでしこ配布キット(Electron版)

- [Electronを利用したなでしこ3配布キット](https://github.com/kujirahand/nadesiko3electron)もあります。配布サイズは大きくなりますが、より安定した環境で動作させることが可能です。

## 本ライブラリのコンパイル(詳細)

本ライブラリを構築するには、Go言語が必要です。

 - [Go言語](https://golang.org/)
 - 利用パッケージ (以下を選択してビルド)
   - [lorca](https://github.com/zserge/lorca)
   - [webview](https://github.com/webview/webview)
   - [webview2](https://github.com/jchv/go-webview2)

## コンパイルの方法

必要なモジュールを取得します。

```bash
go get -u
```

## コンポーネントの切り替え

以下のいずれかのバッチを実行します。

```bash
# lorcaを使う場合(mac/win)
./make_chrome.sh
.\make_chrome.bat

# WebViewを使う場合(mac/win)
./make_webview_mac.sh
./make_webview_win.bat
```

## 配布ファイルの作成

batchフォルダ以下のバッチを実行します。

