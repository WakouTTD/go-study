package main

import (
	"io"
	"log"
	"os"
)

func loggingSettings(logFile string) {
	logfile, _ := os.OpenFile(logFile, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	multiLogFile := io.MultiWriter(os.Stdout, logfile)
	//log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
	log.SetFlags(log.Ldate | log.Ltime | log.Llongfile)
	log.SetOutput(multiLogFile)
}

// ファイルオープンのロギング
func main() {
	loggingSettings("/Users/tateda/work2019/go-study/src/lesson27/test.log")
	_, err := os.Open("./kkkd.log")
	if err != nil {
		log.Fatalln("Exit", err)
	}

}
