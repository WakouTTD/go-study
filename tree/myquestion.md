# 今ある疑問

## goでクローラーかスクレイピングする時って何使えば良いか

## スライスって何

- どうやらarrayのことらしい
- スライスなんだけど、配列という呼び方が一般化してるのでややこしい

## エラー制御

参考：strconv.Atoi(s)では下記のようなエラー制御

```go=
// Slow path for invalid, big, or underscored integers.
i64, err := ParseInt(s, 10, 0)
if nerr, ok := err.(*NumError); ok {
    nerr.Func = fnAtoi
}
return int(i64), err
```

- lesson28,29辺りを参照

## Interfaces

## Goroutines

## Channels

## mutex(排他制御)

## goだと厳密に言うとオブジェクトではなくStruct

## goでシステムコマンド呼ぶ時ってcobraってのを使うのか

## gRPC使う

## json.marshall、json.unmarshall使う

## redis接続と操作

## mysql接続と操作

## elasticsearch接続と操作
