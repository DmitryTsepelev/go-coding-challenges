package main

import (
	"testing"
)

func TestPlus(t *testing.T) {
	expr := "1+1"
	expected := 2
	result, err := Execute(expr)

	if (*result != expected) || (err != nil) {
		t.Fatalf(`expected %d got %d, err %d`, expected, *result, err)
	}
}

func TestMinus(t *testing.T) {
	expr := "3-2"
	expected := 1
	result, err := Execute(expr)

	if (*result != expected) || (err != nil) {
		t.Fatalf(`expected %d got %d, err %d`, expected, *result, err)
	}
}

func TestMult(t *testing.T) {
	expr := "3*2"
	expected := 6
	result, err := Execute(expr)

	if (*result != expected) || (err != nil) {
		t.Fatalf(`expected %d got %d, err %d`, expected, *result, err)
	}
}

func TestComplex(t *testing.T) {
	expr := "3*2-1"
	expected := 5
	result, err := Execute(expr)

	if (*result != expected) || (err != nil) {
		t.Fatalf(`expected %d got %d, err %d`, expected, *result, err)
	}
}

func TestComplex2(t *testing.T) {
	expr := "3*2+4*3-10"
	expected := 8
	result, err := Execute(expr)

	if (*result != expected) || (err != nil) {
		t.Fatalf(`expected %d got %d, err %d`, expected, *result, err)
	}
}
