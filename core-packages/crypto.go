package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
)

func main() {
	var password string
	fmt.Println("Enter your password")
	fmt.Scanln(&password)
	fmt.Println("password encrypted to: ", encodedWithMD5(password))
}

func encodedWithMD5(str string) string {
	var hashCode = md5.Sum([]byte(str))
	return hex.EncodeToString(hashCode[:])

}
