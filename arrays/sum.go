package arrays

func Sum(numbers []int) int {
	sum := 0
	for _, i := range numbers {
		sum += i
	}
	return sum
}

func SumAll(numbersToSum ...[]int) []int {
	sums := []int{}
	for _, numbers := range numbersToSum {
		sums = append(sums, Sum(numbers))
	}
	return sums

}

func SumAllTails(numbersToSumtails ...[]int) []int {
	sumOfTails := []int{}
	for _, numbers := range numbersToSumtails {
		if len(numbers) == 0 {
			sumOfTails = append(sumOfTails, 0)
		} else {
			tail := numbers[1:]
			sumOfTails = append(sumOfTails, Sum(tail))
		}
	}
	return sumOfTails
}
