package main

import (
	"fmt"
	"sort"
)

/*
dir1
--dir2
----dir3
--------temp.txt
*/
func main() {

	//path := filepath.Join("dir2", "../dir1/dir2//", "dir3", "../dir3/temp.txt")
	//path := filepath.Join("dir2", "../dir1/dir2", "dir3", "temp.txt")
	//path := filepath.Join("dir1", "../dir1/dir2", "dir3", "temp.txt")
	//fmt.Println(path)

	// read the temp.txt and give output like this:
	//Lorem Ipsum is simply dummy text of the printing a
	//nd typesetting industry. Lorem Ipsum has been the
	//industry's standard dummy text ever since the 1500
	//s, when an unknown printer took a galley of type a
	//nd scrambled it to make a type specimen book.

	//file, err := os.Open("temp.txt")
	//if err != nil {
	//	log.Println(err)
	//}
	//defer file.Close() // close the file before exiting the program
	//buf := make([]byte, 50)
	//for {
	//	n, err := file.Read(buf)
	//	fmt.Println(string(buf[:n]))
	//	if err != nil {
	//		break
	//	}
	//}

	var strs = []string{"Apple", "Around", "Armor", "An"}
	sort.Strings(strs)
	fmt.Println(strs)
	// log.Fatal: it prints the log and then safely exits the program. To know more about this, refer this page: - https://pkg.go.dev/log#Fatal
}
