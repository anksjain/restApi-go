package instance

import (
	"restsample/db"
	"sync"
)

type instance struct {
	db *db.UserDb
}

var singleton = &instance{}
var once sync.Once

func Init() {
	once.Do(func() {
		// seeding the memory database
		singleton.db = db.Seed()
	})

}

// getting the db object
func GetDb() *db.UserDb {
	return singleton.db
}
