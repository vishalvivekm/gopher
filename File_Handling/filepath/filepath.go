package main

import (
	"fmt"
	"path/filepath"
)

func main() {

	path := filepath.Join("dir1", "dir2/", "text.txt")
	//path := filepath.Join("dir1", "dir2//", "text.txt")
	// path := filepath.Join("dir1", "dir2//", "dir3/../dir4", "text.txt") // dir1/dir2/dir4/text.txt
	//fmt.Println(path)

	fmt.Println(filepath.Dir(path))  // dir1/dir2 : dir of file
	fmt.Println(filepath.Base(path)) // text.txt: name of the file

	//fmt.Println(filepath.IsAbs(path)) // false
	fmt.Println(filepath.IsAbs("/dir/file")) // true // absolute path has / in beginning

	fmt.Println(filepath.Ext(path)) // .txt
}
