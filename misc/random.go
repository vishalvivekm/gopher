package main

import (
	"fmt"
	//"regexp"
	"strings"
)

func main() {
	/*htmlContent := `<a href="https://example.com">Example Link</a>
	                         <a href="https://another.example">Another Link</a>
	<a href="https://github.com/vishal,.vivekm">`

		re := regexp.MustCompile(`<a.*?href=\"(.*?)\"`)
		matches := re.FindAllStringSubmatch(htmlContent, -1)
		for _, match := range matches {
			fmt.Println(match[1])
		}*/
	t := `Hey there
how are you all doing
I'm fine`
	s := []int{1, 2, 3, 4, 5}
	for _, val := range s {
		if val == 3 {
			continue
		}
		fmt.Println(val)
	}
	for _, line := range strings.Split(string(t), "\n") {
		if line == "how are you all doing" {
			continue
		}
		fmt.Println(line)

	}
}
