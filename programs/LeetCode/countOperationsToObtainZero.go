func countOperations(num1 int, num2 int) int {
    i := 0
    for {
    if num1 == 0 || num2 == 0 {
        break
    }
    if num1 >= num2 {
        num1 = num1 - num2
        i++
    } else {
        num2 = num2 - num1
        i++
    }
    }
    return i
}
