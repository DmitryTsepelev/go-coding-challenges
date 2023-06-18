package commands

import (
	"redis-lite/instance"
	"redis-lite/resp"
	"strconv"
	"time"
)

var expire = withCountOfArgs([]int{2}, func(instance *instance.Instance, args []string) (*string, error) {
	key := args[0]
	seconds, err := strconv.Atoi(args[1])
	if err != nil {
		return resp.SerializeError("failed to convert " + args[1] + " to timestamp")
	}
	instance.Expire(key, time.Now().Add(time.Second*time.Duration(seconds)))

	return resp.SerializeSimpleString("OK")
})
