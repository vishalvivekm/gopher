package main

import (
	"fmt"
	"log"
	"os/exec"
	"path/filepath"
	"strings"
)

func main() {

	out, err := exec.Command("git", "config", "--get", "remote.origin.url").CombinedOutput() // .Run()
	if err != nil {
		log.Fatalf("error getting origin url:  %s\n", err)
	}
	originUrl := string(out)
	// fmt.Printf("combined out:\n%s\n", string(out))
	fmt.Println(originUrl)
	//u, err := url.Parse(originUrl)
	//if err != nil {
	//	panic(err)
	//} \n": net/url: invalid control character in URL
	fmt.Println("u.Path: ", u)
	//repoName := filepath.Base(originUrl)
	repoName := filepath.Base(originUrl)
	//fmt.Println(repoName)
	fmt.Println(strings.TrimSuffix(repoName, ".git"))
	//if "meshery" == strings.TrimSuffix(repoName, ".git") {
	//	fmt.Println("repo is meshery")
	//
	//} else {
	//	log.Fatal("Please ensure this command is run from a fork of meshery")
	//}
}
