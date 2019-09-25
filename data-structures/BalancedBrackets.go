package data_structures

import "fmt"

/*
A bracket is considered to be any one of the following characters: (, ), {, }, [, or ].

Two brackets are considered to be a matched pair if the an opening bracket
(i.e., (, [, or {) occurs to the left of a closing bracket (i.e., ), ], or })
of the exact same type. There are three types of matched pairs of brackets: [], {},and ().

A matching pair of brackets is not balanced if the set of brackets it
encloses are not matched. For example, {[(])} is not balanced because the
contents in between { and } are not balanced.
The pair of square brackets encloses a single, unbalanced opening bracket,
(, and the pair of parentheses encloses a single, unbalanced closing square bracket, ].

By this logic, we say a sequence of brackets is balanced if
the following conditions are met:

It contains no unmatched brackets.
The subset of brackets enclosed within the confines of a matched
pair of brackets is also a matched pair of brackets.
Given n strings of brackets, determine whether each sequence of
brackets is balanced. If a string is balanced, return YES. Otherwise, return NO.
 */

/*
AMPLE INPUT
3
{[()]}
{[(])}
{{[[(())]]}}
SAMPLE OUTPUT
YES
NO
YES
Explanation
1.The string {[()]}
meets both criteria for being a balanced string, so we print YES on a new line.

2.The string {[(])} is not balanced because the brackets enclosed by the
matched pair { and } are not balanced: [(]).

3.The string {{[[(())]]}} meets both criteria for being a balanced string,
so we print YES on a new line.
 */

// Solution: Have a mapping of valid balanced bracket pairs.
// For each char we find use a stack to push if not matched. If matched we pop from the stack.
// At the end if we have an empty stack we return true. Otherwise we return false.

func PrintIsBalancedBrackets()  {
	var N int
	fmt.Scanf("%d", &N)

	for i:=0;i<N;i+=1 {
		var S string
		fmt.Scanf("%s", &S)
		if IsBalanced(S) {
			fmt.Println("YES")
		} else {
			fmt.Println("NO")
		}
	}
}

func IsBalanced(s string) bool  {
	stack := []rune{}
	mapping := make(map[rune]rune)
	mapping['{'] = '}'
	mapping['('] = ')'
	mapping['['] = ']'
	mapping['}'] = '{'
	mapping[')'] = '('
	mapping[']'] = '['
	for _, val := range s {
		if len(stack) == 0 {
			stack = append([]rune{val}, stack...)
		} else {
			if val == '{' || val == '(' || val == '[' {
				stack = append([]rune{val}, stack...)
			} else if val == mapping[stack[0]] {
				stack = stack[1:]
			} else {
				stack = append([]rune{val}, stack...)
			}
		}
	}
	return len(stack) == 0
}