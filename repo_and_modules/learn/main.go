package main

import (
	"fmt"
	"github.com/pborman/uuid"
	"github.com/vishalvivekm/cryptit/decrypt"
	"github.com/vishalvivekm/cryptit/encrypt"
	_ "log"
	"os"
	_ "os/exec"
)

func main() {

	uuid := uuid.NewRandom()
	fmt.Println(uuid)

	var message string
	fmt.Print("Enter message to be encoded: ")
	_, err := fmt.Scan(&message)
	/*	if err != nil {
		log.Fatal("error taking input")
	}*/
	check(err)
	encryptedMessage := encrypt.Nimbus(message)

	fmt.Println("encrypted message:", encryptedMessage)
	fmt.Print("do you want to decrypt the message and save it into a file? yes or no? ")
	var response string
	_, err = fmt.Scanf("%s", &response) /*if err != nil {
		fmt.Errorf("Error taking reponse")
	} */
	check(err)
	if response == "yes" {
		writeToFile(encryptedMessage)
	} else {
		fmt.Println("Good Bye")
		os.Exit(3)
	}

}

func check(err error) {
	if err != nil {
		panic(err)
	}

}

func writeToFile(s string) {
	decryptedMessage := decryptmsg(s)
	file, err := os.Create("./lock.txt")
	check(err)
	defer file.Close()

	_, err = file.WriteString(decryptedMessage)
	check(err)
	fmt.Printf("file generated: success\n")

}
func decryptmsg(s string) string {

	decryptedMessage := decrypt.Nimbus(s)

	return decryptedMessage
}
