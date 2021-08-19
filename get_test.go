package kvstore

import (
	"fmt"
	"testing"

	"github.com/ImSingee/tt"
)

var testData = `
{
  "_id": "611e0ba6551aff2f434298d1",
  "index": 0,
  "guid": "8674f065-3ee8-4b40-a699-2397eddf0012",
  "isActive": false,
  "balance": "$1,844.03",
  "picture": "http://placehold.it/32x32",
  "age": 37,
  "eyeColor": "blue",
  "name": "Patton Nieves",
  "gender": "male",
  "company": "XYQAG",
  "email": "pattonnieves@xyqag.com",
  "phone": "+1 (837) 456-3131",
  "address": "172 Claver Place, Juntura, Alaska, 5942",
  "about": "Anim do nulla Lorem velit. Dolore excepteur non dolor Lorem amet ut occaecat irure. Duis velit sunt in sunt ex ex ut cupidatat aliquip in eu cupidatat elit occaecat. Labore ad culpa do id irure duis esse cupidatat velit labore ullamco non consectetur. Ipsum nulla adipisicing cupidatat tempor nostrud non. Quis voluptate eu anim occaecat.\r\n",
  "registered": "2018-11-14T08:30:10 -08:00",
  "latitude": 11.655197,
  "longitude": 37.981309,
  "tags": [
    "ad",
    "anim",
    "quis",
    "ex",
    "velit",
    "voluptate",
    "et"
  ],
  "friends": [
    {
      "id": 0,
      "name": "Molly Castaneda"
    },
    {
      "id": 1,
      "name": "Colleen Jackson"
    },
    {
      "id": 2,
      "name": "Dunlap Wood"
    }
  ],
  "greeting": "Hello, Patton Nieves! You have 5 unread messages.",
  "favoriteFruit": "banana"
}
`

func newStore() *store {
	s, err := LoadFromJSON([]byte(testData))
	if err != nil {
		panic(err)
	}

	return s.(*store)
}

func TestStore_GetInt(t *testing.T) {
	cases := []struct {
		name      string
		in        string
		want      int64
		willError bool
	}{
		{
			name:      `not-exist`,
			in:        "xx",
			willError: true,
		},
		{
			name: `first-level`,
			in:   "age",
			want: 37,
		},
		{
			name: `nested`,
			in:   "friends.1.id",
			want: 1,
		},
		{
			name:      `array-index-invalid`,
			in:        "friends.-1.id",
			willError: true,
		},
		{
			name:      `array-index-overflow`,
			in:        "friends.99.id",
			willError: true,
		},
		{
			name:      `type-mismatch`,
			in:        "greeting",
			willError: true,
		},
		{
			name:      `middle-type-mismatch`,
			in:        "friends.a.xx",
			willError: true,
		},
		{
			name:      `attr-access-on-basic`,
			in:        "greeting.something",
			willError: true,
		},
	}

	store := newStore()

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			got, err := store.GetInt(c.in)

			if !c.willError {
				tt.AssertIsNil(t, err)

				tt.AssertEqual(t, got, c.want)
			} else {
				tt.AssertIsNotNil(t, err)
				fmt.Println("Expected error:", err)
			}
		})
	}
}
