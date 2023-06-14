package calculator

import (
	"calc/tokenizer"
	"testing"
)

func TestPlus(t *testing.T) {
	rpn := []interface{}{tokenizer.Operator{Kind: tokenizer.Plus}, tokenizer.IntNum{Value: 2}, tokenizer.IntNum{Value: 1}}
	result, err := Calculate(&rpn)
	expected := 3

	if (*result != expected) || (err != nil) {
		t.Fatalf(`expected %d got %d, err %d`, expected, *result, err)
	}
}

func TestMinus(t *testing.T) {
	rpn := []interface{}{tokenizer.Operator{Kind: tokenizer.Minus}, tokenizer.IntNum{Value: 1}, tokenizer.IntNum{Value: 2}}
	result, err := Calculate(&rpn)
	expected := 1

	if (*result != expected) || (err != nil) {
		t.Fatalf(`expected %d got %d, err %d`, expected, *result, err)
	}
}

func TestMult(t *testing.T) {
	rpn := []interface{}{tokenizer.Operator{Kind: tokenizer.Mult}, tokenizer.IntNum{Value: 2}, tokenizer.IntNum{Value: 3}}
	result, err := Calculate(&rpn)
	expected := 6

	if (*result != expected) || (err != nil) {
		t.Fatalf(`expected %d got %d, err %d`, expected, *result, err)
	}
}
