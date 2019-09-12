package basic_programming

import "fmt"

/*
Input:
First line contains T, number of test cases.
Each test case contains two integers, N and M. where is N is number of friends and M is number number of chocolates in a packet.

Output:
In each test case output "Yes" if he can buy that packet and "No" if he can't buy that packet.

Constraints:
1<=T<=20
1<=N<=100
1<=M<=10^5

Input
2
5 14
3 21

OUTPUT
No
Yes

Explanation
Test Case 1:
There is no way such that he can distribute 14 chocolates among 5 friends equally.
Test Case 2:
There are 21 chocolates and 3 friends, so he can distribute chocolates eqally.
Each friend will get 7 chocolates.
 */

func PrintBirthdayParty()  {
	var N int
	fmt.Scanf("%d", &N)

	for i := 0;i < N; i += 1 {
		var friends int
		var chocolates int
		fmt.Scanf("%d %d", &friends, &chocolates)
		if EquallyDivides(friends, chocolates) {
			fmt.Println("Yes")
		} else {
			fmt.Println("No")
		}
	}
}


func EquallyDivides(nom, denom int) bool   {
	if denom % nom == 0 {
		return true
	}
	return false
}