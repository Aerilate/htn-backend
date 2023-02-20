# Hack the North Backend
Submission for the Hack the North 2023 Backend Challenge.

## Overview
This Go project compiles to a single binary that is capable of:
1. running migration scripts on a SQLite database
2. populating SQLite with mock data
3. serving user data through a REST API

The main libraries used are:
* cobra: a library for handling CLI arguments and flags
* gin: a web framework for routing HTTP requests
* golang-migrate: a library to run migration scripts from a specified directory
* gorm: an ORM to interact with SQLite

## Project Structure
* cmd: the main initial logic is found in the `migrate.go`, `populate.go`, and `serve.go` files
* migration: contains migration script
* model: defines Go structs for unmarshalling JSON and for the ORM
* repository: contains SQLite query logic
* main.go: starting point

## Setup and Installation
Requirements:
* Git
* Docker
(Go is not required)

### Clone
```bash
$ git clone git@github.com:Aerilate/htn-backend.git
```

### Basic Usage
```bash
$ docker compose up

$ curl localhost:8080/ping
pong
```

### Running Tests
```bash
# this builds the first stage only
$ docker build . --target builder
...
=> => writing image sha256:2de14607582f75261f0580ac906e3fdc9675451fbdfc29745b673163aebf0dad       0.0s
,,,

# run the image from the previous step
$ docker run sha256:2de146
?       github.com/Aerilate/htn-backend [no test files]
ok      github.com/Aerilate/htn-backend/cmd     0.010s
?       github.com/Aerilate/htn-backend/model   [no test files]
?       github.com/Aerilate/htn-backend/repository      [no test files]
```
