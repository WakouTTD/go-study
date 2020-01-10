package main

import (
	"fmt"

	"github.com/markcheno/go-quote"
	"github.com/markcheno/go-talib"
)

/*
RSIとは、テクニカルチャートのひとつで、「Relative Strength Index」の頭文字をとった略語です。
日本語に訳すと「相対力指数」になります。
要するに、買われすぎか、売られすぎかを判断するための指標として利用されています。
*/

// talib ターリブ　株価分析用
func main() {
	// ヤフー(米)の株価　SPYとは、「SPDR(スパイダー)S&P500ETF」の略式
	spy, _ := quote.NewQuoteFromYahoo(
		"spy", "2019-04-01", "2020-01-10", quote.Daily, true)

	// CSVのフォーマットで出力する
	fmt.Print(spy.CSV())
	// RSIというテクニカルインディケーターを使って数値を計算
	rsi2 := talib.Rsi(spy.Close, 2)
	fmt.Println(rsi2)
	mva := talib.Ema(spy.Close, 14)
	fmt.Println(mva)

}
