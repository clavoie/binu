# binu [![GoDoc](https://godoc.org/github.com/clavoie/binu?status.svg)](http://godoc.org/github.com/clavoie/binu) [![Build Status](https://travis-ci.org/clavoie/binu.svg?branch=master)](https://travis-ci.org/clavoie/binu) [![Go Report Card](https://goreportcard.com/badge/github.com/clavoie/binu)](https://goreportcard.com/report/github.com/clavoie/binu)

Binary / byte utilities for Go.

Currently this package is just a wrapper around the encoding/gob package, with hooks
for dependency injection.

```go
type Foo struct {
  Value int
}

foo := &Foo{10}
bytes, err := binu.ToGob(foo)

// if err etc

newFoo := new(Foo)
err = binu.FromGob(bytes, newFoo)

// if err etc
// newFoo.Value == foo.Value
```

If you'd prefer to use dependency injection:

```go
func Serialize(gobber binu.Gobber) error {
  err := gobber.To(value)

  // etc
}
```
