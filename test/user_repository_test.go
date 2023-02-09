package test

import (
	"context"
	"database/sql"
	"fmt"
	"regexp"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/Jocerdikiawann/vm-go/models/request"
	"github.com/Jocerdikiawann/vm-go/repository"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func dbMock(t *testing.T) (*sql.DB, *gorm.DB, sqlmock.Sqlmock) {
	sqlDb, mock, err := sqlmock.New()

	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	gormDb, err := gorm.Open(postgres.New(postgres.Config{
		DSN:                  "sqlmock_db_0",
		DriverName:           "postgres",
		Conn:                 sqlDb,
		PreferSimpleProtocol: true,
	}), &gorm.Config{})

	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	return sqlDb, gormDb, mock
}

func TestCreatedOne(t *testing.T) {
	sqlDb, db, mock := dbMock(t)

	defer sqlDb.Close()

	r := repository.NewUserRepository(db)

	expectedQueryInsertUser := `INSERT INTO "users" 
	("email","password","firstName","lastName","createdAt","updatedAt") 
	VALUES ($1,$2,$3,$4,$5,$6) RETURNING "id"`

	// expectedQueryInsertUsersRoles := `INSERT INTO "users_roles"
	// ("user_id","role_id")
	// VALUES ($1,$2)`

	t.Run("Must return the newly created id, if given user input is valid", func(t *testing.T) {
		timeStamp := uint64(time.Now().UnixNano() / int64(time.Millisecond))

		user := request.SignUpRequest{
			Email:     "user@example.com",
			FirstName: "John",
			LastName:  "Doe",
			Password:  "123",
			CreatedAt: timeStamp,
			UpdatedAt: timeStamp,
		}
		mock.MatchExpectationsInOrder(false)
		mock.ExpectBegin()
		mock.ExpectExec(regexp.QuoteMeta(expectedQueryInsertUser)).
			WithArgs(user.Email, user.Password, user.FirstName,
				user.LastName, user.CreatedAt, user.UpdatedAt).
			WillReturnResult(sqlmock.NewResult(1, 1))
		mock.ExpectCommit()

		// mock.ExpectExec(expectedQueryInsertUsersRoles).
		// 	WithArgs(1, 1).WillReturnResult(sqlmock.NewResult(1, 1))

		_, err := r.CreateOne(context.Background(), user)
		fmt.Println(err)
		// assert.Nil(t, err)
		// assert.NotEqual(t, 0, id.Id)
		// assert.Nil(t, mock.ExpectationsWereMet())
	})

}
