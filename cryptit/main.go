package main

import (
"github.com/vishalvivekm/cryptit/encrypt"
"github.com/vishalvivekm/cryptit/decrypt"
"fmt"
)

func main() {
encryptedStr := encrypt.Nimbus("VivekVishal")
fmt.Println(encryptedStr) // we can't import an identifier if it doesn't start with capital letter, like here, it must be Nimbus, not nimbus

fmt.Println(decrypt.Nimbus(encryptedStr))

}

// not nimbus, but Nimbus
