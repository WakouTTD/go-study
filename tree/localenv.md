# ローカル環境構築

[参照：Golang Mac環境構築メモ](https://qiita.com/beesk/items/b55d1b74b985524c7cf2)

## Goのバージョンを混在させるためgoenvからGoを入れる

```bash=
brew install --HEAD goenv
```

## 更新時

```bash=
brew upgrade --fetch-HEAD goenv
```

## bash_profileにgoenv関連のパスを通しておく

```bash=
vim ~/.bash_profile
```

## 以下を記載して:wq

```txt=
#goenv
export PATH="$HOME/.goenv/bin:$PATH"
# envによるバージョン切替
eval "$(goenv init -)"
#
export GOPATH=$HOME/work2020/go-study
export PATH=$PATH:$GOPATH
export GOBIN=$GOPATH/bin
```

## bash_profileの変更を反映

```bash=
source ~/.bash_profile
```

## goenvコマンドが叩けることの確認

```bash=
goenv -v
```

## ここからGoの構築

goenvでインストールできるGoのバージョンを確認

```bash=
goenv install -l
```

## バージョン指定で入れる

```bash=
goenv install 1.13.5
```

## 有効なバージョンに今回入れたバージョンが表示されれば成功

```bash=
goenv versions
```

## マシン全体で入れたバージョンを有効にする

goenv global 1.13.5
goenv rehash

## 所定のディレクトリのみで有効にしたい場合

```bash=
goenv local 1.13.5
```

[参考:vscodeを使うなら](https://qiita.com/sasaron397/items/ec285b64607c1e7662e0)

## プログラムを作成する

ここではmain.goとしておく

```go=
package main

import "fmt"

func main() {
    fmt.Printf("hello, world\n")
}
```

## 実行するには、以下のようにrun + ファイル名で実行する

```bash=
go run /Users/tateda/work2020/go-study/src/main.go
```

[参考:](https://qiita.com/1000ch/items/e42e7c28cf7a7b798a02)

## パッケージをインストール

`go get`によってリモートのGoパッケージをダウンロードすることが出来る。GitHubやらBitBucketやらをリモートとし、対応するバージョンシステムを利用して取得。

例えば以下のようにGitHubにあるGoのパッケージをダウンロードすると、`$GOPATH/src`にダウンロードされる。

```bash=
go get github.com/golang/lint
```

## パッケージインストール

`go install`によってパッケージをインストール可能。インストールしたパッケージは$GOPATH/binに配置される。パッケージ名は.goファイルのパッケージ名に基づく。

```bash=
go install pkgname
```

バイナリが$GOPATH/binに生成される。これは実行可能ファイル。

```bash=
 ./bin/pkgname
 ```

## パッケージのビルド

`go build`によってパッケージをビルド可能。go installとは異なり、カレントディレクトリに配置される。

```bash=
go build pkgname
```

カレントディレクトリにgo install同様の実行可能ファイルが生成される。

```bash=
./pkgname
```

## VS Code上でデバッグができるようにdelevをインストール

brewでもインストールできますが、認証が面倒なのでgoコマンドでインストールします。

```bash=
go get -u -v github.com/derekparker/delve/cmd/dlv
```

```bash=
dlv version
```

[参考:デバッグ](https://dev.classmethod.jp/go/visual-studio-code-golang-debug/)

```bash=
export GOENV_DISABLE_GOPATH=1
```

## go-sqlite3を使うためのMac環境設定

```bash=
brew install sqlite
install gcc (Xcode)
go get github.com/mattn/go-sqlite3
(export CGO_ENABLED=1)
```

## sqlite実行

- ターミナルで'sqlite3'と打つだけ
- 抜ける時は'.exit'と打つ

```bash=
23:45:35 go-study $ sqlite3
SQLite version 3.28.0 2019-04-15 14:49:49
Enter ".help" for usage hints.
Connected to a transient in-memory database.
Use ".open FILENAME" to reopen on a persistent database.
sqlite> .exit
23:49:36 go-study $
```

## xcodeのコマンドラインツールをインストール

```bash=
xcode-select --install
```

## サーバーと立ち上げようとして、already used errorになったら

- lsofコマンドで確認して殺しましょう

```bash=
04:03:51 mattn $ sudo lsof -i:8080
Password:
COMMAND    PID   USER   FD   TYPE             DEVICE SIZE/OFF NODE NAME
Google     800 tateda   28u  IPv6 0x34f0cf3bb7629283      0t0  TCP localhost:49566->localhost:http-alt (ESTABLISHED)
87_temple 8958 tateda    3u  IPv6 0x34f0cf3bb762ab03      0t0  TCP *:http-alt (LISTEN)
87_temple 8958 tateda    5u  IPv6 0x34f0cf3bb7628023      0t0  TCP localhost:http-alt->localhost:49566 (ESTABLISHED)
04:45:27 mattn $ kill 8958
04:47:08 mattn $ sudo lsof -i:8080
04:47:11 mattn $
```
