// This is custom goose binary with sqlite3 support only.

package main

import (
	"database/sql"
	"flag"
	"fmt"
	"os"
	"time"

	"github.com/adalekin/otusk8s3/internal/appenv"
	"github.com/caarlos0/env"
	_ "github.com/lib/pq"
	"github.com/pressly/goose"
	log "github.com/sirupsen/logrus"
)

var (
	flags                = flag.NewFlagSet("otusk8s3-migrations", flag.ExitOnError)
	dir                  = flags.String("dir", ".", "directory with migration files")
	verbose              = flags.Bool("v", false, "enable verbose mode")
	help                 = flags.Bool("h", false, "print help")
	version              = flags.Bool("version", false, "print version")
	sequential           = flags.Bool("s", false, "use sequential numbering for new migrations")
	maxReconnectAttempts = flags.Int("r", 10, "number of database reconnect attempts")
)

func main() {
	_ = flags.Parse(os.Args[1:])

	if *version {
		fmt.Println(goose.VERSION)
		return
	}
	if *verbose {
		goose.SetVerbose(true)
	}
	if *sequential {
		goose.SetSequential(true)
	}

	args := flags.Args()
	if len(args) == 0 || *help {
		flags.Usage()
		return
	}

	switch args[0] {
	case "create":
		if err := goose.Run("create", nil, *dir, args[2:]...); err != nil {
			log.Fatalf("goose run: %v", err)
		}
		return
	case "fix":
		if err := goose.Run("fix", nil, *dir); err != nil {
			log.Fatalf("goose run: %v", err)
		}
		return
	}

	// Parse environment variables
	config := appenv.Config{}
	if err := env.Parse(&config); err != nil {
		fmt.Printf("%+v\n", err)
	}

	dbConnectionString := fmt.Sprintf("host=%s user=%s password=%s dbname=%s sslmode=disable", config.DbHost, config.DbUser, config.DbPassword, config.DbName)

	command := args[0]

	var db *sql.DB
	var err error

	for databaseConnectionAttemptLoop := 0; databaseConnectionAttemptLoop < *maxReconnectAttempts; databaseConnectionAttemptLoop++ {
		log.Infof("Trying to connect to DB: attempt %d", databaseConnectionAttemptLoop+1)

		if db, err = goose.OpenDBWithDriver("postgres", dbConnectionString); err == nil {
			err = db.Ping()

			if err == nil {
				break
			}
		}
		time.Sleep(1 * time.Duration(databaseConnectionAttemptLoop) * time.Second)
	}

	if err != nil {
		log.Fatalf("migrations: failed to open DB: %v", err)
		return
	}

	defer func() {
		if err := db.Close(); err != nil {
			log.Fatalf("migrations: failed to close DB: %v", err)
		}
	}()

	arguments := []string{}
	if len(args) > 2 {
		arguments = append(arguments, args[2:]...)
	}

	if err := goose.Run(command, db, *dir, arguments...); err != nil {
		log.Fatalf("migrations %v: %v", command, err)
	}
}
