package main

import (
	"bufio"
	"fmt"
	"math"
	//"math"
	"os"
	"sort"
)

type Pair struct {
	level int
	indx  int
}

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var totalLines int
	fmt.Fscan(in, &totalLines)
	for i := 0; i < totalLines; i++ {
		var totalWorkers int
		fmt.Fscan(in, &totalWorkers)
		workers := make([][]float64, totalWorkers, totalWorkers)
		fmt.Printf("%+v\n", workers)
		for j := 0; j < totalWorkers; j++ {
			workers[j] = make([]float64, totalWorkers, totalWorkers)
			var a float64
			fmt.Fscan(in, &a)
			workers[j][j] = a
		}
		for j := 0; j < totalWorkers; j++ {
			for k := j + 1; k < totalWorkers; k++ {
				workers[j][k] = math.Abs(workers[j][j] - workers[k][k])
			}
		}
		for k := 0; k < totalWorkers; k++ {
			workers[k][k] = -1
			sort.Slice(workers[k], func(i, j int) bool {
				return workers[k][i] < workers[k][j] || workers[k][i] == workers[k][j] && i < j
			})
		}
		//sort.Slice(workers, func(i, j int) bool {
		//	return workers[i].level < workers[j].level || workers[i].level == workers[j].level && workers[i].indx == workers[j].indx
		//})
		fmt.Printf("%+v\n", workers)
		//	if workerLevel[j] == 0 {
		//		continue
		//	}
		//	currentLevel := workerLevel[j]
		//	minDiff := 101.0
		//	indxCoworker := -1
		//	for k := j + 1; k < totalWorkers; k++ {
		//		if workerLevel[k] == -1 {
		//			continue
		//		}
		//		if diff := math.Abs(float64(currentLevel - workerLevel[k])); diff < minDiff {
		//			indxCoworker = k
		//			workerLevel[k] = -1
		//		}
		//	}
		//	fmt.Fprintln(out, j+1, indxCoworker+1)
		//}
	}
}

//1
//6
//2 1 3 1 1 4
