package config

import (
	"ethproxy/models"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMock(t *testing.T) {
	mock := NewMock()

	// Must not panic
	mock.GetConfig()
	mock.DebugMode()

	// Classic mock
	var called bool

	mock.GetConfigFunc = func() models.Config {
		called = true
		return models.Config{}
	}
	mock.DebugModeFunc = func() bool {
		called = true
		return false
	}

	called = false
	mock.GetConfig()
	assert.True(t, called)

	called = false
	mock.DebugMode()
	assert.True(t, called)

}
