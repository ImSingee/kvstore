package manager

func (m *manager) Close() {
	m.lock()
	// won't unlock

	m.appendFile.Close()
}
