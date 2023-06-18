package commands

import (
	"redis-lite/instance"
	"redis-lite/resp"
)

var get = withCountOfArgs([]int{1}, func(instance *instance.Instance, args []string) (*string, error) {
	value := instance.Get(args[0])
	return resp.SerializeBulkString(value)
})
