package jwt

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestNewTokenManager(t *testing.T) {
	tokenManager := NewJwt()
	assert.NotNil(t, tokenManager)
}

func TestTokenManager_GenerateToken(t *testing.T) {
	tokenManager := NewJwt()
	_, err := tokenManager.GenerateToken(uint(1), time.Minute)
	assert.NoError(t, err)
}

func TestTokenManager_DecodeToken(t *testing.T) {
	tokenManager := NewJwt()
	accessToken, err := tokenManager.GenerateToken(uint(1), time.Minute)
	assert.NoError(t, err)

	userID, err := tokenManager.DecodeToken(accessToken)
	assert.NoError(t, err)
	assert.Equal(t, uint(1), userID)
}

func TestTokenManager_claimsFromToken(t *testing.T) {
	tokenManager := NewJwt()
	accessToken, err := tokenManager.GenerateToken(uint(1), time.Minute)
	assert.NoError(t, err)

	_, err = tokenManager.claimsFromToken(accessToken)
	assert.NoError(t, err)
}
