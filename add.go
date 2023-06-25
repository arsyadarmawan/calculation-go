package moduler

func Add(numbers []int64) (result int64) {
	for _, v := range numbers {
		result += v
	}
	return
}
