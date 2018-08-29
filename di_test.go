package binu_test

import (
	"testing"

	"github.com/clavoie/binu"
)

func TestNewDiDefs(t *testing.T) {
	defs := binu.NewDiDefs()

	if defs == nil {
		t.Fatal("Expecting non-nil defs")
	}

	defs2 := binu.NewDiDefs()
	if defs[0] == defs2[0] {
		t.Fatal("Not expecting defs to match")
	}
}
