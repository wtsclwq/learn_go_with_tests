package iteration

func Repeat(times int, character string) string {
	repeated := ""
	for i := 0; i < times; i++ {
		repeated += character
	}
	return repeated
}
