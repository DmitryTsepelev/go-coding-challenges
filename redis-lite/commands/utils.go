package commands

import (
	"fmt"
	"redis-lite/instance"
	"redis-lite/resp"
)

type HandlerFn = func(*instance.Instance, []string) (*string, error)

func withCountOfArgs(counts []int, cb HandlerFn) HandlerFn {
	return func(inst *instance.Instance, args []string) (*string, error) {
		validArgs := false
		for _, count := range counts {
			if count == len(args) {
				validArgs = true
				break
			}
		}

		if validArgs {
			return cb(inst, args)
		} else {
			return resp.SerializeError(fmt.Sprintf("expected %d arguments, given %d", counts, len(args)))
		}
	}
}
