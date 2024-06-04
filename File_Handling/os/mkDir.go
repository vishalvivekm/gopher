package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	//err := os.Mkdir("./vcds", os.ModePerm)
	//if err != nil {
	//	log.Fatalln(err)
	//}
	var err error
	createEmptyDir := func(name string) error {
		d := []byte("")
		if err := os.WriteFile(name, d, 0666); err != nil {
			return err
		}
		return nil
	}
	err = createEmptyDir("./vcds/index.txt")
	if err != nil {
		log.Fatalln(err)
	}
	//defer func() {
	//	err := os.RemoveAll("vcds")
	//	if err != nil {
	//		log.Fatalln(err)
	//	}
	//}()
	c, err := os.ReadDir("./")
	if err != nil {
		log.Fatalln(err)
	}
	for _, entry := range c {
		fmt.Println(entry.Name(), entry.IsDir())
	}

}
