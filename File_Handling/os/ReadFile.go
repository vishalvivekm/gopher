package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"unicode/utf8"
)

func main() {
	path := "../temp.txt"
	data, err := os.ReadFile(path)
	if err != nil {
		fmt.Println("cannot read the file ", err)
	}
	fmt.Println(string(data))
	//os.Stdout.Write(data)
	fmt.Println(utf8.RuneCountInString(string(data)))

	// to read in chunk:
	buf := make([]byte, 5)
	file, err := os.Open("./temp.txt")
	defer file.Close()
	if err != nil {
		log.Fatal("error")
	}

	for {
		n, err := file.Read(buf)
		switch err {
		case nil:
			fmt.Println(string(buf[:n]))
		case io.EOF:
			fmt.Println("read the whole file: ", string(buf[:]))
			return
		default:
			fmt.Println("error occured", err)
		}
		//fmt.Println(string(buf[:n]))
	}

}
