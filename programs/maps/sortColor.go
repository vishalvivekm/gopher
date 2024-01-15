package main

import "fmt"

func main() {
	arr := []int{2, 2, 1, 1, 0, 1}
	//fmt.Printf("%T", arr)
	//ans := append(arr, arr...)

	sortColors(arr)
	//fmt.Println(arr)
	fmt.Println(sortArray(arr))
}

func sortColors(nums []int) {
	colorMap := make(map[int]int)
	for _, ele := range nums {
		{
			count, _ := colorMap[ele]
			colorMap[ele] = count + 1

		}
	}
	fmt.Printf("%v", colorMap)
}
func sortArray(arr []int) []int {
	for i := 0; i < len(arr); i++ {
		for j := 1; j < len(arr)-i; j++ {
			if arr[j] <= arr[j-1] {
				temp := arr[j]
				arr[j] = arr[j-1]
				arr[j-1] = temp
			}
		}
	}
	return arr
}
