package main

import (
	"fmt"
)

type APIError struct {
	// your code goes here
}

func (e *APIError) Error() string {
	// your code goes here
}

func NewAPIError(code int, message string, details map[string]interface{}) error {
	// your code goes here
}

func main() {
	err := NewAPIError(400, "Bad request", map[string]interface{}{
		"field": "username",
		"error": "cannot be empty",
	})
	fmt.Println(err.Error())
}

/*
In this question, we are building a Go application that involves interacting with a remote API. We want to handle API errors in a consistent and structured way. We have chosen to use the errors package to accomplish this task.

Our task is to create a new error type called APIError that implements the error interface.
The APIError type should have the following fields:


Message: A string that describes the error.


Code: An integer that represents the error code returned by the API.


Details: A map[string]interface{} that contains additional error details.


We should implement the Error method for the APIError type, which should return the error message. Also, we should implement a function called NewAPIError that creates a new instance of the APIError type.


A Go file is located at /root/code/infra directory.

package main

import (
    "fmt"
)

type APIError struct {
    // your code goes here
}

func (e *APIError) Error() string {
    // your code goes here
}

func NewAPIError(code int, message string, details map[string]interface{}) error {
    // your code goes here
}

func main() {
    err := NewAPIError(400, "Bad request", map[string]interface{}{
        "field": "username",
        "error": "cannot be empty",
    })
    fmt.Println(err.Error())
}


Expected Output:

Bad request
*/

/*

// question:
In this question, we are building a Go application that involves interacting with a remote API. We want to handle API errors in a consistent and structured way. We have chosen to use the errors package to accomplish this task.


Our task is to create a new error type called APIError that implements the error interface.
The APIError type should have the following fields:


Message: A string that describes the error.


Code: An integer that represents the error code returned by the API.


Details: A map[string]interface{} that contains additional error details.


We should implement the Error method for the APIError type, which should return the error message. Also, we should implement a function called NewAPIError that creates a new instance of the APIError type.


A Go file is located at /root/code/infra directory.


package main

import (
    "fmt"
)

type APIError struct {
    // your code goes here
}

func (e *APIError) Error() string {
    // your code goes here
}

func NewAPIError(code int, message string, details map[string]interface{}) error {
    // your code goes here
}

func main() {
    err := NewAPIError(400, "Bad request", map[string]interface{}{
        "field": "username",
        "error": "cannot be empty",
    })
    fmt.Println(err.Error())
}


Expected Output:


Bad request

Answer:
package main

import (
  "fmt"
)

type APIError struct {
  // your code goes here
  code int
  message string
  details map[string]interface{}
}

func (e *APIError) Error() string {
  // your code goes here
  return fmt.Sprintf(e.message)
}

func NewAPIError(code int, message string, details map[string]interface{}) error {
  // your code goes here
  return &APIError{
    code,
    message,
    details,
  }
}

func main() {
  err := NewAPIError(400, "Bad request", map[string]interface{}{
     "field": "username",
     "error": "cannot be empty",
  })
  fmt.Println(err.Error())
}



aliter,


package main

import (
    "fmt"
)

type APIError struct {
    Message string
    Code    int
    Details map[string]interface{}
}

func (e *APIError) Error() string {
    return e.Message
}

func NewAPIError(code int, message string, details map[string]interface{}) error {
    return &APIError{
        Message: message,
        Code:    code,
        Details: details,
    }
}

func main() {
    err := NewAPIError(400, "Bad request", map[string]interface{}{
        "field": "username",
        "error": "cannot be empty",
    })
    fmt.Println(err)
}




*/

/*
We are building a Go application and we want to add logging to help you debug and troubleshoot issues. We have chosen to use the logrus package to accomplish this task.

Problem Statement:


The logrus package provides several functions for logging messages at different levels (e.g. Debug, Info, Warning, Error).
Our task is to use the logrus package to log messages at different levels. We should create a function called logMessage that takes in a level string and a message string as arguments.

If the level string is debug, the function should log the message using the Debug function.

If the level string is info, the function should log the message using the Info function.

If the level string is warning, the function should log the message using the Warning function.

If the level string is error, the function should log the message using the Error function.

If the level string is none of these, the function should log an error message using the Error function in the following format: -

"Invalid log level: " + level

A Go file is located at /root/code/logrus directory.

package main

import log "github.com/sirupsen/logrus"

// your code goes here

func main() {
	log.SetLevel(log.DebugLevel)
	logMessage("debug", "This is a debug message")
	logMessage("info", "This is an info message")
	logMessage("warning", "This is a warning message")
	logMessage("error", "This is an error message")
	logMessage("invalid", "This is an invalid message")
}


Expected Output:

DEBU[0000] This is a debug message
INFO[0000] This is an info message
WARN[0000] This is a warning message
ERRO[0000] This is an error message
ERRO[0000] Invalid log level: invalid

package main

import log "github.com/sirupsen/logrus"


// your code goes here

func main() {
	log.SetLevel(log.DebugLevel)
	logMessage("debug", "This is a debug message")
	logMessage("info", "This is an info message")
	logMessage("warning", "This is a warning message")
	logMessage("error", "This is an error message")
	logMessage("invalid", "This is an invalid message")
}
func logMessage(level string, message string) {
	switch level {
	case "debug":
		log.Debug(message)
	case "info":
		log.Info(message)
	case "warning":
		log.Warning(message)
	case "error":
		log.Error(message)
	default:
		log.Error("Invalid log level: " + level)
	}

}
*/

/*
Tests can't help you optimize your code for better performance.

Explanation:

Other options are all benefits of writing tests for your Go code. Test can help ensure that your code is correct and free of bugs, help you understand how your code is supposed to work and help you refactor your code with confidence. However, writing tests is not a direct way to optimize your code for better performance. That said, writing tests can indirectly lead to better performance, as it can help you find and fix performance issues that you might not have otherwise noticed.
*/
