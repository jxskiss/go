package struct_of_set_test

type TestObject struct {
	Field1 map[int32]bool `thrift:",1,,set"`
}