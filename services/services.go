package services

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var (
	data = make(map[string]string)
	mut  sync.RWMutex
)

func GenerateValue() (string, string) {
	mut.Lock()
	defer mut.Unlock()

	source := rand.NewSource(time.Now().UnixNano())
	rnd := rand.New(source)
	id := fmt.Sprint(rnd.Intn(100000))
	value := fmt.Sprint(rnd.Int())
	data[id] = value
	return id, value
}

func GetValue(id string) (string, bool) {
	mut.RLock()
	defer mut.RUnlock()

	value, isFound := data[id]
	return value, isFound
}
