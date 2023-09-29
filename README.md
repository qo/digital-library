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
