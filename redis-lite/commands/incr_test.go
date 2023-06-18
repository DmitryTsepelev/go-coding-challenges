package commands

import (
	"redis-lite/instance"
	"redis-lite/resp"
	"testing"
)

func TestIncr(t *testing.T) {
	instance := instance.MakeInstance()
	key := "k"

	_, err := incr(instance, []string{key})
	zeroValue := instance.Get(key)
	_, err2 := incr(instance, []string{key})
	oneValue := instance.Get(key)

	if *zeroValue != "1" || *oneValue != "2" || (err != nil) || (err2 != nil) {
		t.Fatalf(`did not incr value`)
	}
}

func TestIncrWithoutKey(t *testing.T) {
	instance := instance.MakeInstance()

	response, err := incr(instance, []string{})
	expectedErr, err2 := resp.SerializeError("expected [1] arguments, given 0")

	if *response != *expectedErr || (err != nil) || (err2 != nil) {
		t.Fatalf(`returned wrong error`)
	}
}

func TestIncrNotIntValue(t *testing.T) {
	instance := instance.MakeInstance()
	key := "k"
	value := "wrong"

	instance.Set(key, value)

	result, err := incr(instance, []string{key})
	expectedErr, _ := resp.SerializeError("failed to incr value " + value)

	if *result != *expectedErr || (err != nil) {
		t.Fatalf(`did not incr value`)
	}
}
