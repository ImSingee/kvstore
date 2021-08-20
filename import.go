package kvstore

import "fmt"

type Applier interface {
	ReplaceInner(new *Map)
	Apply(actions []*Action) error
	ApplyAction(action *Action) error
	MockApply(actions []*Action) (Store, error)

	ApplyByJSON(data []byte) error
}

// ReplaceInner 替换内部的数据
// 该操作将使 UnsafeUnderlying() 返回的内部数据无效
// 该函数不会复制 new，但请注意不要使用另外的函数对 new 进行修改
func (s *store) ReplaceInner(new *Map) {
	s.lock.Lock()
	defer s.lock.Unlock()

	s.replaceInner(new)
}
func (s *store) replaceInner(new *Provider) {
	s.Provider = new
}

// UnsafeReplaceInnerWithoutLock like ReplaceInner but won't use lock
func (s *store) UnsafeReplaceInnerWithoutLock(new *Map) {
	s.replaceInner(new)
}

// Apply 过程会基于 store 执行若干命令
func (s *store) Apply(actions []*Action) error {
	s.lock.Lock()
	defer s.lock.Unlock()

	return s.apply(actions)
}
func (s *store) UnsafeApply(actions []*Action) error {
	return s.apply(actions)
}
func (s *store) apply(actions []*Action) error {
	newStore := s.clone().(*store)
	err := applyOn(newStore, actions)
	if err != nil {
		return err
	}
	s.replaceInner(newStore.Provider)
	return nil
}

func (s *store) ApplyAction(action *Action) error {
	s.lock.Lock()
	defer s.lock.Unlock()

	return applyOn(s, []*Action{action})
}

func (s *store) UnsafeApplyAction(action *Action) error {
	return applyOn(s, []*Action{action})
}

// MockApply 会返回基于 store 虚拟执行一系列 action 的结果
func (s *store) MockApply(actions []*Action) (Store, error) {
	s.lock.RLock()
	defer s.lock.RUnlock()

	newStore := s.clone().(*store)
	err := applyOn(newStore, actions)
	if err != nil {
		return nil, err
	}
	return newStore, nil
}

// UnsafeMockApplyAndLock like MockApply but will Lock store
func (s *store) UnsafeMockApplyAndLock(actions []*Action) (Store, error) {
	s.lock.Lock()

	newStore := s.clone().(*store)
	err := applyOn(newStore, actions)
	if err != nil {
		return nil, err
	}
	return newStore, nil
}

// applyOn 调用方需注意 store 不能上锁
func applyOn(s *store, actions []*Action) error {
	s.lock.Lock()
	defer s.lock.Unlock()

	for _, action := range actions {
		switch action := action.Unwrap().(type) {
		case *Set:
			err := s.setValueByValue(action.Key, action.Value)
			if err != nil {
				return err
			}
		case *Delete:
			err := s.delete(action.Key, false)
			if err != nil {
				return err
			}
		case *Replace:
			s.replaceInner(action.New)
		default:
			return fmt.Errorf("unknown action %T", action)
		}
	}

	return nil
}
