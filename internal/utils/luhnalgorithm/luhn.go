package luhnalgorithm

func IsValid(number int) bool {
	return (number%10+checkSum(number/10))%10 == 0
}

func checkSum(number int) int {
	var sum int
	for i := 0; number > 0; i++ {
		cur := number % 10

		if i%2 == 0 {
			cur = cur * 2
			if cur > 9 {
				cur = cur - 9
			}
		}

		sum = sum + cur
		number = number / 10
	}
	return sum % 10
}
