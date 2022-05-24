package kvstore

func MapToActions(m map[string]interface{}) ([]*Action, error) {
	store, err := CreateStore(m)
	if err != nil {
		return nil, err
	}
	u := store.Unsafe().UnsafeUnderlying().Fields
	actions := make([]*Action, 0, len(u))

	for k, v := range u {
		actions = append(actions, NewSetAction(k, v))
	}

	return actions, nil
}

func (s *store) ApplyByMap(data map[string]interface{}) error {
	actions, err := MapToActions(data)
	if err != nil {
		return err
	}
	return s.Apply(actions) // Apply 内部会上锁
}
