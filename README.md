# WIP

The project is still work in progress.

# Domain

## Specification

### Points of view

#### User

User should be able to:
- view and download books
- read, create, edit and delete their own reviews
- read other people reviews
- mark some books as favorite
- view other people profiles

#### Mod

Mod should be able to:
- do all the things the User can
- delete other people reviews

#### Admin

Admin should be able to:
- do all the things Mod can
- upload and delete books
- create and delete users
- edit user roles from User to Mod and from Mod to User

### Restrictions

#### Book file extension

Book file extension should be PDF.

### Stack

Technology stack should be:
- frontend: ...
- backend: Go
- DBMS: SQLite or MySQL

### Work flow

Both frontend and backend should run locally. If MySQL is used, it should run in a Docker container. There should be an option to specify the environment - local, development or production.

## Entity-Relation Diagram

![ERD](docs/images/ERD.svg)

# Project structure

The project tries to follow the guidelines described in [this](https://github.com/golang-standards/project-layout) repository.

## `cmd`

Contains commands, including the command for running the server.

## `config`

Contains configs.

## `docs`

Contains files that describe the project and are not code.

## `internal`

Contains code not to be imported by anyone else.

# How to use `Makefile`

First of, create a `.env` file. You can leave it empty though.

There is a set of commands for running the server. The commands are listed in the `Makefile`.

`make` runs `make start` which reads `.env` for environment variables. If there is a `DIGITAL_LIBRARY_CONFIG` variable specified, the server will try to open, parse and use the config on specified path.

`make local` sets config to `local.yaml` automatically. 

# How to specify the path to config if you don't want to use `Makefile`

The options are listed in order of priority:
- `-config` flag
- export the environment variable `DIGITAL_LIBRARY_CONFIG`

# How to use REST API

## Using Swagger UI

After running the app, there will be an endpoint (currently `PROTO://HOST:PORT/api`) that will return an [Swagger UI](https://swagger.io/tools/swagger-ui/) page describing the REST API. Note that not all the endpoints described in the Swagger UI are implemented yet.

## Using `curl`

Here are some examples of using REST API via `curl`:

`curl -X GET "http://localhost:PORT/user/ID"` - get the user with id of `ID` while server is running on `PORT` port.

`curl -X PUT "http://localhost:PORT/user` -H "Content-Type: application/json" -d '{"id": ID, "first_name": FIRST_NAME, "second_name": SECOND_NAME}' - add a user with id of `ID` (value of type int, do not put quotes), first name of `FIRST_NAME` (value of type string, do put quotes), second name of `SECOND_NAME` (value of type string, do put quotes) while server is running on `PORT` port.

`curl -X DELETE "http://localhost:PORT/user/ID"` - delete the user with id of `ID` while server is running on `PORT` port.

# How to create a database

The instructions are Fedora-specific, but the process itself should be the same on all Linux distros.

## SQLite

### Install SQLite

`
sudo dnf -y install sqlite
`

## Create SQLite database

`
mkdir .storage && sqlite3 .storage/storage.db
`

## MySQL

I recommend to install it as a Docker container so it's isolated from your main operating system.

### Install Docker 

[Source](https://docs.docker.com/engine/install/fedora/)

#### Set up repository

`
sudo dnf -y install dnf-plugins-core
sudo dnf config-manager --add-repo https://download.docker.com/linux/fedora/docker-ce.repo
`

#### Install Docker Engine

`
sudo dnf install docker-ce docker-ce-cli containerd.io docker-buildx-plugin docker-compose-plugin
`

#### Start Docker

`
sudo systemctl start docker
`

### Install MySQL

[Source](https://earthly.dev/blog/docker-mysql/)

#### Start a MySQL server instance

`
sudo docker run --tty --name digital-library-mysql -p 3306:3306 -e MYSQL_DATABASE=digital-library -e MYSQL_ROOT_PASSWORD=digital-library --restart unless-stopped -v mysql:/var/lib/digital-library mysql:8
`

This will create and start a Docker container named `digital-library-mysql` from `mysql:8` image on port `3306` that creates `digital-library` MySQL database with `digital-library` password for root user that will store it's data on `mysql` volume in `/var/lib/digital-library` directory and will run even after system restart unless the container gets stopped manually.

You can also set up a container network and configure MySQL more thoroughly.

