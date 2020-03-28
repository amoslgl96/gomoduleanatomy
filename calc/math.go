package calc

//Add returns sum of the ints stored in variadic arg numbers
func Add(numbers ...int) int {
	sum := 0
	for _,n := range numbers {
		sum += n
	}

	return sum
}