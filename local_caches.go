package redas

import (
	"encoding/json"
	"time"

	"github.com/my0sot1s/godef/log"
)

type KeySpace = map[string][]byte

type LocalStore struct {
	store   map[string]KeySpace
	timeOut time.Duration
}

func (ls *LocalStore) setKeyRoot(keyRoot string) {
	if ls.store[keyRoot] != nil {
		return
	}
	ls.store[keyRoot] = make(KeySpace)
}

func (ls *LocalStore) getKeyRoot(keyRoot string) KeySpace {
	if ls.store[keyRoot] == nil {
		return nil
	}
	var dat = ls.store[keyRoot]
	return dat
}

func (ls *LocalStore) deleteKeyRoot(keyRoot string) {
	if ls.store[keyRoot] == nil {
		return
	}
	ls.store[keyRoot] = nil
}

func (ls *LocalStore) InitLocalStore(duration time.Duration) {
	ls.store = make(map[string]KeySpace)
	ls.timeOut = duration
}

func (ls *LocalStore) setKeySpace(keyRoot, keyspace string, payload []byte) {
	ls.setKeyRoot(keyRoot)
	ls.store[keyRoot][keyspace] = payload
}

func (ls *LocalStore) getKeySpace(keyRoot, keyspace string) []byte {
	kp := ls.getKeyRoot(keyRoot)
	if kp == nil {
		return nil
	}
	return kp[keyspace]
}

func (ls *LocalStore) deleteKeySpace(keyRoot, keyspace string) {
	kp := ls.getKeyRoot(keyRoot)
	if kp == nil || kp[keyspace] == nil {
		return
	}
	delete(ls.store[keyRoot], keyspace)
	if len(ls.getKeyRoot(keyRoot)) == 0 {
		ls.deleteKeyRoot(keyRoot)
	}
}

func (ls *LocalStore) PushKey(root, key string, payload interface{}) error {
	bin, err := json.Marshal(payload)
	if err != nil {
		log.ErrLog(err)
		return err
	}
	ls.setKeySpace(root, key, bin)
	return nil
}

func (ls *LocalStore) GetAll(root string) KeySpace {
	return ls.getKeyRoot(root)
}

func (ls *LocalStore) DeleteKey(root, key string) {
	ls.deleteKeySpace(root, key)
}

func (ls *LocalStore) DeleteAll(root string) {
	ls.deleteKeyRoot(root)
}
