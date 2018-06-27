package repos

import (
	"clean/models"
	"sync"
)

type TestDB struct {
	mu    *sync.Mutex
	store map[int]string
}

func InitTestDB() *TestDB {
	return &TestDB{store: make(map[int]string), mu: &sync.Mutex{}}
}

func (db *TestDB) newID() int {
	return len(db.store) + 1
}

func (db *TestDB) CreateUser(name string) (int, error) {
	db.mu.Lock()
	defer db.mu.Unlock()
	newUserID := db.newID()
	db.store[newUserID] = name
	return newUserID, nil
}

func (db *TestDB) GetUser(ID int) models.User {
	var user models.User
	db.mu.Lock()
	defer db.mu.Unlock()
	res, ok := db.store[ID]
	if ok {
		user.ID = ID
		user.Name = res
	}
	return user
}
