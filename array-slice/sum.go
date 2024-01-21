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
	for i := 0; i < len(arraysToSum); i++ {
		arrayToSum := arraysToSum[i]
		sums[i] = 0
		for _, number := range arrayToSum {
			sums[i] += number
		}
	}
	return sums

}
