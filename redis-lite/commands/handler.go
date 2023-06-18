package commands

import (
	"net"
	"redis-lite/instance"
	"redis-lite/resp"
)

const PING = "PING"
const ECHO = "ECHO"
const SET = "SET"
const GET = "GET"
const TTL = "TTL"
const EXPIRE = "EXPIRE"
const DEL = "DEL"
const INCR = "INCR"
const DECR = "DECR"

var commandRoutes = map[string](func(*instance.Instance, []string) (*string, error)){
	PING:   ping,
	ECHO:   echo,
	SET:    set,
	GET:    get,
	TTL:    ttl,
	EXPIRE: expire,
	DEL:    del,
	INCR:   incr,
	DECR:   decr,
}

func HandleCommand(inst *instance.Instance, command string, args []string, conn net.Conn) {
	handler := commandRoutes[command]
	if handler == nil {
		handler = func(instance *instance.Instance, args []string) (*string, error) {
			return resp.SerializeError("Unknown command: " + command)
		}
	}

	response, err := handler(inst, args)
	if err == nil {
		conn.Write([]byte(*response))
	} else {
		errorString, _ := resp.SerializeSimpleString("Error: " + err.Error())
		conn.Write([]byte(*errorString))
	}
}
