package main

import (
	"fmt"
)

func Devide(num1, num2 float64) (float64, error) {

	if num2 == 0 {
		err := fmt.Errorf("can't divide %.2f with %.2f", num1, num2)
		return 0, err
	}

	return num1 / num2, nil

}

func main() {
	value, err := Devide(2, 0)
	if err != nil {
		panic(err)
	}
	fmt.Println(value)
}
