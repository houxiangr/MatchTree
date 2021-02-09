package main

import(
	"fmt"
	"github.com/Knetic/govaluate"
)

func main(){
	expression, err := govaluate.NewEvaluableExpression("10 > 0")
	if err != nil {
		fmt.Println(err)
	}
	result, err := expression.Evaluate(nil)
	fmt.Println(result)
}
