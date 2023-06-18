package resp

import "testing"

func TestSerializeSimpleString(t *testing.T) {
	input := "OK"
	result, err := SerializeSimpleString(input)

	if (*result != "+OK\r\n") || (err != nil) {
		t.Fatalf(`got %s, err %d`, *result, err)
	}
}

func TestSerializeBulkString(t *testing.T) {
	input := "OK"
	result, err := SerializeBulkString(&input)

	if (*result != "$2\r\nOK\r\n") || (err != nil) {
		t.Fatalf(`got %s, err %d`, *result, err)
	}
}

func TestSerializeBulkStringNull(t *testing.T) {
	result, err := SerializeBulkString(nil)

	if (*result != "$-1\r\n") || (err != nil) {
		t.Fatalf(`got %s, err %d`, *result, err)
	}
}

func TestSerializeError(t *testing.T) {
	input := "Wrong"
	result, err := SerializeError(input)

	if (*result != "-Wrong\r\n") || (err != nil) {
		t.Fatalf(`got %s, err %d`, *result, err)
	}
}
