package commands

import (
	"redis-lite/instance"
	"redis-lite/resp"
	"testing"
)

func TestSet(t *testing.T) {
	instance := instance.MakeInstance()
	key := "k"
	value := "v"
	_, err := set(instance, []string{key, value})

	if *(*instance.KV)[key] != value || (err != nil) {
		t.Fatalf(`did not set value, err %d`, err)
	}
}

func TestSetEx(t *testing.T) {
	instance := instance.MakeInstance()
	key := "k"
	value := "v"
	seconds := "10"
	_, err := set(instance, []string{key, value, "EX", seconds})

	if *(*instance.KV)[key] != value || (instance.Ttl(key) == nil) || (err != nil) {
		t.Fatalf(`did not set value, err %d`, err)
	}
}

func TestSetPx(t *testing.T) {
	instance := instance.MakeInstance()
	key := "k"
	value := "v"
	seconds := "10"
	_, err := set(instance, []string{key, value, "PX", seconds})

	if *(*instance.KV)[key] != value || (instance.Ttl(key) == nil) || (err != nil) {
		t.Fatalf(`did not set value, err %d`, err)
	}
}

func TestSetExat(t *testing.T) {
	instance := instance.MakeInstance()
	key := "k"
	value := "v"
	unixSeconds := "1617875638"
	_, err := set(instance, []string{key, value, "EXAT", unixSeconds})

	if *(*instance.KV)[key] != value || (instance.Ttl(key) == nil) || (err != nil) {
		t.Fatalf(`did not set value, err %d`, err)
	}
}

func TestSetExWrongArg(t *testing.T) {
	instance := instance.MakeInstance()
	key := "k"
	value := "v"
	seconds := "qwe"
	returnedValue, err := set(instance, []string{key, value, "EX", seconds})
	expectedError, _ := resp.SerializeError("failed to convert " + seconds + " to timestamp")

	if (*returnedValue != *expectedError) || (err != nil) {
		t.Fatalf(`wrong err`)
	}
}

func TestSetExMissingArg(t *testing.T) {
	instance := instance.MakeInstance()
	key := "k"
	value := "v"
	returnedValue, err := set(instance, []string{key, value, "EX"})
	expectedError, _ := resp.SerializeError("expected [2 4] arguments, given 3")

	if (*returnedValue != *expectedError) || (err != nil) {
		t.Fatalf(`wrong err`)
	}
}
