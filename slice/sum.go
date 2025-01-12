package slice

func Sum(numbers []int) (sum int) {
	for _, number := range numbers {
		sum += number
	}
	return
}

func SmallAll(numbers ...[]int) (sums []int) {
	for _, number := range numbers {
		sums = append(sums, Sum(number))
	}
	return
}

func SumAllTails(numbers ...[]int) (sums []int) {
	for _, number := range numbers {
		if len(number) == 0 {
			sums = append(sums, 0)
		} else {
			tail := number[1:]
			sums = append(sums, Sum(tail))
		}
	}
	return
}
