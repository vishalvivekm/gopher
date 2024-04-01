package main

import (
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

func main() {

	out, err := exec.Command("git", "config", "--get", "remote.origin.url").CombinedOutput()
	if err != nil {
		log.Fatalf("error getting origin url:  %s\n", err)
	}
	originUrl := string(out)
	repositoryName := strings.Split(filepath.Base(originUrl), ".")[0]
	//fmt.Println("originURL: ", originUrl)
	//fmt.Println("repo: ", repositoryName)

	if repositoryName != "gopher" {
		log.Fatal("Please ensure the command is executed from a fork of meshery/meshery")
	}
	pwd, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	if filepath.Base(pwd) != "gopher" {
		log.Fatal("Please run this command from the root of the project")
	}
	println(pwd)
}
