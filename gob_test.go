package binu_test

import (
	"bytes"
	"encoding/gob"
	"reflect"
	"testing"

	"github.com/clavoie/binu"
)

func TestGobber(t *testing.T) {
	type Foo struct {
		Value1 int
		Value2 float32
		Value3 string
	}

	intVal := 10
	var floatVal float32 = 20.0
	stringVal := "30"
	baseFoo := &Foo{intVal, floatVal, stringVal}
	gobber := binu.NewGobber()

	buffer := new(bytes.Buffer)
	encoder := gob.NewEncoder(buffer)
	err := encoder.Encode(baseFoo)

	if err != nil {
		t.Fatal(err)
	}

	expectedBytes := buffer.Bytes()

	t.Run("FromSuccess", func(t *testing.T) {
		newFoo := new(Foo)
		err := gobber.From(expectedBytes, newFoo)

		if err != nil {
			t.Fatal(err)
		}

		if reflect.DeepEqual(baseFoo, newFoo) == false {
			t.Fatal(baseFoo, newFoo)
		}
	})
	t.Run("FromFail", func(t *testing.T) {
		type Fail struct {
			X string
		}
		newFail := new(Fail)
		err := gobber.From(expectedBytes, newFail)

		if err == nil {
			t.Fatal("Expected failure")
		}
	})

	t.Run("ToSuccess", func(t *testing.T) {
		bs, err := gobber.To(baseFoo)

		if err != nil {
			t.Fatal(err)
		}

		if bytes.Equal(expectedBytes, bs) == false {
			t.Fatal("Byte slices do not match")
		}
	})
	t.Run("ToFail", func(t *testing.T) {
		_, err := gobber.To(nil)

		if err == nil {
			t.Fatal("Expecting err")
		}
	})
}
