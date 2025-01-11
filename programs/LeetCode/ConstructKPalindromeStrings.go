/*
LC 1400
Given a string s and an integer k, return true if you can use all the characters in s to construct k palindrome strings or false otherwise.

 

Example 1:

Input: s = "annabelle", k = 2
Output: true
Explanation: You can construct two palindromes using all characters in s.
Some possible constructions "anna" + "elble", "anbna" + "elle", "anellena" + "b"

Example 2:

Input: s = "leetcode", k = 3
Output: false
Explanation: It is impossible to construct 3 palindromes using all the characters of s.

Example 3:

Input: s = "true", k = 4
Output: true
Explanation: The only possible solution is to put each character in a separate string.


*/

func canConstruct(s string, k int) bool {
    if k > len(s) {
        return false
    }

  // only beat 50%
    // mp := make(map[string]int)
    // for _, ch := range s {
    //     mp[string(ch)]++
    // }
    // oddCount := 0
    // for _, val := range mp {
    //     if val % 2 != 0 {
    //         oddCount++
    //     }
    // }

  // this beat 100%
    count := [26]int{}
    for i := 0; i < len(s); i++ {
        count[s[i]-'a']++
    }
     oddCount := 0
    for i := 0; i < 26; i++ {
        oddCount += count[i] & 1
    }

    return oddCount <= k   
}
