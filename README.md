# Finite Mock Server

The public facing API through which connectors are exposed as a single abstract API.

## Overview

This server was generated by the [openapi-generator](https://openapi-generator.tech) project.

- API version: v1.5
- Build date: 2021-03-31T17:40:06.545270+01:00[Europe/Lisbon]

For more information, please visit [https://www.trexis.net](https://www.trexis.net)

## Prerequisites

* [Docker Desktop](https://www.docker.com/products/docker-desktop)
* [Homebrew](https://docs.brew.sh/Installation)
* [Taskfile](https://taskfile.dev) 3.x - `brew install go-task/tap/go-task`
* [OpenAPI Generator](https://github.com/OpenAPITools/openapi-generator) 5.1.x - `brew install openapi-generator`
* [Golang](https://golang.org/) 1.16.x - `brew install go`
* [ent.](https://entgo.io/) - `go get entgo.io/ent/cmd/ent`

## Configuration

All environment specific configuration is defined in the file `.env`.
## Generating / updating specs

~~~bash
# (Re-)generate API definitions, API models
$ task generate

# Bring the database up (postgres or mysql)
$ docker-compose up -d postgres
~~~
## Running the server

To run the server, follow these simple steps:

~~~bash
$ go run main.go
~~~

To run the server in a docker container:

~~~bash
$ docker-compose build mock
~~~

Once image is built use:

~~~bash
# Configure the DB_VENDOR and DB_HOST settings in docker-compose
$ docker-compose up mock 
~~~
