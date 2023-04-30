package main

import (
	"fmt"
)

// на вход два канала с неубывающей последовательностью, вывести 1 канал с неубывающей последовательностью из них
func main() {
	array1 := []int{0, 4, 5, 7, 9}
	array2 := []int{0, 1, 2, 3, 4, 7, 13, 17, 19}

	ch1 := make(chan int, len(array1))
	ch2 := make(chan int, len(array2))

	ch3 := make(chan int)

	go func() {
		for _, number := range array1 {
			ch1 <- number
		}
		close(ch1)
	}()

	go func() {
		for _, number := range array2 {
			ch2 <- number
		}
		close(ch2)
	}()

	go func() {
		close1 := false
		close2 := false

		hold1 := -1
		hold2 := -1

		ok1 := true
		ok2 := true
		for {
			first := -1
			second := -1

			if !close1 {
				if hold1 == -1 {
					first, ok1 = <-ch1
					if !ok1 {
						close1 = true
						first = -1
					}
				} else {
					first = hold1
					hold1 = -1
				}
			}

			if !close2 {
				if hold2 == -1 {
					second, ok2 = <-ch2
					if !ok2 {
						close2 = true
						second = -1
					}
				} else {
					second = hold2
					hold2 = -1
				}
			}

			if !close1 && first < second || close2 && !close1 {
				ch3 <- first
				hold2 = second
			} else if !close2 && second <= first || close1 && !close2 {
				ch3 <- second
				hold1 = first
			}

			if close1 && close2 {
				close(ch3)
				return
			}

		}
	}()

	for number := range ch3 {
		fmt.Println(number)
	}
}
