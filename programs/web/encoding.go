// You can edit this code!
// Click here and start typing.
package main

import (
	"encoding/base64"
	"fmt"
)

func main() {
	fmt.Println("Hello, 世界")
	name := []byte("vivek")

	encodedName := base64.URLEncoding.EncodeToString(name) // base64 URL-safe string
	fmt.Println("encodedName: ", encodedName)
	decodedName, _ := base64.URLEncoding.DecodeString(encodedName)
	// fmt.Println("decoded name: ", string(decodedName))
	fmt.Printf("decoded name: %s\n", decodedName)

}
