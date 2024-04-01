// read from data.txt and print the lines as it

package main

import (
	"bufio"
	"fmt"
	"os"
)

func ReadLines(filename string) ([]string, error) {

	// your code goes here
	str := make([]string, 0, 1024)
	file, err := os.Open("data.txt")
	if err != nil {
		return []string{}, fmt.Errorf("Can't open file, ", err)
		//log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		str = append(str, line)

	}
	return str, nil
}

func main() {
	lines, err := ReadLines("data.txt")
	if err != nil {
		fmt.Println(err)
	} else {
		for _, line := range lines {
			fmt.Println(line)
		}
	}

}

/*
package main

import (
   "bufio"
   "fmt"
   "os"
)

func ReadLines(filename string) ([]string, error) {
   // your code goes here
   file, err := os.Open(filename)
   if err != nil {
      return nil, err
   }
   defer file.Close()

   scanner := bufio.NewScanner(file)
   var lines []string
   for scanner.Scan() {
      lines = append(lines, scanner.Text())
   }
   if err := scanner.Err(); err != nil {
      return nil, err
   }

   return lines, nil
}

func main() {
   lines, err := ReadLines("data.txt")
   if err != nil {
      fmt.Println(err)
   } else {
      for _, line := range lines {
         fmt.Println(line)
      }
   }

}


Explanation:

This function uses the os package to open the text file specified by filename, and then it uses the bufio package to read the lines from the file using a bufio.Scanner.


The function reads each line from the file using the Scan method and appends it to a slice of strings.
At the end of the loop, the function returns the slice of strings and any error that occurred while reading the file.
*/

/*problem:

We are working on a project that involves reading data from a text file and processing it. We need to write a function that reads the contents of a text file and returns a slice of strings, where each string represents a line in the file.


A text file called data.txt contains the below content inside the /root/code/data directory.

Lorem Ipsum is simply dummy text of the printing and typesetting industry.
Lorem Ipsum has been the industry's standard dummy text ever since the 1500s.
when an unknown printer took a galley of type and scrambled it to make a type specimen book.



Also, a Go file is located inside the directory called /root/code/data/.

package main

import (
  "bufio"
  "fmt"
  "os"
)

func ReadLines(filename string) ([]string, error) {
    // your code goes here
}

func main() {
  lines, err := ReadLines("data.txt")
  if err != nil {
     fmt.Println(err)
  } else {
     for _, line := range lines {
        fmt.Println(line)
     }
  }

}


Expected Output:

Lorem Ipsum is simply dummy text of the printing and typesetting industry.
Lorem Ipsum has been the industry's standard dummy text ever since the 1500s.
when an unknown printer took a galley of type and scrambled it to make a type specimen book.
*/
