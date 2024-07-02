package main

import (
	"fmt"
	"strconv"
	"time"
)

func main() {
	i, err := strconv.ParseInt("1405544146", 10, 64)
	if err != nil {
		panic(err)
	}
	tm := time.Unix(i, 0)
	fmt.Println(tm)

	var timestamp int64 = 1719373985730
	tmt := time.Unix(timestamp/1000, 0)
	fmt.Println(tmt)

	//t := "2024-06-10T05:54:33Z"
	//t := "2024-06-10T09:46:32.346+0530"
	//t := "2024-06-10T04:21:11Z"
	/* ts, err := time.Parse("2006-01-02T15:04:05.999Z07:00", t)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(ts) */

	timeV, err1 := ParseDateTimeWithFormat("2024-06-10T09:46:32.346+0530")
	//      [7/2, 6:20 PM] dex: // 2024-06-10T09:46:32.346+0530

	if err1 != nil {
		fmt.Println(err1)
	}
	fmt.Println(timeV)

}
func ParseDateTimeWithFormat(dateTimeStr string) (time.Time, error) {
	layout := "2006-01-02T15:04:05.000-0700"
	return time.Parse(layout, dateTimeStr)
}
