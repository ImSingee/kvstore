package manager

import "github.com/ImSingee/kvstore"

func (m *manager) Delete(key string) error {
	return m.writeLogAndDoE(kvstore.NewDeleteAction(key), func(s kvstore.UnsafeStore) error {
		return s.UnsafeDelete(key)
	})
}
