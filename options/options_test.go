package options

import (
	"bytes"
	"testing"
)

type Foo struct {
	Dog string
}

type TestStruct struct {
	Num int
	Str string
	F Foo
}

func TestOptionsWrite(t *testing.T) {
	test := TestStruct{42, "Hello, world!", Foo{"yo man"}}

	options := NewOptions(&test)

	buffer := bytes.NewBufferString("")
	options.Write(buffer)

	if buffer.String() != "Num=42\nStr=Hello, world!\nDog=yo man\n" {
		t.Fail()
	}
}
