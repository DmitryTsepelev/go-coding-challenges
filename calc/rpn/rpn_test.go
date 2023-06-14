package rpn

import (
	"calc/tokenizer"
	"reflect"
	"testing"
)

func TestPlus(t *testing.T) {
	tokens := []interface{}{tokenizer.IntNum{Value: 1}, tokenizer.Operator{Kind: tokenizer.Plus}, tokenizer.IntNum{Value: 1}}
	expected := []interface{}{tokenizer.Operator{Kind: tokenizer.Plus}, tokenizer.IntNum{Value: 1}, tokenizer.IntNum{Value: 1}}
	tokens, err := TokensToRpn(tokens)

	if (!reflect.DeepEqual(tokens, expected)) || (err != nil) {
		t.Fatalf(`expected %d got %d, err %d`, expected, tokens, err)
	}
}

func TestMinus(t *testing.T) {
	tokens := []interface{}{tokenizer.IntNum{Value: 1}, tokenizer.Operator{Kind: tokenizer.Minus}, tokenizer.IntNum{Value: 1}}
	expected := []interface{}{tokenizer.Operator{Kind: tokenizer.Minus}, tokenizer.IntNum{Value: 1}, tokenizer.IntNum{Value: 1}}
	tokens, err := TokensToRpn(tokens)

	if (!reflect.DeepEqual(tokens, expected)) || (err != nil) {
		t.Fatalf(`expected %d got %d, err %d`, expected, tokens, err)
	}
}

func TestMult(t *testing.T) {
	tokens := []interface{}{tokenizer.IntNum{Value: 1}, tokenizer.Operator{Kind: tokenizer.Mult}, tokenizer.IntNum{Value: 1}}
	expected := []interface{}{tokenizer.Operator{Kind: tokenizer.Mult}, tokenizer.IntNum{Value: 1}, tokenizer.IntNum{Value: 1}}
	tokens, err := TokensToRpn(tokens)

	if (!reflect.DeepEqual(tokens, expected)) || (err != nil) {
		t.Fatalf(`expected %d got %d, err %d`, expected, tokens, err)
	}
}

func TestMultAndPlus(t *testing.T) {
	tokens := []interface{}{
		tokenizer.IntNum{Value: 7},
		tokenizer.Operator{Kind: tokenizer.Minus},
		tokenizer.IntNum{Value: 2},
		tokenizer.Operator{Kind: tokenizer.Mult},
		tokenizer.IntNum{Value: 3}}
	expected := []interface{}{
		tokenizer.Operator{Kind: tokenizer.Minus},
		tokenizer.Operator{Kind: tokenizer.Mult},
		tokenizer.IntNum{Value: 3},
		tokenizer.IntNum{Value: 2},
		tokenizer.IntNum{Value: 7}}
	tokens, err := TokensToRpn(tokens)

	if (!reflect.DeepEqual(tokens, expected)) || (err != nil) {
		t.Fatalf(`expected %d got %d, err %d`, expected, tokens, err)
	}
}
