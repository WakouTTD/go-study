/*
   timeパッケージの学習
*/

package lesson67

/* https://golang.org/pkg/time/

const (
        ANSIC       = "Mon Jan _2 15:04:05 2006"
        UnixDate    = "Mon Jan _2 15:04:05 MST 2006"
        RubyDate    = "Mon Jan 02 15:04:05 -0700 2006"
        RFC822      = "02 Jan 06 15:04 MST"
        RFC822Z     = "02 Jan 06 15:04 -0700" // RFC822 with numeric zone
        RFC850      = "Monday, 02-Jan-06 15:04:05 MST"
        RFC1123     = "Mon, 02 Jan 2006 15:04:05 MST"
        RFC1123Z    = "Mon, 02 Jan 2006 15:04:05 -0700" // RFC1123 with numeric zone
        RFC3339     = "2006-01-02T15:04:05Z07:00" // このフォーマットをよく使う
        RFC3339Nano = "2006-01-02T15:04:05.999999999Z07:00"
        Kitchen     = "3:04PM"
        // Handy time stamps.
        Stamp      = "Jan _2 15:04:05"
        StampMilli = "Jan _2 15:04:05.000"
        StampMicro = "Jan _2 15:04:05.000000"
        StampNano  = "Jan _2 15:04:05.000000000"
)
*/

import (
	"fmt"
	"time"
)

func main() {
	t := time.Now()

	// 2020-01-09 15:03:03.129789 +0900 JST m=+0.000073683
	fmt.Println(t)
	// 2020-01-09T15:03:03+09:00
	fmt.Println(t.Format(time.RFC3339))

	//2020 January 9 15 3 3
	fmt.Println(t.Year(), t.Month(), t.Day(),
		t.Hour(), t.Minute(), t.Second())
}
