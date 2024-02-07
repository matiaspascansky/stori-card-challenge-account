package utils

import (
	"math/rand"
	"sync"
	"time"
)

const (
	letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	idLength    = 10
)

// AccountIDGenerator provides auto-incremented alphanumeric strings for account IDs
type AccountIDGenerator struct {
	mutex sync.Mutex
	rand  *rand.Rand
}

// NewAccountIDGenerator creates a new instance of AccountIDGenerator
func NewAccountIDGenerator() *AccountIDGenerator {
	return &AccountIDGenerator{
		rand: rand.New(rand.NewSource(time.Now().UnixNano())),
	}
}

// GenerateID generates a new auto-incremented alphanumeric string
func (g *AccountIDGenerator) GenerateID() string {
	g.mutex.Lock()
	defer g.mutex.Unlock()
	b := make([]byte, idLength)
	for i := range b {
		b[i] = letterBytes[g.rand.Intn(len(letterBytes))]
	}
	return string(b)
}
