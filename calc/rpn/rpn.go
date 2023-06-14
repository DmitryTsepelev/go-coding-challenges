package rpn

import (
	"calc/stackUtils"
	"calc/tokenizer"
	"fmt"
)

// https://en.wikipedia.org/wiki/Shunting_yard_algorithm
func TokensToRpn(tokens []interface{}) ([]interface{}, error) {
	var queue []interface{}
	var stack []interface{}

	for _, token := range tokens {
		if number, ok := token.(tokenizer.IntNum); ok {
			queue = append(queue, number)
		} else if operator, ok := token.(tokenizer.Operator); ok {
			for {
				topOperatorInterface, stackErr := stackUtils.Top(&stack)
				if stackErr != nil {
					break
				}

				topOperator, ok := (*topOperatorInterface).(tokenizer.Operator)

				if !ok {
					return nil, fmt.Errorf("failed to convert")
				}

				if priority(topOperator) > priority(operator) {
					topOperator, newStack, _ := stackUtils.Pop(&stack)
					stack = *newStack
					queue = append(queue, *topOperator)
				} else {
					break
				}
			}
			stack = append(stack, operator)
		}
	}

	for {
		operator, newStack, err := stackUtils.Pop(&stack)
		if err != nil {
			break
		}
		queue = append(queue, *operator)
		stack = *newStack
	}

	return *stackUtils.QueueToStack(&queue), nil
}

func priority(operator tokenizer.Operator) int {
	switch operator.Kind {
	case tokenizer.Plus:
		return 0
	case tokenizer.Minus:
		return 1
	case tokenizer.Mult:
		return 2
	default:
		panic("unexpected operator")
	}
}
