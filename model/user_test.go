package model

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewUser(t *testing.T) {
	id, name, email := 1, "Test", "dummy@test.com"
	user := NewUser(id, name, email)

	assert.Equal(t, name, user.Name)
	assert.Equal(t, email, user.Email)

	assert.NotNil(t, user.Id)
}
