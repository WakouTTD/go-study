# go-study


[参照：Golang Mac環境構築メモ](https://qiita.com/beesk/items/b55d1b74b985524c7cf2)

##### Goのバージョンを混在させるためgoenvからGoを入れる

```bash=
brew install --HEAD goenv
```

更新時

```bash=
brew upgrade --fetch-HEAD goenv
```

bash_profileにgoenv関連のパスを通しておく

```bash=
vim ~/.bash_profile
```

以下を記載して:wq

```txt=
#goenv
export PATH="$HOME/.goenv/bin:$PATH"
export GOPATH=$HOME/go
eval "$(goenv init -)"
```

bash_profileの変更を反映

```bash=
source ~/.bash_profile
```

goenvコマンドが叩けることの確認

```bash=
goenv -v 
```

ここからGoの構築
goenvでインストールできるGoのバージョンを確認

```bash=
goenv install -l
```

バージョン指定で入れる

```bash=
goenv install 1.13.5
```

有効なバージョンに今回入れたバージョンが表示されれば成功

```bash=
goenv versions
```

マシン全体で入れたバージョンを有効にする

goenv global 1.13.5
goenv rehash


所定のディレクトリのみで有効にしたい場合

goenv local 1.13.5


