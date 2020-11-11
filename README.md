![build](https://github.com/Whyeasy/zally-cleaner/workflows/build/badge.svg)
![status-badge](https://goreportcard.com/badge/github.com/Whyeasy/zally-cleaner)
![Github go.mod Go version](https://img.shields.io/github/go-mod/go-version/Whyeasy/zally-cleaner)

# zally-cleaner

A Go app to run as a job to clean up old results from Zally Server and if it's backed with a Postgres Database.

## Requirements

Provide the Postgres host; `--postgresHost` or as env variable `POSTGRES_HOST`.

Provide the Postgres database name: `--postgresDatabase` or as env variable `POSTGRES_DB`.

Provide the Postgres user: `--postgresUser` or as env variable `POSTGRES_USER`.

Provide the Postgres user password: `--postgresPassword` or as env variable `POSTGRES_PASSWORD`.

### Optional

Change the SSL connection mode to Postgres; `--postgresSSL` or as env Variable `POSTGRES_SSL`. The default value is `require`.

Change the retention period of the records to keep. `--zallyRetention <string>` or as env variable `ZALLY_RETENTION`. The default value is `"7"`. Provide value as a string and all entries older than the value will be deleted!

Change if you don't use the Zally violations metrics output. `--zallyMetrics <string>` or as env variable `ZALLY_METRICS`. The value is `"true"`. Provide value as a string.
