package resp

import (
	"fmt"
	"strconv"
)

func parseUntilCrlf(data *string) (*string, *string) {
	var result string
	var length int

	for idx, currentRune := range *data {

		if currentRune == '\r' && (*data)[idx+1] == '\n' {
			break
		}

		result += string(currentRune)
		length = idx
	}

	start := length + 3
	if start >= len(*data) {
		start = len(*data) - 1
	}

	rest := (*data)[start:]
	return &result, &rest
}

func Deserialize(input string) (interface{}, *string, error) {
	// TODO: check input length
	dataTypeByte := input[0]
	dataBytes := input[1:]

	switch dataTypeByte {
	case SimpleString:
		str, rest := parseUntilCrlf(&dataBytes)

		return SimpleStringValue{String: str}, rest, nil
	case Error:
		err, rest := parseUntilCrlf(&dataBytes)

		return ErrorValue{Error: err}, rest, nil
	case Integer:
		result, rest := parseUntilCrlf(&dataBytes)

		value, err := strconv.Atoi(*result)
		if err != nil {
			return nil, nil, err
		}

		return IntegerValue{Value: &value}, rest, nil
	case BulkString:
		result, rest := parseUntilCrlf(&dataBytes)

		length, err := strconv.Atoi(*result)
		if err != nil {
			return nil, nil, err
		}

		if length == -1 {
			return BulkStringValue{String: nil}, rest, nil
		}

		content := (*rest)[0:length]

		rest2 := ((*rest)[length+2:])

		return BulkStringValue{String: &content}, &rest2, nil
	case Array:
		result, rest := parseUntilCrlf(&dataBytes)

		length, err := strconv.Atoi(*result)
		if err != nil {
			return nil, nil, err
		}

		if length == -1 {
			return ArrayValue{Array: nil}, rest, nil
		}

		array := make([]interface{}, 0)

		for i := 0; i < length; i++ {
			element, newRest, deserializeErr := Deserialize(*rest)
			// TODO: check element is not array

			if deserializeErr != nil {
				return nil, newRest, deserializeErr
			}
			array = append(array, element)
			rest = newRest
		}

		return ArrayValue{Array: array}, rest, nil
	default:
		panic("unsupported data type byte")
	}
}

func DeserializeArrayOfBulkStrings(input string) (*[]string, error) {
	arrayWrapped, _, err := Deserialize(input)

	if err != nil {
		return nil, err
	}

	if array, ok := arrayWrapped.(ArrayValue); ok {
		bulkStrings := []string{}

		for _, stringWrapped := range array.Array {
			if str, ok := stringWrapped.(BulkStringValue); ok {
				bulkStrings = append(bulkStrings, *str.String)
			} else {
				return nil, fmt.Errorf("failed to deserialize array")
			}
		}

		return &bulkStrings, nil
	}

	return nil, fmt.Errorf("failed to deserialize array")
}
