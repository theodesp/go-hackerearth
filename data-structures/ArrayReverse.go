package data_structures

import (
	"fmt"
)

/*
Given the size and the elements of array A, print all the elements in reverse order.

Input:
First line of input contains, N - size of the array.
Following N lines, each contains one integer, i{th} element of the array i.e. A[i].

Output:
Print all the elements of the array in reverse order, each element in a new line.
 */

func PrintArrayReverse()  {
	var N int
	fmt.Scanf("%d", &N)
	var arr []int

	for i:=0;i<N;i+=1 {
		var A int
		fmt.Scanf("%d", &A)
		arr = append(arr, A)
	}

	for i:=len(arr)-1;i >=0;i-=1 {
		fmt.Println(arr[i])
	}
}

