package errors

import (
	"errors"
	"sync"
)

var lastError = &errorKeeper{}

type errorKeeper struct {
	err error
	sync.RWMutex
}

func (k *errorKeeper) read() error {
	for !k.TryRLock() {
	}
	defer k.RUnlock()

	return k.err
}

func (k *errorKeeper) store(err error) {
	for !k.TryLock() {
	}
	k.err = err
	k.Unlock()
}

func LastError() error {
	return lastError.read()
}

func LastErrorWas(err error) bool {
	return errors.Is(lastError.read(), err)
}
