package postgresql

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"orm/pkg/logger"
	"os"
	"sync"
	"time"
)

var (
	user     = os.Getenv("USER_DB")
	password = os.Getenv("PASSWORD_DB")
	ip       = os.Getenv("IP_DB")
	port     = os.Getenv("PORT_DB")
)

var (
	once  sync.Once
	db    *gorm.DB
	dbErr error
)

func ConnDB() (*gorm.DB, error) {
	once.Do(func() {
		log := logger.NewLogger("ERROR")

		dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", ip, user, password, "go_avanzado", port)

		db, dbErr = gorm.Open(postgres.Open(dsn), &gorm.Config{})

		if dbErr != nil {
			log.Fatal(dbErr.Error())
		}

		log.Info("Connected!")

		configPool()

	})

	return db, dbErr
}

//Esto debe ser parametrico
func ConfigShema[T interface{}](class T) {
	// Migrate the schema
	db.AutoMigrate(class)
}

func configPool() {
	sqlDB, err := db.DB()

	// SetMaxIdleConns sets the maximum number of connections in the idle connection pool.
	sqlDB.SetMaxIdleConns(10)

	// SetMaxOpenConns sets the maximum number of open connections to the database.
	sqlDB.SetMaxOpenConns(100)

	// SetConnMaxLifetime sets the maximum amount of time a connection may be reused.
	sqlDB.SetConnMaxLifetime(time.Hour)

	if err != nil {
		log.Fatal(fmt.Sprintf("Error configurando el pool de conexiones..... %v", err))
	}

}
