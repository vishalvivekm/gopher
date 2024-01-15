package main

import "fmt"

func main() {
	var sum = 4 + "int"
	fmt.Println(sum)
	// y := [3]int{10, 20, 30}
	// py := &y
	// fmt.Printf("%T %v \n", py, *py)
	// output: ?

	// y := [3]int{10, 20, 30}
	// fmt.Printf("%v \n", y)
	// (*&y)[0] = 100
	// fmt.Printf("%v \n",y)

	//s := 100
	//var ptr *string = &s
	//fmt.Println(s)
	//*ptr += 100
	//fmt.Println(s)
	// error: cannot use &s (type *int) as type *string in assignment

	/*
		s := "100"
		var ptr *string = &s
		fmt.Println(s)
		*ptr += "0"
		fmt.Println(s) 100 <br> 1000
	*/

	/*
		s := "hello"
		var ptr *string = &s
		fmt.Println(s)
		*ptr += strings.ToUpper(s)
		fmt.Println(s)
	*/
}

/*
	var sum1 int
	if _, err := fmt.Scanln(&sum1); err != nil {
		//log.Fatal(err)
		fmt.Println(err)
		return
	}
*/
