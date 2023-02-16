package main

import "fmt"

/*

   matrix processing
   Matrix A = 3 * 4 i.e 3 rows and 4 columns
    A =  [
		    j	j j  j
			0	1 2  3
		i=0	1  2  3  4
		i=1	5  6  7  8
		i=2	9  4  5  1
	      ]

     To store this data in memory we need to use [][] 2 dimensional array

	 Rows are procssed in outer loop and columns are processed in inner loop




*/

func outer(x int) func(y int) int {
	z := func(u int) int {
		return x * x * u
	}
	return z
}

func square(x int) int {
	return x * x
}

func cube(x int) int {
	return x * x * x
}

func applyfunc(f func(x int) int, y []int) []int {
	z := make([]int, len(y))
	for i, v := range y {
		z[i] = f(v)

	}
	return z
}
func greeting(myChannel chan string) {
	myChannel <- "hi"
}

func main() {
	table := [3][4]int{{1, 2, 3, 4}, {5, 6, 7, 8}, {2, 4, 6, 8}}
	fmt.Println("::::::", table[1][2])

	for x := range table {
		for y := range table[x] {
			fmt.Print(table[x][y], " ")
		}
		fmt.Println()
	}
	myChannel := make(chan string)
	go greeting(myChannel)
	fmt.Println(<-myChannel)
	matrix := [2][3]int{{1, 2, 3}, {4, 5, 6}}

	// Iterate over the matrix
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			fmt.Print(matrix[i][j], " ")
		}
		fmt.Println()
	}
	fmt.Println("Hello world", matrix[0][1])

	l := make([]int, 5, 10)
	l = append(l, 1, 2, 3, 4, 5)

	s := applyfunc(square, l)
	c := applyfunc(cube, l)

	j := [4]int{1, 2, 3, 4}
	for i := 0; i < len(j); i++ {
		fmt.Println("::", j[i])
	}

	for i, x := range j {
		fmt.Println(":::", j[i], x)
	}

	fmt.Println(s)
	fmt.Println(c)

}
