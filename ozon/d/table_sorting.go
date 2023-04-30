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
	fmt.Fscan(in, struct{}{})
	for i := 0; i < totalLines; i++ {
		var n int
		var m int
		fmt.Fscan(in, &n, &m)
		table := make([][]int, n, n)
		for j := 0; j < n; j++ {
			table[j] = make([]int, m, m)
			for k := 0; k < m; k++ {
				var a int
				fmt.Fscan(in, &a)
				table[j][k] = a
			}
		}
		var clicks int
		fmt.Fscan(in, &clicks)
		for clickNumber := 0; clickNumber < clicks; clickNumber++ {
			var column int
			fmt.Fscan(in, &column)
			for line2 := 0; line2 < n-1; line2++ {
				for line := 0; line < n-line2-1; line++ {
					if table[line][column-1] > table[line+1][column-1] {
						c := table[line]
						table[line] = table[line+1]
						table[line+1] = c
					}
				}
			}
			//sort.Slice(table, func(i, j int) bool {
			//	return table[i][column-1] < table[j][column-1]
			//})
			//fmt.Println(column)
			//for j := 0; j < n; j++ {
			//	for k := 0; k < m; k++ {
			//		fmt.Printf("%d ", table[j][k])
			//	}
			//	fmt.Println()
			//}
		}
		fmt.Println()
		for j := 0; j < n; j++ {
			for k := 0; k < m; k++ {
				fmt.Printf("%d ", table[j][k])
			}
			fmt.Println()
		}
	}
}
