package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var totalLines int
	fmt.Fscan(in, &totalLines)

	rulesPos := map[string]int{"L": -1, "R": 1, "B": -51, "E": 51}
	rulesLine := map[string]int{"D": 1, "U": -1}
	rulesLinePos := map[string]int{"N": 1}

	for i := 0; i < totalLines; i++ {
		var lineNumber int
		final := []string{""}
		var line string
		var pos int
		fmt.Fscan(in, &line)
		for _, lValue := range line {
			l := string(lValue)
			if rule, ok := rulesPos[l]; ok {
				pos += rule
				if len(final[lineNumber]) <= pos {
					pos = len(final[lineNumber])
				}
				if pos < 0 {
					pos = 0
				}
			} else if rule, ok := rulesLine[l]; ok {
				if want := lineNumber + rule; 0 <= want && len(final) > want {
					lineNumber = want
					if len(final[lineNumber]) <= pos {
						pos = len(final[lineNumber])
					}
				} else {
					fmt.Print()
				}
			} else if _, ok := rulesLinePos[l]; ok {
				final = append(final, final[lineNumber][pos:len(final[lineNumber])])
				final[lineNumber] = fmt.Sprintf("%s", final[lineNumber][0:pos])
				lineNumber++
				if len(final)-1 > lineNumber {
					s := final[len(final)-1]
					final[len(final)-1] = final[lineNumber]
					final[lineNumber] = s
				}
				pos = 0
			} else {
				final[lineNumber] = fmt.Sprintf("%s%s%s", final[lineNumber][0:pos], l, final[lineNumber][pos:len(final[lineNumber])])
				pos++
			}
		}
		for _, l := range final {
			fmt.Fprintln(out, l)
		}
		fmt.Fprintln(out, "-")
	}
}
