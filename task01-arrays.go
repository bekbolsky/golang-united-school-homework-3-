package homework

// returns an average value of array (sum / N)
func average(input [15]float32) (result float32) {
	for _, value := range input {
		result += value
	}
	return result / 15
}
