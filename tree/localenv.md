[参照：Golang Mac環境構築メモ](https://qiita.com/beesk/items/b55d1b74b985524c7cf2)

##### Goのバージョンを混在させるためgoenvからGoを入れる

```bash=
brew install --HEAD goenv
```

##### 更新時

```bash=
brew upgrade --fetch-HEAD goenv
```

##### bash_profileにgoenv関連のパスを通しておく

```bash=
vim ~/.bash_profile
```

##### 以下を記載して:wq

```txt=
#goenv
export PATH="$HOME/.goenv/bin:$PATH"
export GOPATH=/Users/tateda/work2019/go-study
export PATH=$PATH:$GOPATH/bin
eval "$(goenv init -)"

```

##### bash_profileの変更を反映

```bash=
source ~/.bash_profile
```

##### goenvコマンドが叩けることの確認

```bash=
goenv -v 
```

##### ここからGoの構築
goenvでインストールできるGoのバージョンを確認

```bash=
goenv install -l
```

##### バージョン指定で入れる

```bash=
goenv install 1.13.5
```

##### 有効なバージョンに今回入れたバージョンが表示されれば成功

```bash=
goenv versions
```

##### マシン全体で入れたバージョンを有効にする

goenv global 1.13.5
goenv rehash


##### 所定のディレクトリのみで有効にしたい場合

goenv local 1.13.5



##### プログラムを作成する。ここではmain.goとしておく。
```go=
package main

import "fmt"

func main() {
    fmt.Printf("hello, world\n")
}
```

##### 実行するには、以下のようにrun + ファイル名で実行する。

```bash=
go run /Users/tateda/work2019/go-study/src/main.go 
```

[参考:](https://qiita.com/1000ch/items/e42e7c28cf7a7b798a02)

##### パッケージをインストール
`go get`によってリモートのGoパッケージをダウンロードすることが出来る。GitHubやらBitBucketやらをリモートとし、対応するバージョンシステムを利用して取得。

例えば以下のようにGitHubにあるGoのパッケージをダウンロードすると、`$GOPATH/src`にダウンロードされる。

```bash=
go get github.com/golang/lint
```

##### パッケージをインストール

`go install`によってパッケージをインストール可能。インストールしたパッケージは$GOPATH/binに配置される。パッケージ名は.goファイルのパッケージ名に基づく。

```bash=
go install pkgname
```
バイナリが$GOPATH/binに生成される。これは実行可能ファイル。
```bash=
 ./bin/pkgname
 ```

##### パッケージのビルド
`go build`によってパッケージをビルド可能。go installとは異なり、カレントディレクトリに配置される。

```bash=
go build pkgname
```

カレントディレクトリにgo install同様の実行可能ファイルが生成される。
```bash=
./pkgname
```











