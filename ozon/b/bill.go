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
		var totalNumbers int
		sum := 0
		fmt.Fscan(in, &totalNumbers)
		staff := make(map[int]int)
		for j := 0; j < totalNumbers; j++ {
			var a int
			fmt.Fscan(in, &a)
			if n, ok := staff[a]; ok {
				if n == 2 {
					staff[a] = 0
					continue
				}
			}
			staff[a] += 1
			sum += a
		}
		fmt.Fprintln(out, sum)
	}

}
