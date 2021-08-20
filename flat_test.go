package kvstore

import (
	"github.com/ImSingee/tt"
	"sort"
	"testing"
)

func TestFlatForHighLevel(t *testing.T) {
	cases := []struct {
		name string
		in   string
		want []string
	}{
		{
			name: `empty`,
			in:   "",
			want: nil,
		},
		{
			name: `single`,
			in:   "a",
			want: []string{"a"},
		},
		{
			name: `multipart`,
			in:   "a.b.c",
			want: []string{"a", "a.b", "a.b.c"},
		},
		{
			name: `trailing-dot`,
			in:   "a.b.c.",
			want: []string{"a", "a.b", "a.b.c", "a.b.c."},
		},
		{
			name: `prefix-dot`,
			in:   ".a.b.c",
			want: []string{"", ".a", ".a.b", ".a.b.c"},
		},
		{
			name: `joint-dot`,
			in:   "a..b.c",
			want: []string{"a", "a.", "a..b", "a..b.c"},
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			got := FlatForHighLevel(c.in)
			tt.AssertEqual(t, got, c.want)
		})
	}
}

func TestStore_FlatForLowLevel(t *testing.T) {
	store := newStore()

	cases := []struct {
		name string
		in   string
		want []string
	}{
		{
			name: `not-exist`,
			in:   "xx",
			want: nil,
		},
		{
			name: `single`,
			in:   "greeting",
			want: []string{"greeting"},
		},
		{
			name: `first-level-list`,
			in:   "friends",
			want: []string{
				"friends",
				"friends.0",
				"friends.0.id",
				"friends.0.name",
				"friends.1",
				"friends.1.id",
				"friends.1.name",
				"friends.2",
				"friends.2.id",
				"friends.2.name",
			},
		},
		{
			name: `first-level-map`,
			in:   "simple",
			want: []string{
				"simple",
				"simple.eyeColor",
				"simple.name",
				"simple.gender",
				"simple.company",
				"simple.email",
				"simple.phone",
				"simple.address",
				"simple.about",
				"simple.registered",
				"simple.latitude",
				"simple.longitude",
			},
		},
		{
			name: `complicate`,
			in:   "complicate",
			want: []string{
				"complicate",
				"complicate.name",
				"complicate.gender",
				"complicate.company",
				"complicate.tags",
				"complicate.tags.0",
				"complicate.tags.1",
				"complicate.tags.2",
				"complicate.tags.3",
				"complicate.tags.3.id",
				"complicate.tags.3.name",
				"complicate.tags.4",
				"complicate.tags.4.isActive",
				"complicate.tags.4.balance",
				"complicate.tags.4.oops",
				"complicate.tags.4.oops.0",
				"complicate.tags.4.oops.1",
				"complicate.tags.5",
			},
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			got := store.FlatForLowLevel(c.in)
			sort.Strings(got)
			sort.Strings(c.want)
			tt.AssertEqual(t, c.want, got)
		})
	}
}
