// Given a 0-indexed m x n integer matrix matrix, create a new 0-indexed matrix called answer. Make answer equal to matrix, then replace each element with the value -1 with the maximum element in its respective column.
// https://leetcode.com/problems/modify-the-matrix/description/

package main

import "fmt"

func main() {
	/* an array with 5 rows and 2 columns*/
	//   var a = [5][2]int{ {0,0}, {1,2}, {2,4}, {3,6},{4,8}}
	//   var i, j int
	//   fmt.Println(len(a[0]))
	/* output each array element's value */
	//   for  i = 0; i < 5; i++ {
	//       for j = 0; j < 2; j++ {
	//          fmt.Printf("a[%d][%d] = %d\n", i,j, a[i][j] )
	//       }
	//   }
	fmt.Println(modifiedMatrix([][]int{
		{3, -1, 4},
		{5, 2, -1}},
	))
}

func modifiedMatrix(matrix [][]int) [][]int {

	m := len(matrix)
	n := len(matrix[0])
	ans := matrix

	var j int
	fmt.Println(m, n)

	for i := 0; i <= n-1; i++ {
		max := matrix[0][i]
		for j = 0; j <= m-1; j++ {
			if max < ans[j][i] {
				max = ans[j][i]
			}
		}
		fmt.Println(max)
		// j++
		for q := 0; q <= m-1; q++ {

			if ans[q][i] == -1 {
				ans[q][i] = max
			}

		}
		//i++
	}
	return ans

}
