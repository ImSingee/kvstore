package kvstore

import (
	"github.com/ImSingee/tt"
	"google.golang.org/protobuf/proto"
	"testing"
)

func TestStore_ApplyByJSON_stable(t *testing.T) {
	base := newStore()
	cmp := NewStore()
	err := cmp.ApplyByJSON([]byte(testData))
	tt.AssertIsNil(t, err)

	//baseJ, _ := json.Marshal(base.Export())
	//cmpJ, _ := json.Marshal(cmp.Export())

	tt.AssertEqual(t, base.Export(), cmp.Export())
	tt.AssertEqual(t, base.ExportFlat(), cmp.ExportFlat())
	tt.AssertTrue(t, proto.Equal(base.Unsafe().UnsafeUnderlying(), cmp.Unsafe().UnsafeUnderlying()))
}

func TestStore_ApplyByJSON(t *testing.T) {
	base := NewStore()
	err := base.ApplyByJSON([]byte(`{
    "a":10,
    "b.c.d":31,
    "b.c.e":32,
    "c":[
        "a",
        "b",
        "c"
    ],
    "f":{
        "m":3,
        "p": 1.2,
        "q":true
    }
}`))
	tt.AssertIsNil(t, err)

	cmp := map[string]interface{}{
		"a": int64(10),
		"b": map[string]interface{}{
			"c": map[string]interface{}{
				"d": int64(31),
				"e": int64(32),
			},
		},
		"c": []interface{}{"a", "b", "c"},
		"f": map[string]interface{}{
			"m": int64(3),
			"p": float64(1.2),
			"q": true,
		},
	}

	tt.AssertEqual(t, cmp, base.Export())
}
