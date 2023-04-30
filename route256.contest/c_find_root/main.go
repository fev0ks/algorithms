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

	for i := 0; i < totalLines; i++ {
		tree := make(map[int]int)
		var totalNumbers int
		var prevChild int
		subTreeLen := -1
		fmt.Fscan(in, &totalNumbers)
		for j := 0; j < totalNumbers; j++ {
			var value int
			if subTreeLen > 0 {
				var child int
				fmt.Fscan(in, &child)
				tree[child] = prevChild
				subTreeLen--
				if subTreeLen == 0 {
					prevChild = 0
				}
				continue
			}

			fmt.Fscan(in, &value)
			if prevChild != 0 {
				subTreeLen = value
				if subTreeLen == 0 {
					prevChild = 0
					subTreeLen = -1
				}
			} else {
				prevChild = value
				if _, ok := tree[prevChild]; !ok {
					tree[prevChild] = 0
				}
			}
		}

		for k, v := range tree {
			if v == 0 {
				fmt.Fprintln(out, k)
				break
			}
		}
	}

}
