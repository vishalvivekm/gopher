// LC28, runtime beats 100%
func strStr(haystack string, needle string) int {
    k := len(needle)

    i := -1
    start := 0
    sliceStr := strings.Split(haystack, "")
    for start + k <= len(haystack){
        if strings.Join(sliceStr[start:start+k], "") == needle{
            if i == -1 {
                i = start
                break;
            }
        }
        start++
    }
    return i
}
