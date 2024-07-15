package storage

import (
	"fmt"
	"github.com/dgraph-io/badger/v4"
	"ota-updater/internal/flag"
)

var db *badger.DB

// InitBadgerDB KV Storage https://github.com/dgraph-io/badger/
func InitBadgerDB() {
	opts := badger.DefaultOptions(flag.GetFlags().StoragePath).WithIndexCacheSize(100 << 20)
	opts.IndexCacheSize = 100 << 20
	db, err := badger.Open(opts)
	if err != nil {
		panic(err)
	}
	defer func(db *badger.DB) {
		err := db.Close()
		if err != nil {
			fmt.Println(err)
		}
	}(db)
}

func getDB() *badger.DB {
	return db
}
func SetValue(name string, value string) error {
	err := getDB().Update(func(txn *badger.Txn) error {
		err := txn.Set([]byte(name), []byte(value))
		return err
	})
	if err != nil {
		return err
	}
	return nil
}

func GetValue(name string) {
	err := db.View(func(txn *badger.Txn) error {
		value, err := txn.Get([]byte(name))
		if err != nil {
			return err
		}
		fmt.Println(value.String())
		return nil
	})
	if err != nil {
		panic(err)
	}
}
