package helper

func Sum(input []int) int {
	result := 0
	for _, v := range input {
		result += v
	}
	return result
}
