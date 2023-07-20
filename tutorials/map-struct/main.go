// Golang program to show how to
// use structs as map keys
package main

// importing required packages
import (
	"fmt"
	"sync"
)

// concurent safe map
type SafeMap[K comparable, V any] struct {
	mu   sync.RWMutex
	data map[K]V
}

// constructor
func NewSafeMap[K comparable, V any]() *SafeMap[K, V] {
	return &SafeMap[K, V]{
		data: make(map[K]V), // we dont need to specify mutex, automaticly will be added with 0 value
	}
}

func (m *SafeMap[K, V]) Insert(key K, value V) {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.data[key] = value
}

func (m *SafeMap[K, V]) Get(key K) (V, error) {
	m.mu.RLock() // readLock
	defer m.mu.RUnlock()

	// we check if we have key
	value, ok := m.data[key]
	if !ok {
		return value, fmt.Errorf("key %v not found", key)
	}
	return value, nil
}

func (m *SafeMap[K, V]) Update(key K, value V) error {
	m.mu.Lock()
	defer m.mu.Unlock()

	// we check if key exist
	_, ok := m.data[key]
	if !ok {
		return fmt.Errorf("key %v not found", key)
	}

	m.data[key] = value
	return nil
}

func (m *SafeMap[K, V]) Delete(key K) error {
	m.mu.Lock()
	defer m.mu.Unlock()
	// we check if key exist
	_, ok := m.data[key]
	if !ok {
		return fmt.Errorf("key %v not found", key)
	}
	return nil
}

func (m *SafeMap[K, V]) HasKey(key K) bool {
	m.mu.RLock()
	_, ok := m.data[key]
	m.mu.RUnlock() // solution without defer, a little more performent, but less readable
	return ok
}

type AddressStudent struct {
	Name    string
	city    string
	Pincode int
}

func main() {
	// SimpleMap()
	m := NewSafeMap[AddressStudent, int]()

	m.Insert(AddressStudent{Name: "Ram", city: "Delhi", Pincode: 2400}, 1)
	fmt.Println(m.data)

	for str, val := range m.data {
		fmt.Println(str.Name, val)
	}

}
