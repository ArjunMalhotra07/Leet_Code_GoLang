//https://leetcode.com/problems/maximum-number-of-words-found-in-sentences/
package main

import (
	"fmt"
	"strings"
)

func main() {

	fmt.Println("hlo")
	sentences := []string{"alice and bob love leetcode", "i think so too", "this is great thanks very much"}
	ans := mostWordsFound(sentences)

	fmt.Println(ans)

}
func mostWordsFound(sentences []string) int {

	length := len(sentences)
	var max int = 0

	for i := 0; i < length; i++ {
		str := sentences[i]

		words := strings.Fields(str)
		x := len(words)
		if x > max {
			max = x
		}

	}
	return max
}