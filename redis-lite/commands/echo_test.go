package commands

import (
	"redis-lite/resp"
	"testing"
)

func TestEcho(t *testing.T) {
	hello := "hello"

	returnedValue, err := echo(nil, []string{hello})
	serializedValue, serializeErr := resp.SerializeSimpleString(hello)

	if *serializedValue != *returnedValue || (err != nil) || (serializeErr != nil) {
		t.Fatalf(`got %s, err %d`, *returnedValue, err)
	}
}

func TestEchoWithZeroArgs(t *testing.T) {
	returnedValue, err := echo(nil, []string{})
	serializedValue, serializeErr := resp.SerializeError("expected [1] arguments, given 0")

	if *serializedValue != *returnedValue || (err != nil) || (serializeErr != nil) {
		t.Fatalf(`got %s, err %d`, *returnedValue, err)
	}
}

func TestEchoWithTwoArgs(t *testing.T) {
	hello := "hello"
	world := "world"

	returnedValue, err := echo(nil, []string{hello, world})
	serializedValue, serializeErr := resp.SerializeError("expected [1] arguments, given 2")

	if *serializedValue != *returnedValue || (err != nil) || (serializeErr != nil) {
		t.Fatalf(`got %s, err %d`, *returnedValue, err)
	}
}
