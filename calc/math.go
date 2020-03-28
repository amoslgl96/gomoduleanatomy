package calc

import "errors"

//Add returns sum of the ints stored in variadic arg numbers
func Add(numbers ...int) (error,int) {
	sum := 0
	if len(numbers) < 2 {
		return errors.New("Provided less than 2 numbers"),sum
	}
	for _,n := range numbers {
		sum += n
	}

	return nil,sum
}