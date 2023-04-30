package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	//test := []string{
	//	"cat",
	//	"tax",
	//	"city",
	//	"roof",
	//	"chief",
	//	"bus",
	//	"blitz",
	//	"marsh",
	//	"house",
	//	"puppy",
	//	"counterpunch",
	//	"array",
	//}

	sw := map[string]struct{}{"b": {}, "c": {}, "d": {}, "f": {}, "g": {}, "h": {}, "j": {}, "k": {}, "l": {}, "m": {}, "n": {}, "p": {}, "q": {}, "r": {}, "s": {}, "t": {}, "v": {}, "w": {}, "x": {}, "z": {}}

	rules := map[string]struct{}{
		"s":  {},
		"x":  {},
		"z":  {},
		"sh": {},
		"ch": {},
	}

	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var totalLines int
	fmt.Fscan(in, &totalLines)
	for i := 0; i < totalLines; i++ {
		//for i := 0; i < len(test); i++ {
		var word string
		fmt.Fscan(in, &word)
		//word = test[i]
		lastLet := string(word[len(word)-1])

		if lastLet != "y" {
			if _, ok := rules[lastLet]; ok {
				word = fmt.Sprintf("%s%s", word, "es")
				//fmt.Println(lastLet)
				//fmt.Println(word)
				fmt.Fprintln(out, word)
				continue
			}
		}

		if lastLet == "y" {
			if _, ok := sw[string(word[len(word)-2])]; ok {
				word = fmt.Sprintf("%s%s", word[0:len(word)-1], "ies")
				//fmt.Println(lastLet)
				//fmt.Println(word)
				fmt.Fprintln(out, word)
				continue
			}
		}

		lastTwoLet := fmt.Sprintf("%s%s", string(word[len(word)-2]), string(word[len(word)-1]))
		if _, ok := rules[lastTwoLet]; ok {
			word = fmt.Sprintf("%s%s", word, "es")
			//fmt.Println(lastLet)
			//fmt.Println(word)
			fmt.Fprintln(out, word)
			continue
		}
		word = fmt.Sprintf("%s%s", word, "s")
		//fmt.Println(lastLet)
		//fmt.Println(word)
		fmt.Fprintln(out, word)
	}
}
