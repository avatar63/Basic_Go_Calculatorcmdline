package main

import (
	"fmt"
)

func main() {

	var a, b float64

	fmt.Println("Welcome to GoCalc !!!")

	fmt.Println("Enter a number: ")
	fmt.Scan(&a)
	fmt.Println("Enter another number: ")
	fmt.Scan(&b)

	fmt.Println(menu(a, b))

}

func menu(a, b float64) string {
	var choice string
	var flag bool = true
	for flag {

		fmt.Println("Select an operation out of the following: ")

		fmt.Println("1. Addition\n2.Subtraction\n3.Multiplication\n4.Division\n5.Enter New Numbers\n6.Exit")
		fmt.Scan(&choice)

		switch choice {

		case "1":
			fmt.Println("Addition Function")
			fmt.Println("ANSWER ----------------->", Addition(int(a), int(b)))

		case "2":
			fmt.Println("Subtraction Function")
			fmt.Println("ANSWER ----------------->", Subtraction(int(a), int(b)))

		case "3":
			fmt.Println("Multiplication Function")
			fmt.Println("ANSWER ----------------->", Multiplication(a, b))
		case "4":
			fmt.Println("Division Function")
			fmt.Println("ANSWER ----------------->", Division(a, b))
		case "5":
			fmt.Println("Enter a number: ")
			fmt.Scan(&a)
			a = a
			fmt.Println("Enter another number: ")
			fmt.Scan(&b)
			b = b
		case "6":
			flag = false

		default:
			fmt.Println("Error")

		}
	}

	return "Thanks for using GoCalc !!!"
}

func Addition(num1, num2 int) int {

	var sum int
	sum = num1 + num2
	return sum
}

func Subtraction(num1, num2 int) int {

	var ans int
	ans = num1 - num2
	return ans
}

func Multiplication(num1, num2 float64) float64 {

	var ans = num1 * num2
	return ans
}

func Division(num1, num2 float64) float64 {

	var ans = num1 / num2
	return ans
}
