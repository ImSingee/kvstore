package manager

func (m *manager) IsReadonly() bool {
	return m.ro
}
