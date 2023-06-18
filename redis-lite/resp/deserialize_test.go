package resp

import (
	"reflect"
	"testing"
)

// TODO: test rest

func TestDeserializeSimpleString(t *testing.T) {
	input := "+OK\r\n"
	result, _, err := Deserialize(input)

	if str, ok := result.(SimpleStringValue); !ok || (*str.String != "OK") || (err != nil) {
		t.Fatalf(`got %s, err %d`, *str.String, err)
	}
}

func TestDeserializeError(t *testing.T) {
	input := "-Wrong\r\n"
	result, _, err := Deserialize(input)

	if str, ok := result.(ErrorValue); !ok || (*str.Error != "Wrong") || (err != nil) {
		t.Fatalf(`got %s, err %d`, *str.Error, err)
	}
}

func TestDeserializeInt(t *testing.T) {
	input := ":42\r\n"
	result, _, err := Deserialize(input)

	if intVal, ok := result.(IntegerValue); !ok || (*(intVal.Value) != 42) || (err != nil) {
		t.Fatalf(`got %d, err %d`, intVal.Value, err)
	}
}

func TestDeserializeWrongInt(t *testing.T) {
	input := ":qwe\r\n"
	_, _, err := Deserialize(input)

	if err.Error() != "strconv.Atoi: parsing \"qwe\": invalid syntax" {
		t.Fatalf(`expected error`)
	}
}

func TestDeserializeBulkString(t *testing.T) {
	input := "$5\r\nasdfg\r\n"
	result, _, err := Deserialize(input)

	if str, ok := result.(BulkStringValue); !ok || (*str.String != "asdfg") || (err != nil) {
		t.Fatalf(`got %s, err %d`, *str.String, err)
	}
}

func TestDeserializeBulkStringEmpty(t *testing.T) {
	input := "$0\r\n\r\n"
	result, _, err := Deserialize(input)

	if str, ok := result.(BulkStringValue); !ok || (*str.String != "") || (err != nil) {
		t.Fatalf(`got %s, err %d`, *str.String, err)
	}
}

func TestDeserializeBulkStringNull(t *testing.T) {
	input := "$-1\r\n"
	result, _, err := Deserialize(input)

	if str, ok := result.(BulkStringValue); !ok || (str.String != nil) || (err != nil) {
		t.Fatalf(`got %s, err %d`, *str.String, err)
	}
}

func TestDeserializeArrayEmpty(t *testing.T) {
	input := "*0\r\n"
	result, _, err := Deserialize(input)

	if value, ok := result.(ArrayValue); !ok || (!reflect.DeepEqual(value.Array, make([]interface{}, 0))) || (err != nil) {
		t.Fatalf(`got %s, err %d`, value.Array, err)
	}
}

func TestDeserializeArrayWithInt(t *testing.T) {
	input := "*1\r\n:42\r\n"
	result, _, err := Deserialize(input)

	i42 := 42
	expected := ArrayValue{
		Array: []interface{}{
			IntegerValue{Value: &i42},
		}}

	if value, ok := result.(ArrayValue); !ok || (!reflect.DeepEqual(value, expected)) || (err != nil) {
		t.Fatalf(`got %s, err %d`, value.Array, err)
	}
}

func TestDeserializeArrayWithTwoStrings(t *testing.T) {
	input := "*2\r\n$5\r\nhello\r\n$5\r\nworld\r\n"
	result, _, err := Deserialize(input)

	hello := "hello"
	world := "world"
	expected := ArrayValue{
		Array: []interface{}{
			BulkStringValue{String: &hello},
			BulkStringValue{String: &world},
		}}

	if value, ok := result.(ArrayValue); !ok || (!reflect.DeepEqual(value, expected)) || (err != nil) {
		t.Fatalf(`got %s, err %d`, value.Array, err)
	}
}

func TestDeserializeArrayWithThreeIntegers(t *testing.T) {
	input := "*3\r\n:1\r\n:2\r\n:3\r\n"
	result, _, err := Deserialize(input)

	i1 := 1
	i2 := 2
	i3 := 3
	expected := ArrayValue{
		Array: []interface{}{
			IntegerValue{Value: &i1},
			IntegerValue{Value: &i2},
			IntegerValue{Value: &i3},
		}}

	if value, ok := result.(ArrayValue); !ok || (!reflect.DeepEqual(value, expected)) || (err != nil) {
		t.Fatalf(`got %s, err %d`, value.Array, err)
	}
}

func TestDeserializeArrayWithMixedTypes(t *testing.T) {
	input := "*3\r\n:1\r\n:2\r\n$5\r\nhello\r\n"
	result, _, err := Deserialize(input)

	i1 := 1
	i2 := 2
	hello := "hello"

	expected := ArrayValue{
		Array: []interface{}{
			IntegerValue{Value: &i1},
			IntegerValue{Value: &i2},
			BulkStringValue{String: &hello},
		}}

	if value, ok := result.(ArrayValue); !ok || (!reflect.DeepEqual(value, expected)) || (err != nil) {
		t.Fatalf(`got %s, err %d`, value.Array, err)
	}
}

func TestDeserializeArrayWithNestedArrays(t *testing.T) {
	input := "*2\r\n*3\r\n:1\r\n:2\r\n:3\r\n*2\r\n+Hello\r\n-World\r\n"
	result, _, err := Deserialize(input)

	i1 := 1
	i2 := 2
	i3 := 3
	array1 := ArrayValue{
		Array: []interface{}{
			IntegerValue{Value: &i1},
			IntegerValue{Value: &i2},
			IntegerValue{Value: &i3},
		},
	}

	hello := "Hello"
	world := "World"
	array2 := ArrayValue{
		Array: []interface{}{
			SimpleStringValue{String: &hello},
			ErrorValue{Error: &world},
		}}

	expected := ArrayValue{
		Array: []interface{}{array1, array2},
	}

	if value, ok := result.(ArrayValue); !ok || (!reflect.DeepEqual(value, expected)) || (err != nil) {
		t.Fatalf(`got %s, err %d`, value.Array, err)
	}
}

func TestDeserializeArrayNull(t *testing.T) {
	input := "*-1\r\n"
	result, _, err := Deserialize(input)

	expected := ArrayValue{
		Array: nil,
	}

	if value, ok := result.(ArrayValue); !ok || (!reflect.DeepEqual(value, expected)) || (err != nil) {
		t.Fatalf(`got %s, err %d`, value.Array, err)
	}
}

func TestDeserializeArrayWithNull(t *testing.T) {
	input := "*3\r\n$5\r\nhello\r\n$-1\r\n$5\r\nworld\r\n"
	result, _, err := Deserialize(input)

	hello := "hello"
	world := "world"
	expected := ArrayValue{
		Array: []interface{}{
			BulkStringValue{String: &hello},
			BulkStringValue{String: nil},
			BulkStringValue{String: &world},
		}}

	if value, ok := result.(ArrayValue); !ok || (!reflect.DeepEqual(value, expected)) || (err != nil) {
		t.Fatalf(`got %s, err %d`, value.Array, err)
	}
}
