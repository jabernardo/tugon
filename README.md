# Tugon

**Tugon** is a minimal API boilerplate built with Go's standard `net/http` package.

## Go Version

This is currently running on `go 1.23.0` (or later)

## Running

```
cp .env.example .env
go run .
```

## Generating Docs

This uses [swaggo](https://github.com/swaggo/swag) for generating Swagger API Documentation.
To generate docs, run:

```sh
swag init --parseDependency --parseInternal
```
