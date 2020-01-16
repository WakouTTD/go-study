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

- ちょっとわかった
- その後、忘れたので復讐が必要

## Goroutines

- ちょっとわかった
- その後、忘れたので復讐が必要

## Channels

- ちょっとわかった
- その後、忘れたので復讐が必要

## mutex(排他制御)

- ちょっとわかった
- その後、忘れたので復讐が必要

## Unit test

- 標準のtestだと下記のようになる

```go=
import "testing"

// このテスト実行しないよって時は trueにする
var Debug bool = false

func TestAverage(t *testing.T) {
    if Debug {
        t.Skip("Skip Reason")
    }
    v := Average([]int{1, 2, 3, 4, 5})
    if v != 3 {
        t.Error("Expected 3, got", v)
    }
}
```

- もしくはサードパーティのGinkgoを使うか、もっと良いtestフレームワークを探すか gomegaとか

## goでシステムコマンド呼ぶ時ってcobraってのを使うのか

## gRPC使う

## GraphQL使う

## json.marshall、json.unmarshall使う

## redis接続と操作

## postgre接続と操作

## elasticsearch接続と操作

## どこからも生えてない関数って使わない方が良さそう

- lesson39と40をやってみた感想
