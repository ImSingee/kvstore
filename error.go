package kvstore

import (
	"fmt"
)

type ErrKeyNotExist struct {
	Key string
	On  string
}

func (e ErrKeyNotExist) Error() string {
	if e.On == "" || e.Key == e.On {
		return fmt.Sprintf("key %s not exist", e.Key)
	} else {
		return fmt.Sprintf("key %s not exist: miss %s", e.Key, e.On)
	}
}

type ErrKeyTypeNotMatch struct {
	Key    string
	On     string
	Expect string
	Got    string
}

func (e ErrKeyTypeNotMatch) Error() string {
	if e.On == "" || e.Key == e.On {
		return fmt.Sprintf("key %s is invalid: type expected to be %s, but got %s", e.Key, e.Expect, e.Got)
	} else {
		return fmt.Sprintf("key %s is invalid: type of %s is %s, but it should be %s", e.Key, e.On, e.Got, e.Expect)
	}
}

type ErrKeyIndexNotValid struct {
	Key   string
	On    string
	Index int
	Max   int
}

func (e ErrKeyIndexNotValid) Error() string {
	err := ""

	if e.On == "" || e.Key == e.On {
		err += fmt.Sprintf("key %s index error: ", e.Key)
	} else {
		err += fmt.Sprintf("key %s is invalid: %s index error: ", e.Key, e.On)
	}

	if e.Index >= 0 {
		if e.Max >= 0 {
			err += fmt.Sprintf("index %d is greater than max %d", e.Index, e.Max)
		} else {
			err += "array is empty"
		}
	} else {
		err += "index is not int"
	}

	return err
}

func ImpossibleError() error {
	return fmt.Errorf("impossible")
}
