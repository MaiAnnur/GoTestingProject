package arrays

func SumAll(numbersToSum ...[]int) []int {
	lengthOfNumbers := len(numbersToSum)
	sums := make([]int, lengthOfNumbers)

	for i, numbers := range numbersToSum {
		sums[i] = Sum(numbers)
	}
	//sum := 0
	// 	for i := 0; i < 5; i++ {
	// 		sum += numbers[i]
	// 	}
	// 	return sum
	// }
	//for _, number := range numbers {
	//	sum += number
	//}
	return sums
}
