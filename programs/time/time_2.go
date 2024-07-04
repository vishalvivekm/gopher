/*
// You can edit this code!
// Click here and start typing.
package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Hello, 世界")
	tie, err := ParseDateTimeWithFormat("")
	fmt.Println(tie, err)
}
func ParseDateTimeWithFormat(dateTimeStr string) (time.Time, error) {
	layouts := []string{
		time.RFC3339,                   // Layout for strings without the offset
		"2006-01-02T15:04:05.000-0700", // Layout for strings with the offset
	}

	var parsedTime time.Time
	var err error

	for _, layout := range layouts {
		parsedTime, err = time.Parse(layout, dateTimeStr)
		if err == nil {
			return parsedTime, nil
		}
	}

	return time.Time{}, fmt.Errorf("unable to parse date-time string: %s", dateTimeStr)
}

*/
