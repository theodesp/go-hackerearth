package basic_programming

func Factorial(x int) int  {
	if x == 0 {
		return 1
	}
	return x * Factorial(x-1)
}


func FactorialTail(x int) int  {
	if x == 0 {
		return 1
	}
	return loop(1, x)
}

func loop(curr, x int) int  {
	if x == 0 {
		return curr
	}
	return loop(curr * x, x-1)
}

