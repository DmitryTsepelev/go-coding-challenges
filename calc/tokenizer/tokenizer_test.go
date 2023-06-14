package tokenizer

import (
	"reflect"
	"testing"
)

func TestPlus(t *testing.T) {
	expr := "1+1"
	expected := []interface{}{IntNum{Value: 1}, Operator{Kind: Plus}, IntNum{Value: 1}}
	tokens, err := Tokenize([]rune(expr))

	if (!reflect.DeepEqual(tokens, expected)) || (err != nil) {
		t.Fatalf(`expected %d got %d, err %d`, expected, tokens, err)
	}
}

func TestPlusWithSpaces(t *testing.T) {
	expr := "1 + 1"
	expected := []interface{}{IntNum{Value: 1}, Operator{Kind: Plus}, IntNum{Value: 1}}
	tokens, err := Tokenize([]rune(expr))

	if (!reflect.DeepEqual(tokens, expected)) || (err != nil) {
		t.Fatalf(`expected %d got %d, err %d`, expected, tokens, err)
	}
}

func TestPlusWithManySpaces(t *testing.T) {
	expr := "  1  +  1  "
	expected := []interface{}{IntNum{Value: 1}, Operator{Kind: Plus}, IntNum{Value: 1}}
	tokens, err := Tokenize([]rune(expr))

	if (!reflect.DeepEqual(tokens, expected)) || (err != nil) {
		t.Fatalf(`expected %d got %d, err %d`, expected, tokens, err)
	}
}

func TestMinus(t *testing.T) {
	expr := "1-1"
	expected := []interface{}{IntNum{Value: 1}, Operator{Kind: Minus}, IntNum{Value: 1}}
	tokens, err := Tokenize([]rune(expr))

	if (!reflect.DeepEqual(tokens, expected)) || (err != nil) {
		t.Fatalf(`expected %d got %d, err %d`, expected, tokens, err)
	}
}

func TestMult(t *testing.T) {
	expr := "1*1"
	expected := []interface{}{IntNum{Value: 1}, Operator{Kind: Mult}, IntNum{Value: 1}}
	tokens, err := Tokenize([]rune(expr))

	if (!reflect.DeepEqual(tokens, expected)) || (err != nil) {
		t.Fatalf(`expected %d got %d, err %d`, expected, tokens, err)
	}
}
