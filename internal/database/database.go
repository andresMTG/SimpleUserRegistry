package database

import (
	"database/sql"
    "log"
    "time"

    "andresMTG/SimpleUserRegistry/internal/config"
    _ "github.com/go-sql-driver/mysql"
)


func Connect(cfg *config.AppConfig) *sql.DB {
    db, err := sql.Open("mysql", cfg.DBConfig.FormatDSN())
    if err != nil {
        log.Fatal("Failed to open database configuration:", err)
    }

    var pingErr error
    for i := 1; i <= 5; i++ {
        pingErr = db.Ping()
        if pingErr == nil {
            log.Println("Connected!")
            return db 
        }
        time.Sleep(2 * time.Second)
    }

    log.Fatal("Error to connect to data base after retries:", pingErr)
    return nil
}
