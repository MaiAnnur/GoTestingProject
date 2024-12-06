package integers

func Add(x, y int) int {
	return x + y
}

func Subtract(x, y int) int {
	return x - y
}

func Multiply(x, y int) int {
	return x * y
}

func Divide(x, y int) int {
	if y == 0 {
		panic("cannot divide by zero")
	}
	return x / y
}
