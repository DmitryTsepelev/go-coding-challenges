package tokenizer

import (
	"fmt"
	"strconv"
	"unicode"
)

type IntNum struct {
	Value int
}

type OperatorKind uint8

const (
	Plus  OperatorKind = 0
	Minus OperatorKind = 1
	Mult  OperatorKind = 2
)

type Operator struct {
	Kind OperatorKind
}

func parseOperator(operatorRune rune, currentNumber []rune) (*Operator, *IntNum, error) {
	var numPtr *IntNum

	if len(currentNumber) > 0 {
		value, _ := strconv.Atoi(string(currentNumber))
		numPtr = &IntNum{Value: value}
	}

	switch operatorRune {
	case '+':
		return &Operator{Kind: Plus}, numPtr, nil
	case '-':
		return &Operator{Kind: Minus}, numPtr, nil
	case '*':
		return &Operator{Kind: Mult}, numPtr, nil
	default:
		return nil, nil, fmt.Errorf("unexpected operator %d", operatorRune)
	}
}

func Tokenize(expression []rune) ([]interface{}, error) {
	idx := 0
	var tokens []interface{}

	currentNumber := []rune{}

	for {
		if idx == len(expression) {
			if len(currentNumber) > 0 {
				value, _ := strconv.Atoi(string(currentNumber))
				tokens = append(tokens, IntNum{Value: value})
			}

			break
		}

		currentRune := expression[idx]

		if unicode.IsSpace(currentRune) {
			// do nothing
		} else if unicode.IsDigit(currentRune) {
			currentNumber = append(currentNumber, currentRune)
		} else {
			operatorPtr, numPtr, err := parseOperator(currentRune, currentNumber)
			if err != nil {
				return nil, err
			}
			tokens = append(tokens, *numPtr)
			tokens = append(tokens, *operatorPtr)
			currentNumber = []rune{}
		}

		idx += 1
	}

	return tokens, nil
}
