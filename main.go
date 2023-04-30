package main

import "fmt"

func main() {
	//res := make(chan int, 5)
	//fmt.Println(len(res))
	//fmt.Println(cap(res))
	//close(res)
	//
	//go func() { res <- 123 }()
	//res2 := make([]int, 0, 3)
	//fmt.Println(len(res2))
	//fmt.Println(cap(res2))
	//fmt.Println(res2)
	//
	//time.Sleep(1 * time.Second)

	//ar := []int{}
	//fmt.Println(len(ar))
	//fmt.Println(cap(ar))
	//ar = append(ar, 1)
	//fmt.Println(len(ar))
	//fmt.Println(cap(ar))
	//ar = append(ar, 1)
	//fmt.Println(len(ar))
	//fmt.Println(cap(ar))
	//ar = append(ar, 1)
	//fmt.Println(len(ar))
	//fmt.Println(cap(ar))
	//fmt.Println(ar)
	//
	//ar2 := [1]int{1}
	//fmt.Println(len(ar2))
	//fmt.Println(cap(ar2))

	a := make([]int, 2, 4)
	fmt.Printf("a old %p\n", a)

	a = append(a, 1)
	fmt.Printf("a new %p\n", a)
	fmt.Println(a)
}
