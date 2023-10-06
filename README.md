# Project structure

The project tries to follow the guidelines described in [this](https://github.com/golang-standards/project-layout) repository.

## `cmd`

Contains commands, including the command for running the server.

## `internal`

Contains code not to be imported by anyone else.

## `config`

Contains configs.

# How to use Makefile

First of, create a `.env` file. You can leave it empty though.

There is a set of commands for running the server. The commands are listed in the Makefile. Using Makefile you can omit path to config file flag.

`make` runs `make start` which reads `.env` for environment variables. If there is a `DIGITAL_LIBRARY_CONFIG` variable specified, the server will try 

# How to specify the path to config if you don't want to use Makefile

The options are listed in order of priority.
- `-config` flag
- export the environment variable `DIGITAL_LIBRARY_CONFIG`

# How to use REST API

I recommend using `curl`.

Send the following requests to manage users:

`curl -X GET "http://localhost:PORT/user/ID"` - get the user with id of `ID` while server is running on `PORT` port

`curl -X PUT "http://localhost:PORT/user` -H "Content-Type: application/json" -d '{"id": ID, "first_name": FIRST_NAME, "second_name": SECOND_NAME}' - add a user with id of `ID` (value of type int, do not put quotes), first name of `FIRST_NAME` (value of type string, do put quotes), second name of `SECOND_NAME` (value of type string, do put quotes) while server is running on `PORT` port

`curl -X DELETE "http://localhost:PORT/user/ID"` - delete the user with id of `ID` while server is running on `PORT` port
