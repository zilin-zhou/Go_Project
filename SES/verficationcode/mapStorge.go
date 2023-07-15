package verficationcode

import (
	"errors"
	"sync"
	"time"
)

var ms *mapStorage

type item struct {
	Val            string
	ExpirationTime time.Time
}
type mapStorage struct {
	MeMap map[string]*item
	sync.RWMutex
}

func init() {
	ms = &mapStorage{
		MeMap: make(map[string]*item),
	}
	go deleteInvalidKey() //开启协程删除过期的key值
}
func NewMapStorage() storge {
	return ms
}

func deleteInvalidKey() {
	ticker := time.NewTicker(time.Second * 10) //每十秒触发一次
	for {
		select {
		case <-ticker.C:
			for key, item := range ms.MeMap {
				if item.ExpirationTime.Before(time.Now()) {
					ms.Lock()
					delete(ms.MeMap, key)
					ms.Unlock()
				}
			}
		}
	}
}
func (ms *mapStorage) Set(key, val string, duration time.Duration) error {
	ms.RLock()
	defer ms.RUnlock()
	ms.MeMap[key] = &item{
		Val:            val,
		ExpirationTime: time.Now().Add(duration),
	}
	return nil
}
func (ms *mapStorage) Get(key string) (string, error) {
	ms.RLock()
	defer ms.RUnlock()
	res, ok := ms.MeMap[key]
	if !ok {
		return "", errors.New("key值不存在!")
	}
	if res.ExpirationTime.Before(time.Now()) {
		return "", errors.New("验证码过期!")
	}
	return res.Val, nil
}
