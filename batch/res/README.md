# 配布用 なでしこ3 (webkit) バイナリ

このアーカイブは、日本語プログラミング言語「なでしこ3(Web版)」で作ったプログラムを実行ファイルとしてユーザーに配布したい場合に便利なツールです。

Windows版とmacOS版のバイナリファイルが提供されています。

 - [配布URL] https://github.com/kujirahand/nadesiko3webkit/releases

# ダウンロードと動作確認

上記の配布URLから、win-nako3.zip / mac-nako3.zip をダウンロードします。
ZIPファイルを解凍したら、nadesiko3の実行ファイルをダブルクリックしてみてください。
無事になでしこ3が動くごとが分かったら次の作業に移りましょう。

もし、Windowsで動かない場合、Chromium版のEdgeがインストールされていることを確認してください。
また、Chromium版のEdgeに加えて、WebView2のランタイムが必要です。
ページ下部より、「Evergreen Standalone Installer」を選んでインストールしてください。

 - [インストールの案内](https://developer.microsoft.com/en-us/microsoft-edge/webview2/)

# プログラムを差し替えよう

配布用プログラムには、webappフォルダが含まれています。このフォルダを開くと「main.nako3」というファイルがあります。
このファイルが起動してすぐに実行されるプログラムです。
ウィンドウサイズやタイトルを変更したい場合は「index.json」を編集してください。

macOSの場合は、nadesiko3.app/Contents/MacOS/webappにwebappディレクトリがあります。
nadesiko3.appを右クリックして「パッケージ内容を表示」を選択すると、Finderにファイルが表示されます。

# 特殊コマンドについて

特別なコマンドが利用できます。
ファイルの保存と読み込みが可能です。いずれも非同期に実行されます。

```api.nako3
# ファイルへの保存
「test.txt」に「あいうえお」をファイル保存した時には
　　「保存しました」と表示
ここまで。

# ファイルの読み込み
「test.txt」をファイル読んだ時には
　　対象を表示。
ここまで。

# ファイル一覧の取得
ファイル一覧取得した時には
　　対象を反復
　　　　それを表示。
　　ここまで。
ここまで。
```

セキュリティの問題を考慮して、ファイルを任意のフォルダに保存することはできません。
ファイルは以下のフォルダに保存されます。なお、index.jsonでappidを指定します。

```text
<ユーザーフォルダ>/.nadesiko3/<appid>
```

