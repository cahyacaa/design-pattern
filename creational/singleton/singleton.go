package singleton

import (
	"errors"
	"sync"
)

var db *DBObjectExample
var once sync.Once

type DBObjectExample struct {
	TotalCalled int
}

func InitDBExample(dbInstance DBObjectExample) error {

	once.Do(func() {
		db = &dbInstance
	})
	if db == nil {
		return errors.New("db instance is nil!")
	}

	return nil
}

func GetInstance() *DBObjectExample {
	return db
}
