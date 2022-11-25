package repository

import (
	"database/sql"
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/joho/godotenv"
  osx "ms.roomap.jp/os"
	_ "github.com/go-sql-driver/mysql"
)

var Db *sql.DB

func init() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println(err.Error())
	}

  dbn := osx.Getenv("DB_NAME", "")
  dbu := osx.Getenv("DB_USERNAME", "")
  dbpw := osx.Getenv("DB_PASSWORD", "")
  dbh := osx.Getenv("DB_HOST", "")
  dbpt := osx.Getenv("DB_PORT", "3306")
  dbch := osx.Getenv("DB_CHARSET", "utf8mb4")

	dbmll := osx.Getenv("DB_MAX_LIFE_TIME", "25")
	dbmoc := osx.Getenv("DB_MAX_IDLE_CONNS", "25")
	dbmic := osx.Getenv("DB_MAX_OPEN_CONNS", "5")

	dbconf := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=%s",
    dbu, dbpw, dbh, dbpt, dbch, dbn)

	Db, err = sql.Open("mysql", dbconf)
	// SetConnMaxLifetime sets the maximum amount of time a connection may be reused.
	// Expired connections may be closed lazily before reuse.
	// If d <= 0, connections are not closed due to a connection's age
	mll, _ := strconv.Atoi(dbmll)
	Db.SetConnMaxLifetime(time.Duration(mll) * time.Minute)
	if err != nil {
		fmt.Println(err)
	}
	// SetMaxIdleConns sets the maximum number of connections in the idle connection pool.
	// If MaxOpenConns is greater than 0 but less than the new MaxIdleConns, then the new MaxIdleConns will be reduced to match the MaxOpenConns limit.
	// If n <= 0, no idle connections are retained.
	// The default max idle connections is currently 2. This may change in a future release.
	moc, _ := strconv.Atoi(dbmic)
	Db.SetMaxIdleConns(moc)

	// SetMaxOpenConns sets the maximum number of open connections to the database.
	// If MaxIdleConns is greater than 0 and the new MaxOpenConns is less than MaxIdleConns, then MaxIdleConns will be reduced to match the new MaxOpenConns limit.
	// If n <= 0, then there is no limit on the number of open connections. The default is 0 (unlimited).
	mic, _ := strconv.Atoi(dbmoc)
	Db.SetMaxOpenConns(mic)

	if err := Db.Ping(); err != nil {
		fmt.Println(err)
	}

	fmt.Printf("db<%s> connected", dbn)
}
