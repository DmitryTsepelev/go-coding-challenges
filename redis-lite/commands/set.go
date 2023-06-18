package commands

import (
	"redis-lite/instance"
	"redis-lite/resp"
	"strconv"
	"strings"
	"time"
)

const EX = "EX"
const PX = "PX"
const EXAT = "EXAT"

var set = withCountOfArgs([]int{2, 4}, func(instance *instance.Instance, args []string) (*string, error) {
	key := args[0]
	value := args[1]

	instance.Set(key, value)

	if len(args) == 4 {
		expirationOption := strings.ToUpper(args[2])

		// TODO: copy logic to expire
		switch expirationOption {
		case EX:
			{
				seconds, err := strconv.Atoi(args[3])
				if err != nil {
					return resp.SerializeError("failed to convert " + args[3] + " to timestamp")
				}
				instance.Expire(key, time.Now().Add(time.Second*time.Duration(seconds)))
			}
		case PX:
			{
				milliseconds, err := strconv.Atoi(args[3])
				if err != nil {
					return resp.SerializeError("failed to convert " + args[3] + " to timestamp")
				}
				instance.Expire(key, time.Now().Add(time.Millisecond*time.Duration(milliseconds)))
			}
		case EXAT:
			{
				unixSeconds, err := strconv.ParseInt(args[3], 10, 64)
				if err != nil {
					return resp.SerializeError("failed to convert " + args[3] + " to timestamp")
				}
				instance.Expire(key, time.Unix(unixSeconds, 0))
			}
		default:
			{
				return resp.SerializeError("unexpected option")
			}
		}
	}

	return resp.SerializeSimpleString("OK")
})

// func set(instance *instance.Instance, args []string) (*string, error) {
// 	key := args[0]
// 	value := args[1]

// 	instance.Set(key, value)

// 	if len(args) > 2 {
// 		expirationOption := strings.ToUpper(args[2])
// 		if len(args) == 3 {
// 			return resp.SerializeError("expected argument after " + expirationOption)
// 		}

// 		// TODO: copy logic to expire
// 		switch expirationOption {
// 		case EX:
// 			{
// 				seconds, err := strconv.Atoi(args[3])
// 				if err != nil {
// 					return resp.SerializeError("failed to convert " + args[3] + " to timestamp")
// 				}
// 				instance.Expire(key, time.Now().Add(time.Second*time.Duration(seconds)))
// 			}
// 		case PX:
// 			{
// 				milliseconds, err := strconv.Atoi(args[3])
// 				if err != nil {
// 					return resp.SerializeError("failed to convert " + args[3] + " to timestamp")
// 				}
// 				instance.Expire(key, time.Now().Add(time.Millisecond*time.Duration(milliseconds)))
// 			}
// 		case EXAT:
// 			{
// 				unixSeconds, err := strconv.ParseInt(args[3], 10, 64)
// 				if err != nil {
// 					return resp.SerializeError("failed to convert " + args[3] + " to timestamp")
// 				}
// 				instance.Expire(key, time.Unix(unixSeconds, 0))
// 			}
// 		default:
// 			{
// 				return resp.SerializeError("unexpected option")
// 			}
// 		}
// 	}

// 	return resp.SerializeSimpleString("OK")
// }
