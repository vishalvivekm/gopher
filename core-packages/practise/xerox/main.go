package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	// your code goes here
	file, err := os.Open("input.txt")
	if err != nil {
		fmt.Println(err)
		os.Exit(1) // or // return

	}
	var sum int
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {

		}
	}(file)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		num, err := strconv.Atoi(scanner.Text())
		if err != nil {
			fmt.Println(err)
			return
		}
		sum += num
	}
	//fmt.Println(sum)
	outputFile, err := os.OpenFile("output.txt", os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0600)
	if err != nil {
		log.Fatal(err)
	}
	defer func(outputFile *os.File) {
		err := outputFile.Close()
		if err != nil {

		}
	}(outputFile)
	_, err = outputFile.WriteString(strconv.Itoa(sum))
	if err != nil {
		return
	}
}

/*We have given a text file called input.txt that contains a list of integers,
one per line.
Our task is to read the given file, sum the integers,
and write the sum to a new file called output.txt within a same directory.

A file output.txt would be created containing the sum as follows: -

160
*/
/*
soln:
package main

import (
   "bufio"
   "fmt"
   "os"
   "strconv"
)

func main() {
   // Open the input file
   inputFile, err := os.Open("input.txt")
   if err != nil {
      fmt.Println(err)
      return
   }
   defer inputFile.Close()

   // Create a scanner to read the input file
   inputScanner := bufio.NewScanner(inputFile)

   // Initialize the sum to 0
   sum := 0

   // Iterate over the lines of the input file
   for inputScanner.Scan() {
      // Parse the current line as an integer
      n, err := strconv.Atoi(inputScanner.Text())
      if err != nil {
         fmt.Println(err)
         return
      }

      // Add the integer to the sum
      sum += n
   }

   // Check for errors while reading the input file
   if err := inputScanner.Err(); err != nil {
      fmt.Println(err)
      return
   }

   // Open the output file
   outputFile, err := os.Create("output.txt")
   if err != nil {
      fmt.Println(err)
      return
   }
   defer outputFile.Close()

   // Write the sum to the output file as a string
   _, err = outputFile.WriteString(strconv.Itoa(sum))
   if err != nil {
      fmt.Println(err)
      return
   }
}


Then, run the following go run command as follows: -

go run main.go


It will sum the integers, and write the sum 160 to a new file called output.txt.

# List the files of the /root/code/xerox directory
ls -l

# Verify the sum

cat output.txt
*/
