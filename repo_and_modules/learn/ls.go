package main

import (
    "bytes"
    "fmt"
    "log"
    "os/exec"
)

const ShellToUse = "bash"

func Shellout(command string) (string, string, error) {
    var stdout bytes.Buffer
    var stderr bytes.Buffer
    cmd := exec.Command(ShellToUse, "-c", command)
    cmd.Stdout = &stdout
    cmd.Stderr = &stderr
    err := cmd.Run()
    return stdout.String(), stderr.String(), err
}

func main() {
    out, errout, err := Shellout("cat ./lock.txt")
    if err != nil {
        log.Printf("error: %v\n", err)
    }
    fmt.Println("--- stdout ---")
    fmt.Println(out)
    fmt.Println("--- stderr ---")
    fmt.Println(errout)
}
