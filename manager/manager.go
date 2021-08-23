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
	kvstore.Dumper
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

	ApplyByJSON(data []byte) error

	Unsafe() UnsafeManager
}

var _ Manager = (*manager)(nil)

type manager struct {
	kvstore.Store

	full   string
	append string

	appendFile *os.File
}

// NewManager 创建一个新的 Manager 对象，传入的 File 应当是支持 R/W/A 的
func NewManager(full, append string) (Manager, error) {
	var store kvstore.Store

	if fullFile, err := os.ReadFile(full); err == nil {
		store, err = kvstore.Load(fullFile)
		if err != nil {
			return nil, ErrFullFileBroken{Filename: full, Reason: err}
		}
	} else if os.IsNotExist(err) {
		store = kvstore.NewStore()
	} else {
		return nil, ErrFullFileBroken{Filename: full, Reason: err}
	}

	if appendFile, err := os.OpenFile(append, os.O_RDONLY, 0644); err == nil {
		r := sio.NewReader(appendFile)

		err := func() error {
			defer appendFile.Close()

			for {
				size, err := r.ReadVarUInt()
				if err != nil {
					if err == io.EOF {
						return nil
					}
					return ErrAppendFileBroken{Filename: append, DecodeError: err}
				}
				if size > math.MaxInt32 {
					return ErrAppendFileBroken{Filename: append, DecodeError: fmt.Errorf("uvarint %d overflow", size)}
				}
				result, err := r.ReadEnoughBytes(int(size))
				if err != nil {
					return ErrAppendFileBroken{Filename: append, DecodeError: err}
				}
				var action kvstore.Action
				err = proto.Unmarshal(result, &action)
				if err != nil {
					return ErrAppendFileBroken{Filename: append, DecodeError: err}
				}

				err = store.Unsafe().UnsafeApplyAction(&action)
				if err != nil {
					return ErrAppendFileBroken{Filename: append, DecodeError: err}
				}
			}
		}()

		if err != nil {
			return nil, err
		}
	} else if !os.IsNotExist(err) {
		return nil, ErrAppendFileBroken{Filename: append, OpenError: err}
	}

	appendFile, err := os.OpenFile(append, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0644)
	if err != nil {
		return nil, ErrAppendFileCannot{"open for append", err}
	}

	return &manager{
		Store:      store,
		full:       full,
		append:     append,
		appendFile: appendFile,
	}, nil
}
