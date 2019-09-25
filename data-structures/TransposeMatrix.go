package data_structures

import (
	"fmt"
)

/*
Given a matrix A, return the transpose of A.

The transpose of a matrix is the matrix flipped over it's main diagonal,
switching the row and column indices of the matrix.

Example 1:

Input: [[1,2,3],[4,5,6],[7,8,9]]
Output: [[1,4,7],[2,5,8],[3,6,9]]
Example 2:

Input: [[1,2,3],[4,5,6]]
Output: [[1,4],[2,5],[3,6]]
 */

func PrintTranspose()  {
	var A int
	var B int
	fmt.Scanf("%d %d", &A, &B)
	arr := make([][]int, A, A)
	for i:=0;i<A;i+=1 {
		for j:=0;j<B;j+=1 {
			var num int
			fmt.Scanf("%d", &num)
			arr[i] = append(arr[i], num)
		}

	}

	transposed := Transpose(arr)
	for i := 0; i < len(transposed); i += 1 {
		for j := 0; j < len(transposed[0]); j += 1 {
			fmt.Printf("%d ", transposed[i][j])
		}
		fmt.Println()
	}
}

func makeMatrix(n, m int) [][]int {
	board := make([][]int, m, m)
	for i := int(0); i < m; i += 1 {
		board[i] = make([]int, n, n)
	}
	return board
}

/*
 Solution: iterate over A[i][0] to A[i][j] and assign result[i][j] = A[j][i]
 Time Complexity: O(R * C),
 where RR and CC are the number of rows and columns in the given matrix A.

 Space Complexity: O(R * C), the space used by the answer.
 */
func Transpose(A [][]int) [][]int {
	result := makeMatrix(int(len(A)), int(len(A[0])))

	for i:=0;i<len(result);i+=1 {
		for j:=0;j<len(result[0]);j+=1 {
			result[i][j] = A[j][i]
		}
	}
	return result
}

