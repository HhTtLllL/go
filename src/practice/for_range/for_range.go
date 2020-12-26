package main

import (
	"fmt"
)

func main() {

/*
	numbers1 := []int{1, 2, 3, 4, 5, 6}

	for _, i := range numbers1 {

		fmt.Println(i)

		if i == 3 {

			numbers1[i] |= i
		}
	}

	fmt.Println("number", numbers1)
*/
/*
	numbers := [...]int{1, 2, 3, 4, 5, 6}
	maxIndex2 := len(numbers) - 1

	for i, e := range numbers {

		fmt.Printf("%d\n", e)
		if i == maxIndex2 {

			numbers[0] += e
		}else {

			numbers[i + 1] += e
		}
	}

	*/
/*

	for i := 0; i < len(numbers); i ++ {

		fmt.Printf("%d\n", numbers[i])
		if i == maxIndex2 {

			numbers[0] += numbers[i]
		}else {

			numbers[i + 1] += numbers[i]
		}

	}
*/

	numbers3 := []int{1, 2, 3, 4, 5, 6}
	maxIndex3 := len(numbers3) - 1

		for i, e := range numbers3 {

			fmt.Printf("%d\n", e)
			if i == maxIndex3 {

				numbers3[0] += e
			}else {

				numbers3[i + 1] += e
			}
		}


	fmt.Println(numbers3)
}