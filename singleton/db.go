package singleton

import (
	"fmt"
	"sync"
)

type DB struct{}
type Elastic struct{}

var (
	instance        *DB
	elasticInstance *Elastic

	mu   = &sync.Mutex{}
	once = &sync.Once{}
)

func GetDBInsstanceWithLock() *DB {
	if instance == nil {
		mu.Lock()
		instance = &DB{}
		mu.Unlock()
		fmt.Println("set instance")
	}
	return instance
}

func GetElasticInstanceWithOnce() *Elastic {
	once.Do(func() {
		elasticInstance = &Elastic{}
		fmt.Println("set instance")
	})
	return elasticInstance
}
