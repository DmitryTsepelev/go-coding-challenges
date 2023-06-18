package resp

import "fmt"

func SerializeSimpleString(simpleStr string) (*string, error) {
	result := string(SimpleString) + simpleStr + "\r\n"
	return &result, nil
}

func SerializeBulkString(bulkStr *string) (*string, error) {
	var result string

	if bulkStr == nil {
		result = string(BulkString) + "-1\r\n"
	} else {
		result = string(BulkString) + fmt.Sprint(len(*bulkStr)) + "\r\n" + *bulkStr + "\r\n"
	}

	return &result, nil
}

func SerializeError(errStr string) (*string, error) {
	result := string(Error) + errStr + "\r\n"
	return &result, nil
}
