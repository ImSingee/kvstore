package manager

import (
	"fmt"
	"github.com/ImSingee/kvstore"
	"github.com/ImSingee/sio"
	"github.com/golang/protobuf/proto"
	"io"
	"math"
	"os"
)

type Manager interface {
	kvstore.Getter
	kvstore.Checker
	kvstore.Cloner
	kvstore.Exporter
	kvstore.KeySplitter

	Close()

	Set(key string, value *kvstore.Value) error
	SetNull(key string) error
	SetInt64(key string, value int64) error
	SetUint64(key string, value uint64) error
	SetFloat64(key string, value float64) error
	SetBool(key string, value bool) error
	SetString(key string, value string) error
	SetTrue(key string) error
	SetFalse(key string) error
	SetList(key string, value []interface{}) error
	SetMap(key string, value map[string]interface{}) error

	Delete(key string) error
	Clear() error

	ApplyByJSON(data []byte) error

	Dump() error

	Unsafe() UnsafeManager
	IsReadonly() bool
}

var _ Manager = (*manager)(nil)

type manager struct {
	kvstore.Store

	db     string
	dbFile *os.File

	ro bool
}

// NewManager 创建一个新的 Manager 对象
func NewManager(dbFilename string) (Manager, error) {
	return newManager(dbFilename, false)
}
func NewReadonlyManager(dbFilename string) (Manager, error) {
	return newManager(dbFilename, true)
}

func newManager(dbFilename string, readonly bool) (Manager, error) {
	store := kvstore.NewStore()

	if dbFile, err := os.OpenFile(dbFilename, os.O_RDONLY, 0644); err == nil {
		r := sio.NewReader(dbFile)

		err := func() error {
			defer dbFile.Close()

			for {
				size, err := r.ReadVarUInt()
				if err != nil {
					if err == io.EOF {
						return nil
					}
					return ErrDBFileBroken{Filename: dbFilename, DecodeError: err}
				}
				if size > math.MaxInt32 {
					return ErrDBFileBroken{Filename: dbFilename, DecodeError: fmt.Errorf("uvarint %d overflow", size)}
				}
				result, err := r.ReadEnoughBytes(int(size))
				if err != nil {
					return ErrDBFileBroken{Filename: dbFilename, DecodeError: err}
				}
				var action kvstore.Action
				err = proto.Unmarshal(result, &action)
				if err != nil {
					return ErrDBFileBroken{Filename: dbFilename, DecodeError: err}
				}

				err = store.Unsafe().UnsafeApplyAction(&action)
				if err != nil {
					return ErrDBFileBroken{Filename: dbFilename, DecodeError: err}
				}
			}
		}()

		if err != nil {
			return nil, err
		}
	} else if !os.IsNotExist(err) {
		return nil, ErrDBFileBroken{Filename: dbFilename, OpenError: err}
	}

	var appendFile *os.File
	var err error
	if !readonly {
		appendFile, err = os.OpenFile(dbFilename, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0644)
		if err != nil {
			return nil, ErrDBFileCannot{"open for append", err}
		}
	}

	return &manager{
		Store:  store,
		db:     dbFilename,
		dbFile: appendFile,
		ro:     readonly,
	}, nil
}
