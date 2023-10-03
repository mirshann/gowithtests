package iterations

func Repeat(charecter string, repeatCount int) string {
	var repeated string
	for i := 0; i < repeatCount; i++ {
		repeated += charecter
	}
	return repeated
}