package testhelpers

import (
	"errors"
	"strings"
	"sync"
)

var ErrFunctionShouldFail = errors.New("function should fail / return an error")

type MockPassFail interface {

	// SetCode - pass in a code consisting of "F" and "P". "F" for FAIL, "P" for pass.
	// example code: "FFP", meaning fail twice, and then pass the third time.
	SetCode(code string)

	// WillPass - return if the function should pass or not. Will also increment the count.
	WillPassIncrementCount() error

	GetCount() int
}

type impl struct {
	mutex sync.Mutex
	code  string
	count int
}

func NewMockPassFail() MockPassFail {
	i1 := impl{
		mutex: sync.Mutex{},
		code:  "",
		count: 0,
	}
	return &i1
}

func (i1 *impl) WillPassIncrementCount() error {
	var firstChar uint8 = 'P'
	i1.mutex.Lock()
	defer i1.mutex.Unlock()
	i1.count += 1
	if i1.code != "" {
		firstChar = i1.code[0]
		i1.code = i1.code[1:]
	}
	var returnErr error = nil
	if firstChar == 'F' {
		returnErr = ErrFunctionShouldFail
	}
	return returnErr
}

func (i1 *impl) GetCount() int { return i1.count }

func (i1 *impl) SetCode(code string) { i1.code = strings.ToUpper(code) }
