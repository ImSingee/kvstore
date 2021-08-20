package kvstore

import (
	"github.com/ImSingee/tt"
	"testing"
)

func TestStore_Set(t *testing.T) {
	s := NewStore()
	tt.AssertIsNil(t, s.SetInt64("a", 10))
	tt.AssertIsNil(t, s.SetInt64("b.c.d", 31))
	tt.AssertIsNil(t, s.SetInt64("b.c.e", 32))
	tt.AssertIsNotNil(t, s.SetInt64("a.c", 3)) // a 类型不匹配失败
	tt.AssertIsNil(t, s.SetStringList("c", []string{"a", "b", "c"}))
	tt.AssertIsNil(t, s.SetString("c.1", "bb"))
	tt.AssertIsNotNil(t, s.SetString("c.3", "dd"))
	tt.AssertIsNil(t, s.SetInt64("c.+", 8))

	cc := map[string]interface{}{
		"a": int64(10),
		"b": map[string]interface{}{
			"c": map[string]interface{}{
				"d": int64(31),
				"e": int64(32),
			},
		},
		"c": []interface{}{"a", "bb", "c", int64(8)},
	}

	tt.AssertEqual(t, cc, s.Export())
}
