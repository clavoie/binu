package binu

import (
	"bytes"
	"encoding/gob"
)

// Gobber serializes interface{} values to and from
// []byte values using the encoding/gob package.
type Gobber interface {
	// From takes a byte slice encoded with the gob package and
	// deserializes it into dst, returning any error encountered
	// during the process.
	From(data []byte, dst interface{}) error

	// To takes the value provided and encodes it using the gob package.
	// The encoded value and any error encountered are returned.
	To(interface{}) ([]byte, error)
}

// gobber is an implementation of Gobber
type gobber struct{}

// NewGobber returns a new instance of a Gobber
func NewGobber() Gobber {
	return new(gobber)
}

func (g *gobber) From(data []byte, dst interface{}) error {
	buffer := bytes.NewBuffer(data)
	decoder := gob.NewDecoder(buffer)
	return decoder.Decode(dst)
}

func (g *gobber) To(val interface{}) ([]byte, error) {
	buffer := new(bytes.Buffer)
	encoder := gob.NewEncoder(buffer)
	err := encoder.Encode(val)

	if err == nil {
		return buffer.Bytes(), nil
	}

	return nil, err
}
