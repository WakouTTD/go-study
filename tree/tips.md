# メモ

- フォーマット

フォーマット後の表示

```bash=
gofmt sample.go
```

実際に書き換えてもらう場合は、オプションの -w　をつける

```bash=
gofmt -w sample.go
```

このディレクトリ配下の全てのtest実行
```bach=
go test -v ./..
```
