package main
import "fmt"

func main() {
  fmt.Println("Hello World!")
  a0 := 0;
  a1 := 1;

  var n int
  fmt.Scanf("%d", &n)
//  fmt.Println(nums, n)
  nums := make([]int, n+1)
  nums[0] = a0
  nums[1] = a1
for i:=2; i<=n; i++ {
nums[i] = nums[i-1] + nums[i-2];
    
}
fmt.Println(nums)
fmt.Printf("The %d'th term in fibo series: %d",n, nums[n-1] )
}
