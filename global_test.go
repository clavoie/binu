package binu_test

import (
	"bytes"
	"encoding/gob"
	"reflect"
	"testing"

	"github.com/clavoie/binu"
)

func TestGlobals(t *testing.T) {
	type Foo struct {
		Value1 int
		Value2 float32
		Value3 string
	}

	intVal := 10
	var floatVal float32 = 20.0
	stringVal := "30"
	baseFoo := &Foo{intVal, floatVal, stringVal}

	buffer := new(bytes.Buffer)
	encoder := gob.NewEncoder(buffer)
	err := encoder.Encode(baseFoo)

	if err != nil {
		t.Fatal(err)
	}

	expectedBytes := buffer.Bytes()

	t.Run("FromGob", func(t *testing.T) {
		newFoo := new(Foo)
		err := binu.FromGob(expectedBytes, newFoo)

		if err != nil {
			t.Fatal(err)
		}

		if reflect.DeepEqual(baseFoo, newFoo) == false {
			t.Fatal(baseFoo, newFoo)
		}
	})
	t.Run("ToGob", func(t *testing.T) {
		bs, err := binu.ToGob(baseFoo)

		if err != nil {
			t.Fatal(err)
		}

		if bytes.Equal(expectedBytes, bs) == false {
			t.Fatal("Byte slices do not match")
		}
	})
}
