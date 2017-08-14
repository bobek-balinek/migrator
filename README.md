# Migrator
This repo contains the `migrator` tool, which should be used to execute SQL migrations against the specified DashRoots database.

## Building
```
make build
```

This will create a binary named `migrate` in the local `bin` directory.

## Running
```
→ ./bin/migrate
DashRoots Database Migrator

Usage:
  migrate

Application Options:
  --host=      database hostname
  --port=      database port
  --pass=      database password
  --version=   version of database
```

The arguments `--host`, `--port` and `--pass` are all required.

Specify `--version` if you want to migrate the database to a specific version. Otherwise, the database will be brought up to the latest possible version.

## Migration Files
This tool executes the migrations found in the [dashroots/migrations](https://github.com/dashroots/migrations) repo.