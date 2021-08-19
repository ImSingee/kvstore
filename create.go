package kvstore

func NewStore() Store {
	return &store{Provider: NewEmptyMap()}
}
