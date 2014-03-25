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

func TestOptionsRead(t *testing.T) {
	test := TestStruct{0, "", Foo{""}}

	options := NewOptions(&test)

	buffer := bytes.NewBufferString("Num =  42\nStr =Hello, world!\nDog=yo man\n")
	options.Read(buffer)

	if test.Num != 42 {
		t.Log("test.Num = ", test.Num)
		t.Fail()
	}
	if test.Str != "Hello, world!" {
		t.Log("test.Str = ", test.Str)
		t.Fail()
	}
	if test.F.Dog != "yo man" {
		t.Log("test.F.Dog = ", test.F.Dog)
		t.Fail()
	}
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
