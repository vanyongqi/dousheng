package Test

import (
	"dousheng-backend/Databases"
	"testing"
)

func TestDatabaseSessions(t *testing.T) {
	t.Run("", func(t *testing.T) {
		Databases.InitDatabase()
	})

}
