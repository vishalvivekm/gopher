package main

import (
"fmt"
"log"
"os/exec"
)

func main() {

// err := exec.Command("git config --get remote.origin.url | sed 's/.*\/\([^ ]*\/[^.]*\).*/\1/'").Run()
out, err := exec.Command("git",  "config", "--get",  "remote.origin.url").CombinedOutput() // .Run()
        if err != nil { // out, err := cmd 
        log.Fatalf("error getting origin url:  %s\n", err)
    }
    fmt.Printf("combined out:\n%s\n", string(out))
/*if err != nil {
fmt.Errorf("%v", err)
*/

//args := []string{"-y", "-i", "movie.mp4", "movie_audio.mp3", "INVALID-ARG!"}

}

