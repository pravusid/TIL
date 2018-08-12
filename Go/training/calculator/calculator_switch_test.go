package calculator

import "fmt"

func ExampleEval() {
	fmt.Println(Eval("5"))
	fmt.Println(Eval("1 + 2"))
	// Output:
	// 5
	// 3
}
