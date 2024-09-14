// https://leetcode.com/problems/count-the-number-of-consistent-strings/

func countConsistentStrings(allowed string, words []string) int {
    count := 0
    
    mp := make(map[int32]struct{})
    for _, ch := range allowed {
        fmt.Println(ch)
        mp[ch] = struct{}{}
    }
    fmt.Println("%+v", mp)
    for _, str := range words {
        for index, c := range str {
            _, ok := mp[c];
            if !ok {
                break;
            }
            if index == (len(str) - 1) {
                count++
            }
        }
    }

    return count
}
