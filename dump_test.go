package kvstore

import (
	"github.com/ImSingee/tt"
	"testing"
)

func TestDump(t *testing.T) {
	store := newStore()
	b, err := store.Dump()
	tt.AssertIsNil(t, err)
	newStore, err := Load(b)
	tt.AssertIsNil(t, err)

	tt.AssertEqual(t, store.Export(), newStore.Export())
}
