package main

import (
	"fmt"
	"log"
	"os"

	"github.com/dashroots/migrator"
	flags "github.com/jessevdk/go-flags"
)

const (
	databaseName   = "dashroots"
	databaseUser   = "dashroots"
	migrationsRepo = "github://dashroots/migrations"
)

var args struct {
	Host    string `long:"host" description:"database hostname" required:"true"`
	Port    uint   `long:"port" description:"database port" required:"true"`
	Pass    string `long:"pass" description:"database password" required:"true"`
	Version uint   `long:"version" description:"version of database to migrate to" required:"false"`
}

func printHeader() {
	fmt.Printf("DashRoots Database Migrator\n\n")
}

func parseFlags() {
	parser := flags.NewParser(&args, flags.None)
	_, err := parser.Parse()

	if err != nil {
		printHeader()
		parser.WriteHelp(os.Stdout)
		os.Exit(1)
	}
}

func buildConnString(host string, port uint, password string) string {
	format := "postgres://%s:%d/%s?sslmode=disable&user=%s&password=%s"
	return fmt.Sprintf(format, host, port, databaseName, databaseUser, password)
}

func main() {
	parseFlags()

	connection := buildConnString(args.Host, args.Port, args.Pass)
	version, err := migrator.Run(migrationsRepo, connection, args.Version)

	if err != nil {
		printHeader()
		log.Fatal(err)
	}

	printHeader()
	fmt.Printf("Database now at version: %d\n", version)
}
