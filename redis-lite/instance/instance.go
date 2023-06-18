package instance

import (
	"math/rand"
	"sync"
	"time"
)

type Instance struct {
	KV      *map[string]*string
	expires *map[string]*time.Time
	mutex   *sync.RWMutex
}

func MakeInstance() *Instance {
	kv := make(map[string]*string, 0)
	expires := make(map[string]*time.Time, 0)
	mutex := sync.RWMutex{}

	instance := Instance{KV: &kv, mutex: &mutex, expires: &expires}

	return &instance
}

func (instance *Instance) Set(key string, value string) {
	(*instance.mutex).Lock()
	(*instance.KV)[key] = &value
	(*instance.mutex).Unlock()
}

func (instance *Instance) Get(key string) *string {
	(*instance.mutex).Lock()

	if instance.IsExpired(key) {
		instance.Del(key)
		(*instance.mutex).Unlock()
		return nil
	}

	value := (*instance.KV)[key]
	(*instance.mutex).Unlock()
	return value
}

func (instance *Instance) IsExpired(key string) bool {
	expiration := (*instance.expires)[key]
	if expiration == nil {
		return false
	}

	return time.Now().Compare(*expiration) == +1
}

func (instance *Instance) Expire(key string, expiration time.Time) {
	(*instance.mutex).Lock()
	(*instance.expires)[key] = &expiration
	(*instance.mutex).Unlock()
}

func (instance *Instance) Ttl(key string) *time.Time {
	(*instance.mutex).Lock()
	ttl := (*instance.expires)[key]
	(*instance.mutex).Unlock()
	return ttl
}

func (instance *Instance) Del(key string) {
	(*instance.mutex).Lock()
	instance.delUnstafe(key)
	(*instance.mutex).Unlock()
}

func (instance *Instance) delUnstafe(key string) {
	delete((*instance.expires), key)
	delete((*instance.KV), key)
}

func (instance *Instance) RunExpirationChecker() {
	for {
		(*instance.mutex).Lock()

		keys := make([]string, 0, len(*instance.expires)/10)
		for key := range *instance.KV {
			if rand.Intn(100) < 10 {
				keys = append(keys, key)
			}
		}

		expiredCount := 0
		for _, key := range keys {
			if instance.IsExpired(key) {
				expiredCount += 1
				instance.delUnstafe(key)
			}
		}

		(*instance.mutex).Unlock()

		if len(keys) == 0 || expiredCount < len(keys)/10 {
			time.Sleep(time.Duration(10) * time.Second)
		}
	}
}
