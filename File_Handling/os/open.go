package main

import (
	"fmt"
	"os"
)

// to read in chunk: os.Open, opens file in Read-only mode, refer to os.OpenFile for more modes
func main() {
	path := "../temp.txt"

	file, err := os.Open(path)
	buffer := make([]byte, 4)
	if err != nil {
		fmt.Println(err)
	}
	for {
		n, err := file.Read(buffer)

		//switch err {
		//case nil:
		//	fmt.Println(string(buffer[:n]))
		//case io.EOF:
		//	//fmt.Println(string(buffer[:]))
		//	fmt.Println("done reading the string")
		//	return
		//default:
		//	fmt.Println("error occured", err)
		//}
		//if err == nil {
		//	fmt.Println(string(buffer[:n]))
		//} else if err == io.EOF {
		//	fmt.Println("done reading the file")
		//	break
		//} else {
		//	fmt.Println("Error: ", err)
		//}

		if err != nil {
			fmt.Println("Error: ", err)
			break
		}
		fmt.Println(string(buffer[:n]))

	}
	file.Close()
}
