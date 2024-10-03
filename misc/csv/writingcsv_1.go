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

    // Close the file to ensure all data is written
    file.Close()

    // Reopen the file for reading
    file, err = os.Open(filename)
    if err != nil {
        log.Fatalln("Cannot open file for reading:", err)
    }
    defer file.Close()

    content, err := io.ReadAll(file)
    if err != nil {
        log.Fatalln("Cannot read file:", err)
    }
    log.Println("File contents:")
    log.Println(string(content))

    log.Println("INFO: Data successfully written to and read from csv", filename)
}