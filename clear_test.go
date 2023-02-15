package kvstore

import (
	"github.com/ImSingee/tt"
	"testing"
)

func TestStore_Clear(t *testing.T) {
	t.Run("simple-clear", func(t *testing.T) {
		store := newStore()
		tt.AssertNotEqual(t, map[string]interface{}{}, store.Export())
		store.Clear()
		tt.AssertEqual(t, map[string]interface{}{}, store.Export())
	})
}
