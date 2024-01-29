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

func SumAllTails(arraysToSum ...[]int) []int {
	var sums []int
	for _, numbersToSum := range arraysToSum {
		var tail []int
		if len(numbersToSum) > 0 {
			tail = numbersToSum[1:]

		} else {
			tail = []int{0}
		}
		sums = append(sums, Sum(tail))
	}

	return sums
}
