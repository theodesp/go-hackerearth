package data_structures

import "fmt"

/*
Micro just purchased a queue and wants to perform N operations on the queue. The operations are of following type:

 : Enqueue x in the queue and print the new size of the queue.
D : Dequeue from the queue and print the element that is deleted and the new size of the queue separated by space.
If there is no element in the queue then print 1 in place of deleted element.

Since Micro is new with queue data structure, he need your help.

Input:
First line consists of a single integer denoting N
Each of the following N lines contain one of the operation as described above.

Output:
For each enqueue operation print the new size of the queue.
And for each dequeue operation print two integers, deleted element
(1, if queue is empty) and the new size of the queue.
 */

type Op string
const (
	ENQ Op = "E"
	DEL = "D"
)

func PrintQueueOps()  {
	var N int
	fmt.Scanf("%d", &N)
	var queue []int

	for i:=0;i<N;i+=1 {
		var op Op
		fmt.Scanf("%s", &op)
		switch op {
		case ENQ:
			var x int
			fmt.Scanf("%d", &x)
			queue = append(queue, x)
			fmt.Println(len(queue))
		case DEL:
			if len(queue) == 0 {
				fmt.Println(-1, 0)
			} else {
				front, rest := queue[0], queue[1:]
				fmt.Println(front, len(rest))
				queue = rest
			}
		}
	}
}

