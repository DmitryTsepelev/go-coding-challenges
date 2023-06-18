package commands

import (
	"redis-lite/instance"
	"redis-lite/resp"
	"testing"
)

func TestGet(t *testing.T) {
	instance := instance.MakeInstance()
	key := "k"
	value := "v"
	(*instance.KV)[key] = &value

	returnedValue, err := get(instance, []string{key})
	serializedValue, serializeErr := resp.SerializeBulkString(&value)

	if *serializedValue != *returnedValue || (err != nil) || (serializeErr != nil) {
		t.Fatalf(`got %s, err %d`, *returnedValue, err)
	}
}

func TestGetWithoutKey(t *testing.T) {
	instance := instance.MakeInstance()

	response, err := get(instance, []string{})
	expectedErr, err2 := resp.SerializeError("expected [1] arguments, given 0")

	if *response != *expectedErr || (err != nil) || (err2 != nil) {
		t.Fatalf(`returned wrong error`)
	}
}
