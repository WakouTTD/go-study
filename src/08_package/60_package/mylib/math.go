package mylib

// Average intgerのスライスの平均値を求める関数
func Average(s []int) int {
	total := 0
	for _, i := range s {
		total += i
	}
	return int(total / len(s))
}
