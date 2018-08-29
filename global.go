package binu

// FromGob takes a byte slice encoded with the gob package and
// deserializes it into dst, returning any error encountered
// during the process.
func FromGob(data []byte, dst interface{}) error {
	return NewGobber().From(data, dst)
}

// ToGob takes the value provided and encodes it using the gob package.
// The encoded value and any error encountered are returned.
func ToGob(src interface{}) ([]byte, error) {
	return NewGobber().To(src)
}
