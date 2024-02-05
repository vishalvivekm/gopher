package main

import "fmt"

type Expense struct {
	expenseName string
	amount float64
	date string
}

func Total (expense []Expense) float64 {
	var sum float64
   for _, ele := range expense {
	   sum += ele.amount
   }
   return sum
}

func (expense *Expense) getName() string {
return expense.expenseName
}

func main() {
	expenses := []Expense{
		Expense{"Grocery", 50.0, "2022-01-01"},
		Expense{"Gas", 30.0, "2022-01-02"},
		Expense{"Restaurant", 40.0, "2022-01-03"},
	}

	fmt.Println(Total(expenses))
	fmt.Println(expenses[0].getName())
}

/*

package main

import "fmt"

type Expense struct {
	name string
	amount float64
	date string
}

func Total(expense []Expense) float64 {
	var sum float64
	for _, ele := range expense {
		sum += ele.amount
	}
	return sum
}

func (expense *Expense) getName() string {
	return expense.name
}

func main() {
	expenses := []Expense{
		Expense{"Grocery", 50.0, "2022-01-01"},
		Expense{"Gas", 30.0, "2022-01-02"},
		Expense{"Restaurant", 40.0, "2022-01-03"},
	}

	fmt.Println(Total(expenses))
	fmt.Println(expenses[0].getName())
}

*/
