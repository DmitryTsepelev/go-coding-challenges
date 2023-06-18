package commands

import (
	"redis-lite/instance"
	"redis-lite/resp"
)

const PONG = "PONG"

func ping(_ *instance.Instance, args []string) (*string, error) {
	return resp.SerializeSimpleString(PONG)
}
