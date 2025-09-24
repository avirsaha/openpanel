package db

import (
    "database/sql"
    "fmt"
    "github.com/go-sql-driver/mysql"

    "gopanel/config"
)

var DB *sql.DB

// Init initializes the database connection
func Init(cfg config.DatabaseConfig) error {
    dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?parseTime=true",
        cfg.User,
        cfg.Password,
        cfg.Host,
        cfg.Port,
        cfg.Name,
    )

    db, err := sql.Open(cfg.Driver, dsn)
    if err != nil {
        return fmt.Errorf("failed to open DB connection: %v", err)
    }

    // Ping to verify connection
    if err := db.Ping(); err != nil {
        return fmt.Errorf("failed to ping DB: %v", err)
    }

    DB = db
    return nil
}

// Close closes the DB connection
func Close() error {
    if DB != nil {
        return DB.Close()
    }
    return nil
}
