package manager

import "fmt"

type ErrDBFileBroken struct {
	Filename    string
	OpenError   error
	DecodeError error
}

func (e ErrDBFileBroken) Error() string {
	if e.OpenError != nil {
		return fmt.Sprintf("kvstore manager: db file %s cannot open: %v", e.Filename, e.OpenError)
	} else {
		return fmt.Sprintf("kvstore manager: db file %s is broken: %v", e.Filename, e.DecodeError)
	}
}

func (e ErrDBFileBroken) Unwrap() error {
	if e.OpenError != nil {
		return e.OpenError
	} else {
		return e.DecodeError
	}
}

type ErrDBFileCannot struct {
	Action     string
	Underlying error
}

func (e ErrDBFileCannot) Error() string {
	return fmt.Sprintf("kvstore manager: db file cannot %s: %v", e.Action, e.Underlying)
}

func (e ErrDBFileCannot) Unwrap() error {
	return e.Underlying
}
