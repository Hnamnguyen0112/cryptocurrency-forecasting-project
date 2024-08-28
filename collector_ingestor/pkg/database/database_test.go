package database

import (
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func TestConnect(t *testing.T) {
	testDB, mock, err := sqlmock.New()

	dialector := postgres.New(postgres.Config{
		DSN:                  "sqlmock_db_0",
		DriverName:           "postgres",
		Conn:                 testDB,
		PreferSimpleProtocol: true,
	})

	DB, err := Connect(dialector)
	assert.NoError(t, err)
	assert.NotNil(t, DB)

	sqlDB, err := DB.DB()
	assert.NoError(t, err)
	defer sqlDB.Close()

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}

func TestDisconnect(t *testing.T) {
	testDB, mock, err := sqlmock.New()

	dialector := postgres.New(postgres.Config{
		DSN:                  "sqlmock_db_0",
		DriverName:           "postgres",
		Conn:                 testDB,
		PreferSimpleProtocol: true,
	})

	db, _ := gorm.Open(dialector, &gorm.Config{})

	DB = db

	mock.ExpectClose()

	err = Disconnect()
	assert.NoError(t, err)

	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}
