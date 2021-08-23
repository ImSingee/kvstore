package manager

import (
	"fmt"
	"github.com/ImSingee/kvstore"
	"github.com/ImSingee/sio"
	"google.golang.org/protobuf/proto"
	"os"
)

// Dump 将当前数据全量存储至 DB 文件
// 如果出现极端异常可能会直接 panic
func (m *manager) Dump() error {
	m.lock()
	defer m.unlock()

	// 刷新当前 db 文件
	err := m.dbFile.Sync()
	if err != nil {
		return fmt.Errorf("kvstore manager: cannot fsync db file: %w", err)
	}

	// 备份当前数据
	success := false
	newFile, err := os.OpenFile(m.db+".tmp", os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		return fmt.Errorf("kvstore manager: cannot create new db file: %w", err)
	}
	defer func() {
		if !success {
			newFile.Close()
		}
	}()

	// 编码当前 action 数据
	action := kvstore.NewReplaceAction(m.Store.Unsafe().UnsafeUnderlying())

	// 写入 action 到文件
	w := sio.NewWriter(newFile)
	b, err := proto.Marshal(action)
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

	// 将当前 db 文件关掉并删除
	err = m.dbFile.Close()
	if err != nil {
		panic("kvstore manager: cannot close db file")
	}
	err = os.Remove(m.db)
	if err != nil {
		panic("kvstore manager: cannot remove old db file")
	}

	// 使用新的来替换
	err = os.Rename(m.db+".tmp", m.db)
	if err != nil {
		panic("kvstore manager: cannot rename new db file")
	}
	m.dbFile = newFile

	success = true
	return nil
}
