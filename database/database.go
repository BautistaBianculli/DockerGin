package database

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type PostgresStore struct {
	Db *gorm.DB
}

var PostgresDB = NewPostgresStore()

func NewPostgresStore() *PostgresStore {
	dsn := "host=postgres user=postgres password=admin123 dbname=postgres port=5432 sslmode=disable TimeZone=America/Argentina/Buenos_Aires"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		log.Fatal(err.Error())
	}
	return &PostgresStore{
		Db: db,
	}
}

func ConectionCheck() bool {

	sqlDB, err := PostgresDB.Db.DB()
	if err != nil {
		return false
	}

	err = sqlDB.Ping()
	if err != nil {
		return false
	}
	return true
}

// dsn := "host=localhost user=gorm password=gorm dbname=gorm port=9920 sslmode=disable TimeZone=Asia/Shanghai"
// db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
func (s *PostgresStore) Init() error {
	err := s.createUserTable()
	if err != nil {
		return err
	}
	return nil
}

func (s *PostgresStore) createUserTable() error {
	query := `
	CREATE TABLE IF NOT EXISTS users (
		id SERIAL PRIMARY KEY,
		username VARCHAR UNIQUE NOT NULL,
		password VARCHAR NOT NULL,
		create_time TIMESTAMP
	)`

	result := s.Db.Exec(query)
	if result.Error != nil {
		return result.Error
	}

	result = s.Db.Exec("SELECT * FROM users")
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		result = s.Db.Exec("INSERT INTO users (username, password, create_time) VALUES ('admin', 'admin', NOW())")
		if result.Error != nil {
			return result.Error
		}
	}
	return nil
}
