package kvstore

import (
	"github.com/ImSingee/tt"
	"testing"
)

func TestStore_Delete(t *testing.T) {
	t.Run("simple-delete", func(t *testing.T) {
		store := newStore()
		base := store.Export()
		tt.AssertIsNil(t, store.Delete("children.a"))
		tt.AssertNotEqual(t, base, store.Export())
		delete(base["children"].(map[string]interface{}), "a")
		tt.AssertEqual(t, base, store.Export())
	})

	t.Run("delete-not-exist", func(t *testing.T) {
		store := newStore()
		base := store.Export()
		tt.AssertIsNil(t, store.Delete("not-exist"))
		tt.AssertEqual(t, base, store.Export())
	})

	t.Run("delete-middle-not-exist", func(t *testing.T) {
		store := newStore()
		base := store.Export()
		tt.AssertIsNotNil(t, store.Delete("children.c.name"))
		tt.AssertEqual(t, base, store.Export())
	})

	t.Run("delete-type-not-match", func(t *testing.T) {
		store := newStore()
		base := store.Export()
		tt.AssertIsNotNil(t, store.Delete("isActive.k"))
		tt.AssertEqual(t, base, store.Export())
	})

	t.Run("delete-in-list", func(t *testing.T) {
		store := newStore()
		base := store.Export()
		tt.AssertIsNotNil(t, store.Delete("tags.3"))
		tt.AssertEqual(t, base, store.Export())
	})
}
