package manager

import "github.com/ImSingee/kvstore"

func (m *manager) Clear() error {
	return m.writeLogAndDo(kvstore.NewClearAction(), func(s kvstore.UnsafeStore) {
		s.UnsafeClear()
	})
}
