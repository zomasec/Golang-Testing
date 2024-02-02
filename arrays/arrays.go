package arrays

func Sum(nums [5]int) int {
	sum := 0
	for _, num := range nums {
		sum += num
	}
	return sum
}

