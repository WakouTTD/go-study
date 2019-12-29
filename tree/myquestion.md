# 今ある疑問

## goでクローラーかスクレイピングする時って何使えば良い？

## スライスって何？
  - どうやらarrayのことらしい

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

## goだと厳密に言うとオブジェクトではなくStruct

## goでシステムコマンド呼ぶ時ってcobraってのを使うのか

## gRPC使う

## json.marshall、json.unmarshall使う

## redis

## mysql

## elasticsearch
