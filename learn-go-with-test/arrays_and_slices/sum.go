package main

// Sum get array and return calculated sum for every value inside of it.
func Sum(numbers [5]int) int {
	sum := 0
	for _, number := range numbers {
		sum += number
	}
	return sum
}
