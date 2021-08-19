package kvstore

import (
	"github.com/ImSingee/tt"
	"testing"
)

func TestTrie(t *testing.T) {
	tree := NewTrie([]string{
		"a.b.c",
		"a.c.d",
		"b.c",
	})
	tt.AssertTrue(t, tree.CheckExist("a.b.c"))
	tt.AssertTrue(t, tree.CheckExist("a.b.c.d"))
	tt.AssertTrue(t, tree.CheckExist("a.c.d"))
	tt.AssertTrue(t, tree.CheckExist("b.c"))
	tt.AssertTrue(t, tree.CheckExist("b.c.e"))
	tt.AssertTrue(t, tree.CheckExist("b.c.e.f"))

	tt.AssertFalse(t, tree.CheckExist("a"))
	tt.AssertFalse(t, tree.CheckExist("a.b"))
	tt.AssertFalse(t, tree.CheckExist("a.c"))
	tt.AssertFalse(t, tree.CheckExist("b"))
	tt.AssertFalse(t, tree.CheckExist("c"))

}
