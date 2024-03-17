package main

import (
	"fmt"
	"os"
)

func main() {
	fileInfo, err := os.Stat("./temp.txt")
	if err != nil {
		fmt.Println(err)
	}
	//fmt.Println(fileInfo)
	fmt.Println(fileInfo.Name())
	fmt.Println(fileInfo.Size())
	fmt.Println(fileInfo.Mode())
	fmt.Println(fileInfo.IsDir())

}

/*
temp.txt
24
-rw-rw-r--
false

*/
