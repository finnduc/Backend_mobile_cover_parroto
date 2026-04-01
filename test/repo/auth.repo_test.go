package repo

import (
	"context"
	"go-familytree/internal/models"
	"go-familytree/internal/repo"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func TestAuthRepo_CreateUser(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("failed to open sqlmock: %s", err)
	}
	defer db.Close()

	gormDB, err := gorm.Open(postgres.New(postgres.Config{
		Conn: db,
	}), &gorm.Config{})
	if err != nil {
		t.Fatalf("failed to open gorm: %s", err)
	}

	r := repo.NewAuthRepo(gormDB)

	user := &models.User{
		Email: "test@example.com",
		Name:  "Test User",
	}

	mock.ExpectBegin()
	mock.ExpectQuery(`INSERT INTO "users"`).
		WithArgs(sqlmock.AnyArg(), sqlmock.AnyArg(), sqlmock.AnyArg(), user.Email, sqlmock.AnyArg(), user.Name, sqlmock.AnyArg(), sqlmock.AnyArg()).
		WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(1))
	mock.ExpectCommit()

	err = r.CreateUser(context.Background(), user)

	assert.NoError(t, err)
	assert.Equal(t, uint(1), user.ID)
	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestAuthRepo_FindUserByEmail(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("failed to open sqlmock: %s", err)
	}
	defer db.Close()

	gormDB, err := gorm.Open(postgres.New(postgres.Config{
		Conn: db,
	}), &gorm.Config{})
	if err != nil {
		t.Fatalf("failed to open gorm: %s", err)
	}

	r := repo.NewAuthRepo(gormDB)

	email := "test@example.com"
	// GORM First() returns all fields usually, but we can mock just what we need
	rows := sqlmock.NewRows([]string{"id", "email", "name", "created_at", "updated_at", "deleted_at"}).
		AddRow(1, email, "Test User", nil, nil, nil)

	mock.ExpectQuery(`SELECT \* FROM "users" WHERE email = \$1 AND "users"\."deleted_at" IS NULL ORDER BY "users"\."id" LIMIT \$2`).
		WithArgs(email, 1).
		WillReturnRows(rows)

	user, err := r.FindUserByEmail(context.Background(), email)

	assert.NoError(t, err)
	assert.NotNil(t, user)
	assert.Equal(t, email, user.Email)
	assert.NoError(t, mock.ExpectationsWereMet())
}
