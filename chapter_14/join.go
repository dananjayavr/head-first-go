package main

import (
	"fmt"
	"strings"
)

func JoinWithCommas(phrases []string) string {
	if len(phrases) == 0{
		return ""
	} else if len(phrases) == 1 {
		return phrases[0]
	} else if len(phrases) == 2 {
		return phrases[0] + " and " + phrases[1]
	} else {
		result := strings.Join(phrases[:len(phrases)-1], ", ")
		result += ", and "
		result += phrases[len(phrases)-1]
		return result
	}
}

func main() {
	/*	testStr := make([]string,4)
		testStr[0] = "a"
		testStr[1] = "b"
		testStr[2] = "c"
		testStr[3] = "d"

		fmt.Println(testStr[:len(testStr)-1])*/

	/*	fmt.Println(strings.Join([]string{"05","14","2018"},"/"))
		fmt.Println(strings.Join([]string{"state","of","the","art"}, "-"))*/

	phrases := []string{"my parents", "a rodeo clown"}
	fmt.Println("A photo of", JoinWithCommas(phrases))
	phrases = []string{"my parents", "a rodeo clown", "a prize bull"}
	fmt.Println("A photo of", JoinWithCommas(phrases))
	phrases = []string{"my parents"}
	fmt.Println("A photo of", JoinWithCommas(phrases))
}
