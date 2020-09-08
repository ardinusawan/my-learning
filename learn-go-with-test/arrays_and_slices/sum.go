package main

// Sum get array and return calculated sum for every value inside of it.
func Sum(numbers []int) int {
	sum := 0
	for _, number := range numbers {
		sum += number
	}
	return sum
}

// SumAllTails take slices of array and return calculated all value inside of it.
func SumAllTails(numbersToSum ...[]int) []int {
	var sums []int
	for _, numbers := range numbersToSum {
		if len(numbers) == 0 {
			sums = append(sums, 0)
		} else {
			tail := numbers[1:]
			sums = append(sums, Sum(tail))
		}
	}

	return sums
}

// TODO: https://quii.gitbook.io/learn-go-with-tests/go-fundamentals/arrays-and-slices#wrapping-up
