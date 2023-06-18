package commands

import (
	"redis-lite/resp"
	"testing"
)

func TestPing(t *testing.T) {
	returnedValue, err := ping(nil, []string{})
	serializedValue, serializeErr := resp.SerializeSimpleString("PONG")

	if *serializedValue != *returnedValue || (err != nil) || (serializeErr != nil) {
		t.Fatalf(`got %s, err %d`, *returnedValue, err)
	}
}
