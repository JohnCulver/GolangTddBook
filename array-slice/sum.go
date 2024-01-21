package arrayslice

func Sum(numbers []int) int {
	total := 0
	for _, number := range numbers {
		total += number
	}
	return total
}

func SumAll(arraysToSum ...[]int) []int {

	sums := []int{}

	for _, numberArray := range arraysToSum {
		sums = append(sums, Sum(numberArray))
	}

	return sums

}
