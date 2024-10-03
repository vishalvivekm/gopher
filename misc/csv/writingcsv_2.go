package main

import (
    "encoding/csv"
    "io"
    "log"
    "os"
)

func main() {
    filename := "new_friends.csv"
    file, err := os.Create(filename)
    if err != nil {
        log.Fatalln("Cannot create new_friends.csv:", err)
    }
    defer file.Close()

    data := [][]string{
        {"1", "John", "Doe", "john@google.in"},
        {"2", "Jane", "Doe", "jane@google.in"},
        {"3", "Michael", "Smith", "michael@google.in"},
    }

    writer := csv.NewWriter(file)
    for _, row := range data {
        err := writer.Write(row)
        if err != nil {
            log.Fatalln("Cannot write to csv file:", err)
        }
    }
    writer.Flush()

    if err := writer.Error(); err != nil {
        log.Fatalln("Error flushing writer:", err)
    }

    // reset cursor to the beginning of the file
    _, err = file.Seek(0, 0)
    if err != nil {
        log.Fatalln("can't seek to beginning of file:", err)
    }

    content, err := io.ReadAll(file)
    if err != nil {
        log.Fatalln("can not read file:", err)
    }
    log.Println("File contents:")
    log.Println(string(content))

    log.Println("INFO: Data successfully written to and read from csv", filename)
}