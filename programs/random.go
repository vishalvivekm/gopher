package main

import (
	"bytes"
	"fmt"
)

func main() {

	// encode simple maps to json
	var product = []byte(`{"name": "chair", "quantity": 1, "price": 100}`)
	fmt.Println(product)
	buf := bytes.NewBuffer(product)
	fmt.Println(buf)

}
