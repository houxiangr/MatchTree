package main

import (
	"fmt"
	"github.com/Knetic/govaluate"
)

func main() {
	expression, err := govaluate.NewEvaluableExpression("10 in (10,20)")
	if err != nil {
		fmt.Println(err)
	}
	result, err := expression.Evaluate(nil)
	fmt.Println(err)
	fmt.Println(result)
}
