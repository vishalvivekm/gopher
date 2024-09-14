// https://leetcode.com/contest/weekly-contest-414/problems/convert-date-to-binary/

import (
"strconv"
"strings"
)
func convertDateToBinary(date string) string {
	slice := strings.Split(date, "-")
	yyStr, _ := strconv.Atoi(slice[0])
	mmStr, _ := strconv.Atoi(slice[1])
	ddStr, _ := strconv.Atoi(slice[2])
	yy := strconv.FormatInt(int64(yyStr), 2)
	mm := strconv.FormatInt(int64(mmStr), 2)
	dd := strconv.FormatInt(int64(ddStr), 2)
	ans := strings.Join([]string{yy, mm, dd}, "-")
    return ans
}
