package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strings"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var totalLines int
	fmt.Fscan(in, &totalLines)

	for i := 0; i < totalLines; i++ {
		var totalNumbers int
		fmt.Fscan(in, &totalNumbers)
		lenInterval := 0
		prevValue := -1002
		startInterval := -1002
		var compressed []string
		for j := 0; j < totalNumbers; j++ {
			var value int
			fmt.Fscan(in, &value)
			if prevValue > -1002 {
				if signed := value - prevValue; math.Abs(float64(signed)) == 1 && (lenInterval == 0 || math.Signbit(float64(signed)) == math.Signbit(float64(lenInterval))) {
					lenInterval = lenInterval + signed
					prevValue = value
				} else {
					compressed = append(compressed, fmt.Sprintf("%d", startInterval))
					compressed = append(compressed, fmt.Sprintf("%d", lenInterval))
					startInterval = value
					prevValue = value
					lenInterval = 0
				}
			} else {
				startInterval = value
				prevValue = value
			}
		}
		compressed = append(compressed, fmt.Sprintf("%d", startInterval))
		compressed = append(compressed, fmt.Sprintf("%d", lenInterval))
		fmt.Fprintln(out, len(compressed))
		fmt.Fprintln(out, strings.Join(compressed, " "))
	}
}
