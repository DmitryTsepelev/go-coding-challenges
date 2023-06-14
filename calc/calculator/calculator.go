package calculator

import (
	"calc/stackUtils"
	"calc/tokenizer"
	"fmt"
)

func performBinary(stack *[]int, cb func(int, int) int) (*[]int, error) {
	var left *int
	var right *int
	var stackErr error
	var newStack = &([]int{})

	left, newStack, stackErr = stackUtils.Pop(stack)
	if stackErr != nil {
		println("stackErr", *stack)
		return nil, stackErr
	}

	right, newStack, stackErr = stackUtils.Pop(newStack)
	if stackErr != nil {
		return nil, stackErr
	}

	appendedStack := append(*newStack, cb(*left, *right))
	return &appendedStack, nil
}

func Calculate(input *[]interface{}) (*int, error) {
	var token *interface{}
	var stack = &([]int{})

	println("calculate!", *input)

	for {
		var stackErr error
		token, input, stackErr = stackUtils.Pop(input)

		fmt.Println("stack")
		for _, el := range *stack {
			fmt.Print(el, ",")
		}
		fmt.Println("")

		if stackErr != nil {
			// TODO handle stack not empty
			fmt.Println("!!!", stack)
			newStack := *stack
			return &newStack[0], nil
		}

		if number, ok := (*token).(tokenizer.IntNum); ok {
			newStack := append(*stack, number.Value)
			println("number.Value", number.Value)
			stack = &newStack
		} else if operator, ok := (*token).(tokenizer.Operator); ok {
			println("operator.Kind", operator.Kind)
			switch operator.Kind {
			case tokenizer.Plus:
				{
					stack, stackErr = performBinary(stack, func(left int, right int) int {
						return left + right
					})

					if stackErr != nil {
						return nil, stackErr
					}
				}
			case tokenizer.Minus:
				{
					stack, stackErr = performBinary(stack, func(left int, right int) int {
						return right - left
					})

					if stackErr != nil {
						return nil, stackErr
					}
				}
			case tokenizer.Mult:
				{
					stack, stackErr = performBinary(stack, func(left int, right int) int {
						return left * right
					})

					if stackErr != nil {
						return nil, stackErr
					}
				}
			default:
				{
					panic("unsupported operator")
				}
			}
		}
	}
}
