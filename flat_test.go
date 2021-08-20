package kvstore

import (
	"github.com/ImSingee/tt"
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
