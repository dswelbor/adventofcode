package utility

// Add up all the calibration numbers and return their sum
func SumNumbers(numbers *[]int) int {
	sum := 0
	for _, num := range *numbers {
		sum += num
	}
	return sum
}
