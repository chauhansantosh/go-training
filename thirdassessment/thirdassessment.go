package thirdassessment

import (
	"fmt"
	"sort"
)

func Hello() {
	fmt.Println("Hello from third assessment!")
}

func SolutionOne() {
	arrLength := 0
	var evenSlice, oddSlice []int
	fmt.Print("Enter the size of the array: ")
	fmt.Scanln(&arrLength)
	userArr := make([]int, arrLength)
	for i := 0; i < arrLength; i++ {
		fmt.Print("Enter the number: ")
		fmt.Scanln(&userArr[i])
	}
	for i := 0; i < arrLength; i++ {
		if userArr[i]%2 == 0 {
			evenSlice = append(evenSlice, userArr[i])
		} else {
			oddSlice = append(oddSlice, userArr[i])
		}
	}
	fmt.Println("Even elements of the array are: ", evenSlice)
	fmt.Println("Odd elements of the array are: ", oddSlice)
}

func SolutionTwo() {
	arrLength := 0
	fmt.Print("Enter the size of the array: ")
	fmt.Scanln(&arrLength)
	userArr := make([]int, arrLength)
	for i := 0; i < arrLength; i++ {
		fmt.Print("Enter the number: ")
		fmt.Scanln(&userArr[i])
	}
	var searchNumber int
	fmt.Print("Enter the search element: ")
	fmt.Scanln(&searchNumber)

	for index, value := range userArr {
		if value == searchNumber {
			fmt.Println("Searched number exists at index: ", index)
			return
		}
	}
	fmt.Println("Searched number does not exist")
}

func SolutionThree() {
	arrLength := 0
	fmt.Print("Enter the size of the array: ")
	fmt.Scanln(&arrLength)
	userArr := make([]int, arrLength)
	for i := 0; i < arrLength; i++ {
		fmt.Print("Enter the number: ")
		fmt.Scanln(&userArr[i])
	}
	var sum, multiplication int = 0, 1
	for _, value := range userArr {
		sum += value
		multiplication *= value
	}
	fmt.Println("Sum of the elements = ", sum)
	fmt.Println("Multiplication of the elements = ", multiplication)
}

func SolutionFour() {
	sliceLength := 0
	fmt.Print("Enter the size of the slice : ")
	fmt.Scanln(&sliceLength)
	userSlice := make([]int, sliceLength)
	for i := 0; i < sliceLength; i++ {
		fmt.Print("Enter the elements of the slice: ")
		fmt.Scanln(&userSlice[i])
	}
	sort.Slice(userSlice, func(p, q int) bool { return userSlice[p] < userSlice[q] })
	fmt.Println("Sorted slice = ", userSlice)

	searchNumber := 0
	fmt.Print("Enter the search element: ")
	fmt.Scanln(&searchNumber)
	binarySearch(userSlice, searchNumber, 0, sliceLength-1)
}

func SolutionFive() {
	var empId int
	var name string
	var empMap = map[int]string{}
	for {
		fmt.Print("Enter employee id (enter -999 to stop) : ")
		fmt.Scanln(&empId)
		if empId != -999 {
			fmt.Printf("Enter employee name for id %d: ", empId)
			fmt.Scanln(&name)
			empMap[empId] = name
		} else {
			break
		}
	}
	fmt.Println("The list of employees:")
	for id, name := range empMap {
		fmt.Printf("Employee ID = %d, Employee Name = %s\n", id, name)
	}
}

func binarySearch(mySlice []int, searchElement int, low int, high int) {
	if low > high {
		fmt.Println("The number does not exist in slice")
	} else {
		mid := (low + high) / 2
		if searchElement == mySlice[mid] {
			fmt.Printf("Found the number %d at index %d", searchElement, mid)
		} else if searchElement > mySlice[mid] { // searchElement is on the right side
			binarySearch(mySlice, searchElement, mid+1, high)
		} else { // searchElement is on the left side
			binarySearch(mySlice, searchElement, low, mid-1)
		}
	}
}
