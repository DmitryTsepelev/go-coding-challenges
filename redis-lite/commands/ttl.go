package commands

import (
	"redis-lite/instance"
	"redis-lite/resp"
	"strconv"
	"time"
)

var ttl = withCountOfArgs([]int{1}, func(instance *instance.Instance, args []string) (*string, error) {
	ttl := instance.Ttl(args[0])

	if ttl == nil {
		return resp.SerializeBulkString(nil)
	}

	seconds := strconv.Itoa(-int(time.Since(*ttl).Seconds()))
	return resp.SerializeBulkString(&seconds)
})
