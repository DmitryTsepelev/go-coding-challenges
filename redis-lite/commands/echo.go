package commands

import (
	"redis-lite/instance"
	"redis-lite/resp"
)

var echo = withCountOfArgs([]int{1}, func(instance *instance.Instance, args []string) (*string, error) {
	return resp.SerializeSimpleString(args[0])
})
