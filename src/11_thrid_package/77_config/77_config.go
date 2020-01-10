package main

import (
	"fmt"

	"github.com/ini"
	//"gopkg.in/ini.v1"
)

// ConfigList どんなstructに入れたいか
type ConfigList struct {
	Port      int
	DbName    string
	SQLDriver string
}

// Config だよ
var Config ConfigList

// init main関数が呼ばれる前に実行される
func init() {
	cfg, _ := ini.Load("/Users/tateda/work2019/go-study/src/11_thrid_package/77_config/config.ini")
	Config = ConfigList{
		Port: cfg.Section("web").Key("port").MustInt(),
		// confから値が取れなかった場合、MustStringはデフォルト値、Stringは空、を設定
		DbName:    cfg.Section("db").Key("name").MustString("example.sql"),
		SQLDriver: cfg.Section("db").Key("driver").String(),
	}
}

func main() {
	fmt.Printf("%T %v\n", Config.Port, Config.Port)
	fmt.Printf("%T %v\n", Config.DbName, Config.DbName)
	fmt.Printf("%T %v\n", Config.SQLDriver, Config.SQLDriver)
}
