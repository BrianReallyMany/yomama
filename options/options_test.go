package options

import (
	"testing"
)

type Foo struct {
	Dog string
}

type testStruct struct {
	Num int
	Str string
	F Foo
}

func TestConfigMapConstruction(t *testing.T) {
	test := testStruct{42, "Hello, world!", Foo{"yo man"}}

	NewOptions(&test)
}
