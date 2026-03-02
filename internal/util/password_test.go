package util

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHashPassword(t *testing.T) {
	password := "TestPassword123!"
	hash, err := HashPassword(password)
	
	assert.NoError(t, err)
	assert.NotEmpty(t, hash)
	assert.NotEqual(t, password, hash)
}

func TestCheckPasswordHash(t *testing.T) {
	password := "TestPassword123!"
	hash, err := HashPassword(password)
	assert.NoError(t, err)

	// Correct password
	assert.True(t, CheckPasswordHash(password, hash))

	// Incorrect password
	assert.False(t, CheckPasswordHash("WrongPassword", hash))
	assert.False(t, CheckPasswordHash("", hash))
}

func TestPasswordHash_Unique(t *testing.T) {
	password := "TestPassword123!"
	hash1, _ := HashPassword(password)
	hash2, _ := HashPassword(password)

	// Each hash should be unique due to salt
	assert.NotEqual(t, hash1, hash2)

	// But both should validate the same password
	assert.True(t, CheckPasswordHash(password, hash1))
	assert.True(t, CheckPasswordHash(password, hash2))
}

