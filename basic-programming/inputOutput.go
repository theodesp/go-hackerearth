package basic_programming

import "fmt"

func InputOutput()  {
	var N int
	var name string
	fmt.Scanf("%d", &N)
	fmt.Scanf("%s", &name)

	fmt.Println(N * 2)
	fmt.Println(name)
}
