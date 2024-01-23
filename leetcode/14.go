package main

import "fmt"

func longestCommonPrefix(strs []string) string {

	length := len(strs[0])

	r := ""

	for _, v := range strs {
		if len(v) < length {
			length = len(v)
		}
	}

	for i := 0; i < length; i++ {
		isPrefixMatching := true
		for j := 1; j < len(strs); j++ {
			if strs[j][i] != strs[j-1][i] {
				isPrefixMatching = false
			}
		}
		if isPrefixMatching {
			r += string(strs[0][i])
		} else {
			break
		}
	}

	return r
}

func main() {
	fmt.Printf("%s\n", longestCommonPrefix([]string{"flower", "flow", "flight"}))
	fmt.Printf("%s\n", longestCommonPrefix([]string{"bigger", "banger", "barbie"}))

}
