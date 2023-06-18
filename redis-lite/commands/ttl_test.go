package commands

import (
	"redis-lite/instance"
	"redis-lite/resp"
	"strconv"
	"testing"
	"time"
)

func TestTtl(t *testing.T) {
	instance := instance.MakeInstance()
	key := "k"
	duration := time.Now().Add(time.Second * time.Duration(10))
	instance.Expire(key, duration)

	returnedValue, err := ttl(instance, []string{key})

	seconds := strconv.Itoa(-int(time.Since(duration).Seconds()))
	expectedValue, err2 := resp.SerializeBulkString(&seconds)

	if *returnedValue != *expectedValue || (err != nil) || (err2 != nil) {
		t.Fatalf(`got %s`, *returnedValue)
	}
}

func TestTtlWithoutKey(t *testing.T) {
	instance := instance.MakeInstance()

	response, err := ttl(instance, []string{})
	expectedErr, err2 := resp.SerializeError("expected [1] arguments, given 0")

	if *response != *expectedErr || (err != nil) || (err2 != nil) {
		t.Fatalf(`returned wrong error`)
	}
}
