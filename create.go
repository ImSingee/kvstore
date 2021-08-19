package kvstore

func NewStore() Store {
	return &store{Provider: NewEmptyMap()}
}

func CreateStore(data map[string]interface{}) (Store, error) {
	if data == nil {
		return NewStore(), nil
	}

	m, err := NewMap(data)
	if err != nil {
		return nil, err
	}

	return &store{Provider: m}, nil
}

func LoadFromJSON(data []byte) (Store, error) {
	m := NewEmptyMap()

	err := m.UnmarshalJSON(data)
	if err != nil {
		return nil, err
	}

	return &store{Provider: m}, nil
}
