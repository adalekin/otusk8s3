// This is custom goose binary with sqlite3 support only.

package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq"
	"github.com/pressly/goose"
)

var (
	flags      = flag.NewFlagSet("otusk8s3-migrations", flag.ExitOnError)
	dir        = flags.String("dir", ".", "directory with migration files")
	verbose    = flags.Bool("v", false, "enable verbose mode")
	help       = flags.Bool("h", false, "print help")
	version    = flags.Bool("version", false, "print version")
	sequential = flags.Bool("s", false, "use sequential numbering for new migrations")
)

func main() {
	flags.Parse(os.Args[1:])

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

	switch args[1] {
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

	if len(args) < 2 {
		flags.Usage()
		return
	}

	dbstring, command := args[0], args[1]

	db, err := goose.OpenDBWithDriver("postgres", dbstring)
	if err != nil {
		log.Fatalf("goose: failed to open DB: %v\n", err)
	}

	defer func() {
		if err := db.Close(); err != nil {
			log.Fatalf("goose: failed to close DB: %v\n", err)
		}
	}()

	arguments := []string{}
	if len(args) > 3 {
		arguments = append(arguments, args[3:]...)
	}

	if err := goose.Run(command, db, *dir, arguments...); err != nil {
		log.Fatalf("goose %v: %v", command, err)
	}
}
