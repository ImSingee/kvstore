package manager

import "github.com/ImSingee/kvstore"

func (m *manager) ApplyByJSON(data []byte) error {
	actions, err := kvstore.JSONToActions(data)
	if err != nil {
		return err
	}

	return m.writeLogsAndDoE(actions, func(s kvstore.UnsafeStore) error {
		return s.UnsafeApply(actions)
	})
}
