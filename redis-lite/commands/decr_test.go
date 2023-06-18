package commands

import (
	"redis-lite/instance"
	"redis-lite/resp"
	"testing"
)

func TestDecr(t *testing.T) {
	instance := instance.MakeInstance()
	key := "k"

	_, err := decr(instance, []string{key})
	zeroValue := instance.Get(key)
	_, err2 := decr(instance, []string{key})
	oneValue := instance.Get(key)

	if *zeroValue != "-1" || *oneValue != "-2" || (err != nil) || (err2 != nil) {
		t.Fatalf(`did not decr value`)
	}
}

func TestDecrWithoutKey(t *testing.T) {
	instance := instance.MakeInstance()

	response, err := decr(instance, []string{})
	expectedErr, err2 := resp.SerializeError("expected [1] arguments, given 0")

	if *response != *expectedErr || (err != nil) || (err2 != nil) {
		t.Fatalf(`returned wrong error`)
	}
}

func TestDecrNotIntValue(t *testing.T) {
	instance := instance.MakeInstance()
	key := "k"
	value := "wrong"

	instance.Set(key, value)

	result, err := decr(instance, []string{key})
	expectedErr, _ := resp.SerializeError("failed to decr value " + value)

	if *result != *expectedErr || (err != nil) {
		t.Fatalf(`did not decr value`)
	}
}
