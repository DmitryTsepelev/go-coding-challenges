package stackUtils

import "fmt"

func QueueToStack[T interface{}](list *([]T)) *[]T {
	reversed := []T{}

	for idx := len(*list) - 1; idx >= 0; idx-- {
		el := (*list)[idx]
		reversed = *Push(&el, &reversed)
	}

	return &reversed
}

func Push[T interface{}](el *T, stack *([]T)) *[]T {
	newStack := append(*stack, *el)
	return &newStack
}

func Pop[T interface{}](stack *([]T)) (*T, *[]T, error) {
	if len(*stack) == 0 {
		return nil, nil, fmt.Errorf("stack is empty")
	}

	el := (*stack)[len(*stack)-1]

	if len(*stack) == 1 {
		rest := (*stack)[0:0]
		return &el, &rest, nil
	} else {
		rest := (*stack)[0 : len(*stack)-1]
		return &el, &rest, nil
	}
}

func Top[T interface{}](stack *([]T)) (*T, error) {
	if len(*stack) == 0 {
		return nil, fmt.Errorf("stack is empty")
	}

	el := (*stack)[len(*stack)-1]

	return &el, nil
}
