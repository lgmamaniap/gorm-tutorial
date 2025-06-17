package main

func sum(a, b int) int {
	return a + b
}

func multiply(a, b int) int {
	return a * b
}

func divide(a, b int) int {
	if b == 0 {
		return 0 // Handle division by zero
	}
	return a / b
}

func main() {
	result := sum(3, 5)
	println("The sum is:", result) // Output: The sum is: 8
	println("Sum function executed successfully")

	result = multiply(3, 5)
	println("The product is:", result) // Output: The product is: 15
	println("Multiply function executed successfully")

	result = divide(10, 2)
	println("The division result is:", result) // Output: The division result is: 5
	println("Divide function executed successfully")
}
