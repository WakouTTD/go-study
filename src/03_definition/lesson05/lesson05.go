package main

import "fmt"

// goの仕様
// https://golang.org/ref/spec
// numeric types
/*
uint8       the set of all unsigned  8-bit integers (0 to 255)
uint16      the set of all unsigned 16-bit integers (0 to 65535)
uint32      the set of all unsigned 32-bit integers (0 to 4294967295)
uint64      the set of all unsigned 64-bit integers (0 to 18446744073709551615)

int8        the set of all signed  8-bit integers (-128 to 127)
int16       the set of all signed 16-bit integers (-32768 to 32767)
int32       the set of all signed 32-bit integers (-2147483648 to 2147483647)
int64       the set of all signed 64-bit integers (-9223372036854775808 to 9223372036854775807)

float32     the set of all IEEE-754 32-bit floating-point numbers
float64     the set of all IEEE-754 64-bit floating-point numbers

complex64   the set of all complex numbers with float32 real and imaginary parts
complex128  the set of all complex numbers with float64 real and imaginary parts

byte        alias for uint8
rune        alias for int32
*/

// complex number 複素数
func main() {

	// pythonなどではイコールの位置を合わせないのが推奨されているが、
	// go langではイコールの位置を合わせるように推奨されている
	var (
		u8  uint8     = 255
		i8  int8      = 127
		f32 float32   = 0.3
		c64 complex64 = -5 + 12i // 複素数
	)

	// 255 127 0.3 (-5+12i)
	fmt.Println(u8, i8, f32, c64)

	// https://golang.org/pkg/fmt/
	// 型と値
	fmt.Printf("type=%T value=%v", u8, u8)

}
