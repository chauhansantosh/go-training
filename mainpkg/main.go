package main

import (
	"fmt"

	"github.com/chauhansantosh/go-training/firstassessment"
	"github.com/chauhansantosh/go-training/fourthassessment"
	"github.com/chauhansantosh/go-training/secondassessment"
	"github.com/chauhansantosh/go-training/thirdassessment"
)

func main() {
	fmt.Println("Hello from main!")
	firstassessment.Hello()
	//firstassessment.SolutionOne()
	//firstassessment.SolutionTwo()
	//firstassessment.SolutionThree()
	//firstassessment.SolutionFour()
	//firstassessment.SolutionFive()

	secondassessment.Hello()
	//secondassessment.SolutionOne()
	//secondassessment.SolutionTwo()
	//secondassessment.SolutionThree()
	//secondassessment.SolutionFour()
	//secondassessment.SolutionFive()

	thirdassessment.Hello()
	//thirdassessment.SolutionOne()

	//Output of SolutionOne
	/*Enter the size of the array: 5
	Enter the number: 1
	Enter the number: 2
	Enter the number: 3
	Enter the number: 6
	Enter the number: 7
	Even elements of the array are:  [2 6]
	Odd elements of the array are:  [1 3 7] */

	//thirdassessment.SolutionTwo()

	//Output of SolutionTwo
	/* Enter the size of the array: 5
	Enter the number: 1
	Enter the number: 2
	Enter the number: 3
	Enter the number: 4
	Enter the number: 5
	Enter the search element: 3
	Searched number exists at index:  2 */

	//thirdassessment.SolutionThree()

	//Output of SolutionThree
	/* 	Enter the size of the array: 5
	   	Enter the number: 1
	   	Enter the number: 2
	   	Enter the number: 3
	   	Enter the number: 4
	   	Enter the number: 5
	   	Sum of the elements =  15
	   	Multiplication of the elements =  120 */

	//thirdassessment.SolutionFour()

	/* Enter the size of the slice : 5
	Enter the elements of the slice: 1
	Enter the elements of the slice: 3
	Enter the elements of the slice: 2
	Enter the elements of the slice: 4
	Enter the elements of the slice: 7
	Sorted slice =  [1 2 3 4 7]
	Enter the search element: 4
	Found the number 4 at index 3 */

	//thirdassessment.SolutionFive()

	/* Enter employee id (enter -999 to stop) : 123
	Enter employee name for id 123: santosh
	Enter employee id (enter -999 to stop) : 456
	Enter employee name for id 456: deepika
	Enter employee id (enter -999 to stop) : 789
	Enter employee name for id 789: aman
	Enter employee id (enter -999 to stop) : -999
	The list of employees:
	Employee ID = 123, Employee Name = santosh
	Employee ID = 456, Employee Name = deepika
	Employee ID = 789, Employee Name = aman */

	fourthassessment.Hello()

	/* santoshAccount := fourthassessment.BankAccount{}
	santoshAccount.Initialization("Santosh", 51691852, "Savings", 10000.5)
	santoshAccount.DisplayBalance()
	santoshAccount.Withdraw()
	santoshAccount.Deposit()

	deepikaAccount := fourthassessment.BankAccount{}
	deepikaAccount.Initialization("Deepika", 51691853, "Current", 20000.5)
	deepikaAccount.DisplayBalance()
	deepikaAccount.Withdraw()

	amanAccount := fourthassessment.BankAccount{}
	amanAccount.Initialization("Aman", 51691854, "Saving", 30000.0)
	amanAccount.DisplayBalance()
	amanAccount.Withdraw() */

	fifthassessment.Hello()

	var account1, account2, account3 fifthassessment.BankAccountInterface
	account1 = &fifthassessment.BankAccount{}
	account1.Init("Santosh", 51691852, "Current", 10000.5)
	account1.BalanceCheck()
	err := account1.Withdraw(1000)
	if err != nil {
		fmt.Println(err)
	}

	account2 = &fifthassessment.BankAccount{}
	account2.Init("Jharana", 51691853, "Savings", 20000.5)
	account2.BalanceCheck()
	account2.Deposit(5000)
	err = account2.Withdraw(10000)
	if err != nil {
		fmt.Println(err)
	}

	account3 = &fifthassessment.BankAccount{}
	account3.Init("Aman", 51691854, "SV", 30000.0)
	account3.BalanceCheck()
	err = account3.Withdraw(5000)
	if err != nil {
		fmt.Println(err)
	}
	err = account3.Withdraw(30000)
	if err != nil {
		fmt.Println(err)
	}

	/* Hello from fifth assessment!
	Balance for the Santosh account = 10000.500000
	Amount withdrawn 1000
	Balance for the Santosh account = 9000.500000
	Balance for the Jharana account = 20000.500000
	Amount deposited 5000
	Balance for the Jharana account = 25000.500000
	Amount withdrawn 10000
	Balance for the Jharana account = 15000.500000
	Balance for the Aman account = 30000.000000
	Amount withdrawn 5000
	Balance for the Aman account = 25000.000000
	Amount withdrawn 30000
	insufficient fund */
}
