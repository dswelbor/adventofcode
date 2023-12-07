package utility

// Add up all the calibration numbers and return their sum
func SumNumbers(numbers *[]int) int {
	sum := 0
	for _, num := range *numbers {
		sum += num
	}
	return sum
}

// Takes a list of ints and multiples all elements together
func MultipleNumbers(numbers *[]int) int {
	product := 1
	for _, num := range *numbers {
		product *= num
	}
	return product
}
