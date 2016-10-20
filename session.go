package flushbbs

import (
	"sync"
	"fmt"
)

type Manager struct {
	cookieName string
	lock sync.Mutex
	provider Provider
	maxLifetime int64
}

type Provider interface {
	SessionInit(sid string) (Session, error)
	SessionRead(sid string) (Session, error)
	SessionGc(maxLifetime int64)
	SessionDestroy() error
}

type Session interface {
	Set(key, val interface{}) error
	Get(key interface{}) interface{}
	Delete(key interface{}) error
	SessionId() string
}

var provides = make(map[string]Provider)

func NewManager(provideName, cookieName string, maxLifetime int64) (*Manager ,error) {

	provider, ok := provides[provideName]
	if !ok {
		return nil, fmt.Errorf("session: unknown provide %q (forgotten import?)", provideName)
	}
	return &Manager{provider: provider, cookieName: cookieName, maxLifetime: maxLifetime}, nil
}

func Register(name string, provider Provider) {
	if provider == nil {
		panic("session: Register provide is nil")
	}
	if _, dup := provides[name]; dup {
		panic("session: Register called twice for provide " + name)
	}
	provides[name] = provider
}

func main() {
}
