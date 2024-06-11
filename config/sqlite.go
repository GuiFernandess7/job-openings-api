package config

import (
	"fmt"
	"os"

	"github.com/GuiFernandess7/job-openings-api/schemas"
	"gorm.io/driver/sqlite"

	"gorm.io/gorm"
)

func InitializeSQLite() (*gorm.DB, error) {
    logger := GetLogger("sqlite")
    dbPath := "./db/main.db"

    dbExists, err := DBExists(dbPath)
    if err != nil {
        return nil, fmt.Errorf("error checking if database exists: %v", err)
    }

    if !dbExists {
        logger.Info("database file not found, creating...")
        if err := createDB(dbPath); err != nil {
            return nil, fmt.Errorf("error creating database: %v", err)
        }
    }

    db, err := OpenDBConnection(dbPath)
    if err != nil {
        return nil, fmt.Errorf("error opening database connection: %v", err)
    }

    logger.Info("applying migrations...")
    if err := applyMigrations(db); err != nil {
        return nil, fmt.Errorf("error applying migrations: %v", err)
    }

    return db, nil
}

func OpenDBConnection(dbPath string) (*gorm.DB, error) {
    db, err := gorm.Open(sqlite.Open(dbPath), &gorm.Config{})
    if err != nil {
        return nil, fmt.Errorf("error opening database connection: %v", err)
    }
    return db, nil
}

func DBExists(dbPath string) (bool, error) {
    _, err := os.Stat(dbPath)
    if err != nil {
        if os.IsNotExist(err) {
            return false, nil
        }
        return false, fmt.Errorf("error checking database existence: %v", err)
    }
    return true, nil
}

func createDB(dbPath string) error {
    if err := os.MkdirAll("./db", os.ModePerm); err != nil {
        return fmt.Errorf("error creating database directory: %v", err)
    }

    file, err := os.Create(dbPath)
    if err != nil {
        return fmt.Errorf("error creating database file: %v", err)
    }
    defer file.Close()

    return nil
}

func applyMigrations(db *gorm.DB) error {
    if err := db.AutoMigrate(&schemas.Opening{}); err != nil {
        return fmt.Errorf("error applying migrations: %v", err)
    }
    return nil
}



/* func InitializeSQL() (*gorm.DB, error){
	logger := GetLogger("sqlite")
	db_path := "./db/main.db"
	_, err := os.Stat(db_path)

	if os.IsNotExist(err){
		logger.Info("database file not found, creating...")
		err = os.MkdirAll("./db", os.ModePerm)

		if err != nil {
			return nil, err
		}

		file, err := os.Create(db_path)
		if err != nil {
			return nil, err
		}

		file.Close()
	}

	db, err := gorm.Open(sqlite.Open(db_path), &gorm.Config{})
	if err != nil {
		logger.Error("sqlite opening error: %v", )
		return nil, err
	}

	err = db.AutoMigrate(&schemas.Opening{})
	if err != nil {
		logger.Error("slite automigration error: %v", err)
		return nil, err
	}

	return db, nil
}
 */