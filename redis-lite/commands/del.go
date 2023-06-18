package commands

import (
	"redis-lite/instance"
	"redis-lite/resp"
)

var del = withCountOfArgs([]int{1}, func(inst *instance.Instance, args []string) (*string, error) {
	key := args[0]

	inst.Del(key)

	return resp.SerializeSimpleString("OK")
})
