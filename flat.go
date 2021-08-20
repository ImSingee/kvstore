package kvstore

import "strings"

func FlatForHighLevel(key string) []string {
	if key == "" {
		return nil
	}

	parts := make([]string, 0, 8)

	findAndDo(key, '.', func(i int) {
		parts = append(parts, key[:i])
	})

	parts = append(parts, key)

	return parts
}

func findAndDo(s string, c byte, f func(int)) {
	skip := 0
	sub := s

	for {
		next := strings.IndexByte(sub, c)
		if next == -1 {
			return
		}
		f(skip + next)

		skip += next + 1
		sub = s[skip:]
	}
}
