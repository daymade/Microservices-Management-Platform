//go:build integration
// +build integration

package storage

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestPostgresStorage_Integration(t *testing.T) {
	storage, err := NewPostgresStorage()
	require.NoError(t, err)

	t.Run("ListServices", func(t *testing.T) {
		services, total, err := storage.ListServices("", "created_at", "desc", 1, 10)
		assert.NoError(t, err)
		assert.Greater(t, total, 0)
		assert.NotEmpty(t, services)
	})

	t.Run("GetService", func(t *testing.T) {
		service, err := storage.GetService("1")
		assert.NoError(t, err)
		assert.NotEmpty(t, service.Name)
	})

	// 添加更多测试用例...
}
