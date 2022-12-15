package firstassessment

import (
	"fmt"
	"math"
)

func Hello() {
	fmt.Println("Hello from first assesment!")
}

func SolutionOne() {
	var year int
	fmt.Print("Enter the year: ")
	fmt.Scanf("%d", &year)
	if isLeapYear(year) {
		fmt.Println("The entered year is a Leap year")
	} else {
		fmt.Println("The entered year is not a leap year")
	}
}

func isLeapYear(year int) bool {
	return year%4 == 0 && year%100 != 0 || year%400 == 0
}

func SolutionTwo() {
	var num int
	fmt.Print("Enter the number: ")
	fmt.Scanf("%d", &num)

	if num%2 == 0 {
		fmt.Printf("Entered number %d is an even number", num)
	} else {
		fmt.Printf("Entered number %d is an odd number", num)
	}

	switch {
	case num < 0:
		fmt.Printf("\nEntered number %d is a negative number", num)
	case num > 0:
		fmt.Printf("\nEntered number %d is a positive number", num)
	default:
		fmt.Printf("\nEntered number %d is Zero", num)
	}

	fmt.Printf("\nEntered number has %d digits", countDigit(num))

}

func countDigit(number int) int {
	return int(math.Floor(math.Log10(math.Abs(float64(number))) + 1))
}

func SolutionThree() {
	var num1, num2 int
	fmt.Print("Enter two numbers: ")
	fmt.Scan(&num1, &num2)
	fmt.Printf("\nSum of %d and %d = %d", num1, num2, addition(num1, num2))
	fmt.Printf("\nSubtraction of %d and %d = %d", num1, num2, subtraction(num1, num2))
	fmt.Printf("\nMultiplication of %d and %d = %d", num1, num2, multiplication(num1, num2))
	fmt.Printf("\nDivision of %d and %d = %f", num1, num2, division(num1, num2))
}

func addition(a, b int) int {
	return a + b
}

func subtraction(a, b int) int {
	return a - b
}

func multiplication(a, b int) int {
	return a * b
}
func division(a, b int) float64 {
	return float64(a) / float64(b)
}

func SolutionFour() {
	var str string
	fmt.Print("Enter the string: ")
	fmt.Scanln(&str)
	fmt.Printf("\nLength of the string = %d", len(str))
	fmt.Printf("\nLast character in the string = %v", str[len(str)-1:])
	fmt.Println("\nASCII value of each character in the string: ")
	printASCIIValue(str)
	fmt.Printf("\nReversed string of %s = %s", str, reverseString(str))
}

func printASCIIValue(str string) {
	for _, value := range str {
		fmt.Printf("\nASCII value of %c = %d", value, int(value))
	}
}

func reverseString(str string) (result string) {
	for _, value := range str {
		result = string(value) + result
	}
	return
}

func SolutionFive() {
	var str1, str2 string
	fmt.Print("Enter two strings: ")
	fmt.Scan(&str1, &str2)
	fmt.Printf("Concatenated string = %s", str1+str2)
}
