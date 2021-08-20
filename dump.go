package kvstore

import "github.com/golang/protobuf/proto"

type Dumper interface {
	Dump() ([]byte, error)
}

var _ Dumper = (*store)(nil)

func (s *store) Dump() ([]byte, error) {
	s.lock.RLock()
	defer s.lock.RUnlock()

	return proto.Marshal(s.Provider)
}
