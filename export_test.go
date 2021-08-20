package kvstore

import (
	"github.com/ImSingee/tt"
	"testing"
)

func newComplicateStore() Store {
	data, err := newStore().GetMap("complicate")
	if err != nil {
		panic(err)
	}
	s, err := CreateStore(data)
	if err != nil {
		panic(err)
	}

	return s
}

func TestExport(t *testing.T) {
	store := newComplicateStore()

	cmp := map[string]interface{}{
		"name":    "Hannah Mack",
		"gender":  "female",
		"company": "VERTON",
		"tags": []interface{}{
			"duis",
			"irure",
			"non",
			map[string]interface{}{
				"id":   int64(0),
				"name": "Naomi Burris",
			},
			map[string]interface{}{
				"isActive": false,
				"balance":  "$2,422.40",
				"oops": []interface{}{
					"duis",
					"irure",
				},
			},
			"id",
		},
	}

	tt.AssertEqual(t, cmp, store.Export())
}

func TestExportFlat(t *testing.T) {
	store := newComplicateStore()

	cmp := map[string]interface{}{
		"name":            "Hannah Mack",
		"gender":          "female",
		"company":         "VERTON",
		"tags.0":          "duis",
		"tags.1":          "irure",
		"tags.2":          "non",
		"tags.3.id":       int64(0),
		"tags.3.name":     "Naomi Burris",
		"tags.4.isActive": false,
		"tags.4.balance":  "$2,422.40",
		"tags.4.oops.0":   "duis",
		"tags.4.oops.1":   "irure",
		"tags.5":          "id",
	}

	tt.AssertEqual(t, cmp, store.ExportFlat())
}

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
