package redas

import (
	"testing"
	"time"

	"github.com/my0sot1s/godef/log"
)

func Test_set_get(t *testing.T) {
	l := LocalStore{}
	l.InitLocalStore(24 * time.Hour)
	l.PushKey("test", "test1", "abc")
	l.PushKey("test", "test2", "zxv")
	l.PushKey("test", "test3", "iop")

	l2 := l.GetAll("test")

	for _, v := range l2 {
		log.Log(string(v))
	}
}

func Test_get_dele(t *testing.T) {
	l := LocalStore{}
	l.InitLocalStore(24 * time.Hour)
	l.PushKey("test", "test1", "abc")
	l.PushKey("test", "test2", "zxv")
	l.PushKey("test", "test3", "iop")

	l2 := l.GetAll("test")

	for _, v := range l2 {
		log.Log(string(v))
	}

	l.deleteKeySpace("test", "test1")
	l2 = l.GetAll("test")
	for _, v := range l2 {
		log.Log(string(v))
	}
}

func Test_get_deleAll(t *testing.T) {
	l := LocalStore{}
	l.InitLocalStore(24 * time.Hour)
	l.PushKey("test", "test1", "abc")
	l.PushKey("test", "test2", "zxv")
	l.PushKey("test", "test3", "iop")

	l2 := l.GetAll("test")

	for _, v := range l2 {
		log.Log(string(v))
	}

	l.deleteKeyRoot("test")
	l2 = l.GetAll("test")
	for _, v := range l2 {
		log.Log(string(v))
	}
}

func Test_update(t *testing.T) {
	l := LocalStore{}
	l.InitLocalStore(24 * time.Hour)
	l.PushKey("test", "test1", "abc")
	l.PushKey("test", "test2", "zxv")
	l.PushKey("test", "test3", "iop")

	l2 := l.GetAll("test")

	for _, v := range l2 {
		log.Log(string(v))
	}

	l.PushKey("test", "test1", "ffgfg")
	l2 = l.GetAll("test")
	for _, v := range l2 {
		log.Log(string(v))
	}
}
