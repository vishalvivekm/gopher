package main

import "fmt"

func main() {
	var T int
	fmt.Scanf("%d", &T)

	for i := 1; i <= T; i++ {

		arr := make([]int, 5)
		sum := 0

		for j := 0; j < 5; j++ {
			fmt.Scanf("%d", &arr[j])
			sum = sum + arr[j]
		}

		if sum >= 4 {
			fmt.Println("YES")
		} else {
			fmt.Println("NO")
		}
	}
}

