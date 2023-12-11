package main

import (
	"sync"
	"time"
)

var (
	cachelock sync.Mutex
)

type Cache struct {
	uploads map[string]*UploadMetaData
	stored  map[string]*StoreMetaData
}

func NewCache() *Cache {
	res := &Cache{
		uploads: make(map[string]*UploadMetaData),
		stored:  make(map[string]*StoreMetaData),
	}
	return res
}

// SetUpload: sets a key and deletes it 5 minutes later
func (cache *Cache) SetUpload(key string, item *UploadMetaData) {
	cachelock.Lock()
	cache.uploads[key] = item
	cachelock.Unlock()

	go func() {
		time.Sleep(5 * time.Minute)
		cache.DelUpload(key)
	}()
}

// GetUpload :
func (cache *Cache) GetUpload(key string) *UploadMetaData {
	cachelock.Lock()
	x := cache.uploads[key]
	cachelock.Unlock()
	return x
}

// DelUpload :
func (cache *Cache) DelUpload(key string) {
	//	fmt.Println("Cleaning up upload metadata...")
	cachelock.Lock()
	delete(cache.uploads, key)
	cachelock.Unlock()
}

// SetStored : sets a key and deletes it 5 minutes later
func (cache *Cache) SetStored(id string, item *StoreMetaData) {
	cachelock.Lock()
	cache.stored[id] = item
	cachelock.Unlock()

	go func() {
		time.Sleep(5 * time.Minute)
		cache.DelStored(id)
	}()
}

// GetStored :
func (cache *Cache) GetStored(id string) *StoreMetaData {
	cachelock.Lock()
	x := cache.stored[id]
	cachelock.Unlock()
	return x
}

//DelStored :
func (cache *Cache) DelStored(id string) {
	//	fmt.Println("Cleaning up stored metadata...")
	cachelock.Lock()
	delete(cache.stored, id)
	cachelock.Unlock()
}










