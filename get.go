package kvstore

func (s *store) Get(key string) (AnyValue, error) {
	s.lock.RLock()
	defer s.lock.RUnlock()

	return s.readValue(key)
}

func (s *store) GetInt(key string) (int64, error) {
	s.lock.RLock()
	defer s.lock.RUnlock()

	v, err := s.readValue(key)
	if err != nil {
		return 0, err
	}

	if vv, ok := v.(int64); ok {
		return vv, nil
	}

	return 0, ErrKeyTypeNotMatch{
		Key:    key,
		Expect: "int",
		Got:    TypeName(v),
	}
}

func (s *store) GetFloat(key string) (float64, error) {
	s.lock.RLock()
	defer s.lock.RUnlock()

	v, err := s.readValue(key)
	if err != nil {
		return 0, err
	}

	if vv, ok := v.(float64); ok {
		return vv, nil
	}

	return 0, ErrKeyTypeNotMatch{
		Key:    key,
		Expect: "float",
		Got:    TypeName(v),
	}
}

func (s *store) GetBool(key string) (bool, error) {
	s.lock.RLock()
	defer s.lock.RUnlock()

	v, err := s.readValue(key)
	if err != nil {
		return false, err
	}

	if vv, ok := v.(bool); ok {
		return vv, nil
	}

	return false, ErrKeyTypeNotMatch{
		Key:    key,
		Expect: "bool",
		Got:    TypeName(v),
	}
}

func (s *store) GetString(key string) (string, error) {
	s.lock.RLock()
	defer s.lock.RUnlock()

	v, err := s.readValue(key)
	if err != nil {
		return "", err
	}

	if vv, ok := v.(string); ok {
		return vv, nil
	}

	return "", ErrKeyTypeNotMatch{
		Key:    key,
		Expect: "string",
		Got:    TypeName(v),
	}
}

func (s *store) GetMap(key string) (map[string]interface{}, error) {
	s.lock.RLock()
	defer s.lock.RUnlock()

	v, err := s.readValue(key)
	if err != nil {
		return nil, err
	}

	if vv, ok := v.(map[string]interface{}); ok {
		return vv, nil
	}

	return nil, ErrKeyTypeNotMatch{
		Key:    key,
		Expect: "map",
		Got:    TypeName(v),
	}
}

func (s *store) GetList(key string) ([]interface{}, error) {
	s.lock.RLock()
	defer s.lock.RUnlock()

	v, err := s.readValue(key)
	if err != nil {
		return nil, err
	}

	if vv, ok := v.([]interface{}); ok {
		return vv, nil
	}

	return nil, ErrKeyTypeNotMatch{
		Key:    key,
		Expect: "list",
		Got:    TypeName(v),
	}
}
