package commands

import (
	"redis-lite/instance"
	"redis-lite/resp"
	"strconv"
)

var decr = withCountOfArgs([]int{1}, func(instance *instance.Instance, args []string) (*string, error) {
	key := args[0]

	value := instance.Get(key)

	if value == nil {
		instance.Set(key, strconv.Itoa(-1))
	} else {
		intValue, err := strconv.Atoi(*value)
		if err != nil {
			return resp.SerializeError("failed to decr value " + *value)
		}

		instance.Set(key, strconv.Itoa(intValue-1))
	}

	return resp.SerializeSimpleString("OK")
})
