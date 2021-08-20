package kvstore

import (
	"github.com/ImSingee/tt"
	"testing"
)

func TestFilterDeep(t *testing.T) {
	s := newStore()

	t.Run("nil-filter", func(t *testing.T) {
		v := s.ExportAndFilter(nil)
		tt.AssertIsNotNil(t, v)
		tt.AssertEqual(t, 0, len(v))
	})

	t.Run("empty-filter", func(t *testing.T) {
		v := s.ExportAndFilter([]string{})
		tt.AssertIsNotNil(t, v)
		tt.AssertEqual(t, 0, len(v))
	})

	t.Run("single", func(t *testing.T) {
		v := s.ExportAndFilter([]string{"balance"})
		tt.AssertIsNotNil(t, v)
		tt.AssertEqual(t, 1, len(v))
		tt.AssertEqual(t, "$1,844.03", v["balance"])
	})

	t.Run("single-and-extra", func(t *testing.T) {
		v := s.ExportAndFilter([]string{"balance", "something", "something.not.exist"})
		tt.AssertIsNotNil(t, v)
		tt.AssertEqual(t, 1, len(v))
		tt.AssertEqual(t, "$1,844.03", v["balance"])
	})

	t.Run("first-level-map", func(t *testing.T) {
		v := s.ExportAndFilter([]string{"children"})
		tt.AssertIsNotNil(t, v)
		tt.AssertEqual(t, 1, len(v))
		children, _ := v["children"].(map[string]interface{})
		tt.AssertIsNotNil(t, children)
		tt.AssertEqual(t, 2, len(children))
		tt.AssertEqual(t, s.Export()["children"], children)
	})

	t.Run("second-level-map", func(t *testing.T) {
		v := s.ExportAndFilter([]string{"children.a"})
		tt.AssertIsNotNil(t, v)
		tt.AssertEqual(t, 1, len(v))
		children, _ := v["children"].(map[string]interface{})
		tt.AssertIsNotNil(t, children)
		tt.AssertEqual(t, 1, len(children))
		expect, _ := s.GetMap("children.a")
		tt.AssertIsNotNil(t, expect)
		tt.AssertEqual(t, expect, children["a"])
	})
}
