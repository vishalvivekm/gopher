package main

import (
	"log"
	"os"
)

func main() {
	// creates a new file, or truncates the already existing file
	file, err := os.Create("new_file.txt")
	if err != nil {
		log.Fatal(err)
	}
	// os.OpenFile() - doesn't truncate the file in case it already exists
	// creates a new file using the os.O_CREATE flag and the 0666 permission
	file, err := os.OpenFile("text.txt", os.O_CREATE, 0666)
	if err != nil {
		log.Fatal(err)
	}

}

// to delete file:
//os.Remove() - a single file. Also: will error if directory isn't empty,
//os.RemoveAll() - all the folders and file inside the path

// to create dir
// os.Mkdir  // throws error if the directory already exists
// os.MkdirAll - directories and subdirectories, doesn't throw error if dir already exist
