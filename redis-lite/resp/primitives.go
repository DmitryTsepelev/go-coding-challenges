package resp

const SimpleString = '+'
const Error = '-'
const Integer = ':'
const BulkString = '$'
const Array = '*'

type IntegerValue struct {
	Value *int
}

type SimpleStringValue struct {
	String *string
}

type ErrorValue struct {
	Error *string
}

type BulkStringValue struct {
	String *string
}

type ArrayValue struct {
	Array []interface{}
}
