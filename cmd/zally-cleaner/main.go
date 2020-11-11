package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"

	log "github.com/sirupsen/logrus"
	"github.com/whyeasy/zally-cleaner/internal"
	"github.com/whyeasy/zally-cleaner/lib/cleaner"
	"github.com/whyeasy/zally-cleaner/lib/client"
)

var (
	config internal.Config
)

func init() {
	flag.StringVar(&config.User, "postgresUser", os.Getenv("POSTGRES_USER"), "User to use for connecting to the Postgres Database")
	flag.StringVar(&config.Password, "postgresPassword", os.Getenv("POSTGRES_PASSWORD"), "Password to use for connecting to the Postgres Database")
	flag.StringVar(&config.Host, "postgresHost", os.Getenv("POSTGRES_HOST"), "Host address of the Postgres Instance")
	flag.StringVar(&config.Database, "postgresDatabase", os.Getenv("POSTGRES_DB"), "Database to select within the Postgres Instance")
	flag.StringVar(&config.Port, "postgresPort", os.Getenv("POSTGRES_PORT"), "Port to use to connect to Postgres Instance")
	flag.StringVar(&config.SSLMode, "postgresSSL", os.Getenv("POSTGRES_SSL"), "SSL mode to use to connect to the Postgres Instance")
	flag.StringVar(&config.Retention, "zallyRetention", os.Getenv("ZALLY_RETENTION"), "Provide retention period of records in Days")
	flag.StringVar(&config.Metrics, "zallyMetrics", os.Getenv("ZALLY_METRICS"), "By Default true, so it will clean the custom_label_mapping table.")
}

func main() {
	if err := parseConfig(); err != nil {
		log.Error(err)
		flag.Usage()
		os.Exit(2)
	}
	log.Info("Running Zally cleaner")

	dbClient := client.New(config)

	retention, _ := strconv.Atoi(config.Retention)
	metrics, _ := strconv.ParseBool(config.Metrics)
	cleaner.CleanUp(dbClient, retention, metrics)

}

func parseConfig() error {
	flag.Parse()
	required := []string{"postgresUser", "postgresPassword", "postgresHost", "postgresDatabase"}
	var err error
	flag.VisitAll(func(f *flag.Flag) {
		for _, r := range required {
			if r == f.Name && (f.Value.String() == "" || f.Value.String() == "0") {
				err = fmt.Errorf("%v is empty", f.Usage)
			}
		}
		if f.Name == "postgresSSL" && (f.Value.String() == "" || f.Value.String() == "0") {
			err = f.Value.Set("require")
			if err != nil {
				log.Error(err)
			}
		}
		if f.Name == "postgresPort" && (f.Value.String() == "" || f.Value.String() == "0") {
			err = f.Value.Set("5432")
			if err != nil {
				log.Error(err)
			}
		}
		if f.Name == "zallyRetention" && (f.Value.String() == "" || f.Value.String() == "0") {
			err = f.Value.Set("7")
			if err != nil {
				log.Error(err)
			}
		}
		if f.Name == "zallyMetrics" && (f.Value.String() == "" || f.Value.String() == "0") {
			err = f.Value.Set("true")
			if err != nil {
				log.Error(err)
			}
		}
	})

	return err
}
