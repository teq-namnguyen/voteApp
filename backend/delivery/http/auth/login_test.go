package auth_test

import (
	"github.com/namnguyen/backend/fixture/database"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestLogin(t *testing.T) {
	db := database.InitDatabase()
	defer db.TruncateTables()
	//repo := repository.New(func(c context.Context) *gorm.DB { return db.DB })
	//r := user.Route{UseCase: usecase.New(repo)}
	t.Run("run", func(t *testing.T) {
		require.Equal(t, 1, 1)
	})
}
