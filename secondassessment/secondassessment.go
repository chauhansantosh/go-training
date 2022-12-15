package secondassessment

import (
	"fmt"
	"math"
	"strconv"
)

func Hello() {
	fmt.Println("Hello from second assesment!")
}

func SolutionOne() {
	var num int
	fmt.Print("Enter the number: ")
	fmt.Scanf("%d", &num)
	checkPrime(num)
	if isPalindrome(num) {
		fmt.Printf("\n%d is Palimdrome", num)
	} else {
		fmt.Printf("\n%d is not Palimdrome", num)
	}
}

func checkPrime(num int) {
	if num < 2 {
		fmt.Println("\nNumber must be greater than or equal to 2.")
		return
	}
	sq_root := int(math.Sqrt(float64(num)))
	for i := 2; i <= sq_root; i++ {
		if num%i == 0 {
			fmt.Printf("\n%d is a Non Prime Number", num)
			return
		}
	}
	fmt.Printf("%d is a Prime Number", num)
}

func isPalindrome(num int) bool {
	str := strconv.Itoa(num)
	lastIndex := len(str) - 1
	for i := 0; i < lastIndex/2 && i < (lastIndex-i); i++ {
		if str[i] != str[lastIndex-i] {
			return false
		}
	}
	return true
}

func SolutionTwo() {
	var baseNum, exponent int
	fmt.Printf("Enter a base number and its exponent: ")
	fmt.Scan(&baseNum, &exponent)
	squareRoot := math.Sqrt(float64(baseNum))
	fmt.Printf("\nSquare root of the base number %d = %f", baseNum, squareRoot)
	power := math.Pow(float64(baseNum), float64(exponent))
	fmt.Printf("\n%d^%d = %f", baseNum, exponent, power)
}

func SolutionThree() {
	var numArray1, numArray2, arraySum, arrayMultiplication [3]int
	for i := 0; i < 3; i++ {
		fmt.Printf("Enter %dth element of first array: ", i)
		fmt.Scanln(&numArray1[i])
	}
	for i := 0; i < 3; i++ {
		fmt.Printf("Enter %dth element of second array: ", i)
		fmt.Scanln(&numArray2[i])
	}
	for i := 0; i < 3; i++ {
		arraySum[i] = numArray1[i] + numArray2[i]
		arrayMultiplication[i] = numArray1[i] * numArray2[i]
	}
	fmt.Println("Array Sum:", arraySum)
	fmt.Println("Array Multiplication:", arrayMultiplication)
}

func SolutionFour() {
	var arraySize int
	fmt.Printf("Enter size of your array: ")
	fmt.Scanln(&arraySize)
	var numArray = make([]int, arraySize)
	var sumArray int
	for i := 0; i < arraySize; i++ {
		fmt.Printf("Enter %dth element: ", i)
		fmt.Scanln(&numArray[i])
		sumArray += numArray[i]
	}
	fmt.Println("Sum of array elements: ", sumArray)
	fmt.Println("Original array: ", numArray)
	rvereseArray(numArray, 0, arraySize-1)
	fmt.Println("Reversed array: ", numArray)
}

func rvereseArray(arr []int, start, end int) {
	if start >= end {
		return
	}
	temp := arr[start]
	arr[start] = arr[end]
	arr[end] = temp
	rvereseArray(arr, start+1, end-1)
}

func SolutionFive() {
	var arraySize int
	fmt.Printf("Enter size of your array: ")
	fmt.Scanln(&arraySize)
	var numArray = make([]int, arraySize)
	for i := 0; i < arraySize; i++ {
		fmt.Printf("Enter %dth element: ", i)
		fmt.Scanln(&numArray[i])
	}
	fmt.Println("Length of the array: ", len(numArray))
	fmt.Println("First element of the array: ", numArray[0])
	fmt.Println("Last element of the array: ", numArray[arraySize-1])
}
