// https://codingchallenges.fyi/challenges/challenge-calculator

package main

import (
	"calc/calculator"
	"calc/rpn"
	"calc/tokenizer"
	"fmt"
	"log"
	"os"
)

func Execute(expression string) (*int, error) {
	tokens, tokenizerErr := tokenizer.Tokenize([]rune(expression))
	if tokenizerErr != nil {
		return nil, tokenizerErr
	}

	rpn, rpnError := rpn.TokensToRpn(tokens)
	if rpnError != nil {
		return nil, rpnError
	}

	result, calculateErr := calculator.Calculate(&rpn)
	if rpnError != nil {
		return nil, calculateErr
	}

	return result, nil
}

func main() {
	result, err := Execute(os.Args[1])
	if err != nil {
		log.Fatalln(err)
		return
	}
	fmt.Println(*result)
}
