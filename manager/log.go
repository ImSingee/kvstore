package manager

import (
	"bytes"
	"github.com/ImSingee/kvstore"
	"github.com/ImSingee/sio"
	"github.com/golang/protobuf/proto"
	"io"
)

// 内部会上锁，f 中只能使用无锁函数
func (m *manager) writeLogsAndDoE(actions []*kvstore.Action, f func(s kvstore.UnsafeStore) error) error {
	m.lock()
	defer m.unlock()

	buf := bytes.Buffer{}
	w := sio.NewWriter(&buf)

	for _, a := range actions {
		b, err := proto.Marshal(a)
		if err != nil {
			return err
		}

		_, err = w.WriteVarUInt(uint64(len(b)))
		if err != nil {
			return err
		}
		_, err = w.Write(b)
		if err != nil {
			return err
		}
	}

	err := f(m.Store.Unsafe())
	if err != nil {
		return err
	}

	_, err = io.Copy(m.dbFile, &buf)
	if err != nil {
		panic("kvstore manager: cannot write db: " + err.Error())
	}

	return nil
}

func (m *manager) writeLogAndDoE(a *kvstore.Action, f func(s kvstore.UnsafeStore) error) error {
	return m.writeLogsAndDoE([]*kvstore.Action{a}, f)
}

func (m *manager) writeLogAndDo(a *kvstore.Action, f func(s kvstore.UnsafeStore)) error {
	return m.writeLogsAndDoE([]*kvstore.Action{a}, func(s kvstore.UnsafeStore) error {
		f(s)
		return nil
	})
}
