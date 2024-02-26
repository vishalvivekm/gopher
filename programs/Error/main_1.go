package main

import "fmt"

func main() {

	var num1, num2 int
	//fmt.Scan(&num1)
	//fmt.Scan(&num2)
	//
	//fmt.Print(num1 + num2)

	if _, err := fmt.Scan(&num1, &num2); err != nil {
		fmt.Println("error taking two integers input")
	}
	fmt.Println(num1 + num2)
	/*_, err := fmt.Scanf("%d %d", &firstNum, &secondNum)
		if err != nil {
			return
		}
	    fmt.Println(firstNum + secondNum)*/

}
