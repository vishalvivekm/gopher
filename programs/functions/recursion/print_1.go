func print(n int) {
        if n == 0 {
                return
        }
        print(n - 1)
        fmt.Print(n)

}

func main() {
        print(5)
}
