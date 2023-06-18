package instance

import (
	"testing"
	"time"
)

func TestSet(t *testing.T) {
	instance := MakeInstance()
	key := "k"
	value := "v"
	instance.Set(key, value)

	if *(*instance.KV)[key] != value {
		t.Fatalf(`did not set value`)
	}
}

func TestGet(t *testing.T) {
	instance := MakeInstance()

	key := "k"

	nilValue := instance.Get(key)

	value := "v"
	(*instance.KV)[key] = &value

	fetchedValue := instance.Get(key)

	if *fetchedValue != value || nilValue != nil {
		t.Fatalf(`did not get value`)
	}
}

func TestExpireAndTtl(t *testing.T) {
	instance := MakeInstance()

	key := "k"
	noTtl := instance.Ttl(key)

	ex := time.Now().Add(time.Second * time.Duration(10))
	instance.Expire(key, ex)

	ttl := instance.Ttl(key)

	if *ttl != ex || noTtl != nil {
		t.Fatalf(`did not get expiration`)
	}
}

func TestDel(t *testing.T) {
	instance := MakeInstance()

	key := "k"
	value := "v"
	(*instance.KV)[key] = &value

	instance.Del(key)

	fetchedValue := instance.Get(key)

	if fetchedValue != nil {
		t.Fatalf(`did not get value`)
	}
}
