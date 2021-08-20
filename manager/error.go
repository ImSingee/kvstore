package manager

import "fmt"

type ErrFullFileBroken struct {
	Filename string
	Reason   error
}

func (e ErrFullFileBroken) Error() string {
	return fmt.Sprintf("kvstore manager: full db file %s is broken: %v", e.Filename, e.Reason)
}

func (e ErrFullFileBroken) Unwrap() error {
	return e.Reason
}

type ErrAppendFileBroken struct {
	Filename    string
	OpenError   error
	DecodeError error
}

func (e ErrAppendFileBroken) Error() string {
	if e.OpenError != nil {
		return fmt.Sprintf("kvstore manager: append file %s cannot open: %v", e.Filename, e.OpenError)
	} else {
		return fmt.Sprintf("kvstore manager: append file %s is broken: %v", e.Filename, e.DecodeError)
	}
}

func (e ErrAppendFileBroken) Unwrap() error {
	if e.OpenError != nil {
		return e.OpenError
	} else {
		return e.DecodeError
	}
}

type ErrAppendFileCannot struct {
	Action     string
	Underlying error
}

func (e ErrAppendFileCannot) Error() string {
	return fmt.Sprintf("kvstore manager: append file cannot %s: %v", e.Action, e.Underlying)
}

func (e ErrAppendFileCannot) Unwrap() error {
	return e.Underlying
}
