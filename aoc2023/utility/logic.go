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

// Calculates the Least Common Multiple for 2 numbers. This assumes that
// LCM(a, b) = ab/GCD(a,b)) = a/GCD(a,b)*b. There's a good discussion on this
// at: https://stackoverflow.com/a/3154503
func LCM(a int, b int) int {
	lcm := a / RecursiveGCD(a, b) * b
	return lcm
}

// This is a recursive implementation of GCD in Go borrowed under the MIT license at:
// https://github.com/TheAlgorithms/Go/blob/master/LICENSE
func RecursiveGCD(a, b int) int {
	if b == 0 {
		return a
	}
	return RecursiveGCD(b, a%b)
}
