# binu [![GoDoc](https://godoc.org/github.com/clavoie/erru?status.svg)](http://godoc.org/github.com/clavoie/erru) [![Build Status](https://travis-ci.org/clavoie/erru.svg?branch=master)](https://travis-ci.org/clavoie/erru) [![codecov](https://codecov.io/gh/clavoie/erru/branch/master/graph/badge.svg)](https://codecov.io/gh/clavoie/erru) [![Go Report Card](https://goreportcard.com/badge/github.com/clavoie/erru)](https://goreportcard.com/report/github.com/clavoie/erru)

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
