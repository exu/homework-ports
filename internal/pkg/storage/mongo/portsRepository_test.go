package mongo

import (
	"ports/internal/pkg/pb"
	"ports/internal/pkg/test/assert"
	"testing"
)

func TestPortsRepository(t *testing.T) {
	repo, err := NewPortsRepository()
	assert.NoError(t, err)

	t.Run("inserts data", func(t *testing.T) {
		port := &pb.Port{Code: "A"}
		err := repo.Save(port)
		assert.NoError(t, err)

		port, exists := repo.Get("A")
		assert.Equals(t, true, exists)
		assert.Equals(t, "A", port.Code)

		repo.Delete(port.Code)

	})

	t.Run("change data", func(t *testing.T) {
		port := &pb.Port{Code: "A"}
		err := repo.Save(port)
		assert.NoError(t, err)

		port.Name = "Gdansk"

		err = repo.Save(port)
		assert.NoError(t, err)

		port, _ = repo.Get(port.Code)
		assert.Equals(t, "Gdansk", port.Name)
		repo.Delete(port.Code)
	})

	t.Run("list data", func(t *testing.T) {
		port1 := &pb.Port{Code: "A"}
		port2 := &pb.Port{Code: "B"}
		err := repo.Save(port1)
		assert.NoError(t, err)
		err = repo.Save(port2)
		assert.NoError(t, err)

		ports, err := repo.List("A", "B")
		assert.NoError(t, err)
		assert.Equals(t, 2, len(ports))
	})
}
