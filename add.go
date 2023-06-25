package moduler

func Add(numbers []int64) (result int64) {
	for i := 0; i < len(numbers); i++ {
		result += numbers[i]
	}
	return
}
