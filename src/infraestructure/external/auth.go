package infrastructure

import (
	"sync"
	"time"
)

// AuthManager implementa AuthService
type AuthManager struct {
	mu          sync.RWMutex
	vercelToken string
	expires     time.Time
}

// NewAuthManager crea una nueva instancia de AuthManager
func NewAuthManager() *AuthManager {
	return &AuthManager{}
}

func (a *AuthManager) GetToken() string {
	a.mu.RLock()
	defer a.mu.RUnlock()
	return a.vercelToken
}

func (a *AuthManager) SetToken(token string) {
	a.mu.Lock()
	defer a.mu.Unlock()
	a.vercelToken = token
	a.expires = time.Now().Add(1 * time.Hour)
}

func (a *AuthManager) IsTokenValid() bool {
	a.mu.RLock()
	defer a.mu.RUnlock()
	return a.vercelToken != "" && time.Now().Before(a.expires)
}
