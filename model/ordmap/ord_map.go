package ordmap

import (
	"fmt"
	"sync"
)

type OrdMap struct {
	// Contains the key-value store
	m map[string]string

	// Contains a linked list of keys for ordering,
	// where head is the first and tail is the latest
	list linkedList

	mu sync.Mutex
}

func New() OrdMap {
	return OrdMap{
		m:    make(map[string]string),
		list: *NewLinkedList(),
		mu:   sync.Mutex{},
	}
}

// Update updates the map with the key and corresponding value if either:
// (1) there is no existing key in the map; or
// (2) the key exists but the corresponding value is different
// In addition, in the both cases the linked list is updated where the key
// is appended to the end of the list in order to keep track of order.
// If the same key-value entry already exists, nothing happens.
func (o *OrdMap) Update(k, v string) {
	o.mu.Lock()
	defer o.mu.Unlock()

	// if the key-value pair already exists, ignore
	if o.m[k] == v {
		return
	}

	// if the key already exists but with a different value, we
	// update the map and the linked list
	if val, ok := o.m[k]; ok && val != v {
		o.m[k] = v
		o.list.Swap(k)
		return
	}

	// this is a new entry so add to map and list
	o.m[k] = v
	o.list.Append(k)
}

func (o *OrdMap) Delete(k string) {
	o.mu.Lock()
	defer o.mu.Unlock()

	// key is missing, noop
	if _, ok := o.m[k]; !ok {
		return
	}

	// wipe the value from the key
	o.m[k] = ""
	o.list.Delete(k)
}

func (o *OrdMap) Get(k string) string {
	return fmt.Sprintf("%s,%s", k, o.m[k])
}

func (o *OrdMap) GetAll() []string {
	var result []string

	keys := o.list.ToSlice()
	for _, key := range keys {
		s := fmt.Sprintf("%s,%s\n", key, o.m[key])
		result = append(result, s)
	}

	return result
}
