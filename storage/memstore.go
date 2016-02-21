package storage

import (
	"fmt"
	"github.com/loofort/softswiss/models"
	"sync"
)

type notFound string

func (e notFound) Error() string {
	return string(e)
}

func IsNotFound(err error) bool {
	_, ok := err.(notFound)
	return ok
}

func clone(acc *models.Account) *models.Account {
	newacc := *acc
	return &newacc
}

type Storage struct {
	sync.RWMutex
	// data storage and ID index
	// we don't have delete operation, so the storage is simple array
	db []*models.Account
}

func MustConnect(adress string) *Storage {
	return &Storage{
		db: make([]*models.Account, 0),
	}
}

func (s *Storage) AccountItem(id int64) (*models.Account, error) {
	s.RLock()
	defer s.RUnlock()

	if id >= int64(len(s.db)) || id < 0 {
		return nil, notFound(fmt.Sprintf("account id=%d is not found", id))
	}

	return clone(s.db[id]), nil
}

func (s *Storage) AccountList() ([]*models.Account, error) {
	s.RLock()
	defer s.RUnlock()

	accs := make([]*models.Account, 0, len(s.db))
	for _, acc := range s.db {
		accs = append(accs, clone(acc))
	}

	return accs, nil
}

func (s *Storage) AccountInsert(acc *models.Account) (*models.Account, error) {
	s.Lock()
	defer s.Unlock()

	newacc := clone(acc)
	newacc.ID = int64(len(s.db))

	s.db = append(s.db, newacc)

	return clone(newacc), nil
}

func (s *Storage) AccountUpdate(acc *models.Account) error {
	s.Lock()
	defer s.Unlock()

	id := acc.ID
	if id >= int64(len(s.db)) || id < 0 {
		return notFound(fmt.Sprintf("account id=%d is not found", id))
	}

	*s.db[id] = *acc
	return nil
}

// return fake transaction
func (s *Storage) Begin() Tx {
	return Tx{s}
}

type Tx struct {
	*Storage
}

// fake
func (Tx) Commit() error {
	return nil
}
func (Tx) Rollback() {
}
