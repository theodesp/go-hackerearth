package basic_programming

import (
	"fmt"
)

/*
{a,e,i,o,u,A,E,I,O,U}

Natural Language Understanding is the subdomain of Natural Language Processing where people used to design AI based applications have ability to understand the human languages. HashInclude Speech Processing team has a project named Virtual Assistant. For this project they appointed you as a data engineer (who has good knowledge of creating clean datasets by writing efficient code). As a data engineer your first task is to make vowel recognition dataset. In this task you have to find the presence of vowels in all possible substrings of the given string. For each given string you have to print the total number of vowels.



Input

First line contains an integer T, denoting the number of test cases.

Each of the next lines contains a string, string contains both
lower case and upper case .

Output

Print the vowel sum
Answer for each test case should be printed in a new line.

Input Constraints

1<=T<=10

1<=|S|<=100000

Input
1
baceb


Output
16

Explanation
First line is number of input string, In given example, string is 
"baceb" so the substrings will be like -"b, ba, bac, bace, a, ac, ace, aceb, c, ce, ceb, e, eb, baceb" 
now the number of vowels in each substring will be 0, 1, 1, 2, 1, 1, 2, 2, 0, 1, 1, 1, 1, 2  
and the total number will be sum of all presence which is 16.
 */

/*
 Solution:
 Have counts that represent a list of substring vowel counts.
 Calculate substrings and on each iteration append a count if the new char is a vowel.
 At the end compute tha total sum of the counts slice.
 */

func PrintCountVowels()  {
	var N int
	var S string
	fmt.Scanf("%d", &N)

	for i := 0;i < N; i += 1 {
		fmt.Scanf("%s", &S)
		fmt.Printf("%d ", CountVowels(S))
	}
}

func CountVowels(s string) int  {
	runes := []rune(s)
	total := 0
	curr := 0
	counts := []int{}
	for i := 0;i < len(s); i += 1 {
		curr = 0
		for j := i+1; j <= len(s); j += 1 {
			if isVowel(runes[j-1]) {
				curr += 1
			}
			counts = append(counts, curr)
		}
	}

	for _, val := range counts {
		total += val
	}

	return total
}

func isVowel(word rune) bool  {
	switch word {
	case 'a', 'e', 'i', 'o', 'u', 'A', 'E', 'I', 'O', 'U':
		return true
	default:
		return false
	}
}

func Substrings(s string) []string {
	result := []string{}
	seen := make(map[string]interface{})
	for i := 0;i < len(s); i += 1 {
		for j := i+1; j <= len(s); j += 1 {
			sub := s[i:j]
			if _, ok := seen[sub]; !ok {
				result = append(result, sub)
				seen[sub] = struct {}{}
			}
		}
	}

	return result
}