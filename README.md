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
... # some Docker build output
=> => writing image sha256:2de14607582f75261f0580ac906e3fdc9675451fbdfc29745b673163aebf0dad       0.0s
,,,

# run the image from the previous step
$ docker run sha256:2de146
?       github.com/Aerilate/htn-backend [no test files]
ok      github.com/Aerilate/htn-backend/cmd     0.010s
?       github.com/Aerilate/htn-backend/model   [no test files]
?       github.com/Aerilate/htn-backend/repository      [no test files]
```

Alternatively, in one command:
```bash
# please be patient!
$ docker run $(docker build . --target builder -q)
?       github.com/Aerilate/htn-backend [no test files]
ok      github.com/Aerilate/htn-backend/cmd     0.010s
?       github.com/Aerilate/htn-backend/model   [no test files]
?       github.com/Aerilate/htn-backend/repository      [no test files]
```


## API
### GET /users/
Retrieves information on all users in the database.

#### Return
* On failure, returns 404.
* On success, returns 200 with a body of the following schema:
```json
[
    {
        "name": <string>,
        "company": <string>,
        "email": <string>,
        "phone": <string>,
        "skills": [
            {
                "skill": <string>,
                "rating": <int>
            }
        ]
    },
    ...
]
```

#### Example
```bash
$ curl localhost:8080/users/
[{"name":"Breanna Dillon","company":"Jackson Ltd","email":"lorettabrown@example.net","phone":"+1-924-116-7963","skills":[{"skill":"OpenCV","rating":1},{"skill":"Swift","rating":4}]},{"name":"Kimberly Wilkinson","company":"Moon, Mendoza and Carter","email":"frederickkyle@example.org","phone":"(186)579-0542","skills":[{"skill":"Elixir","rating":4},{"skill":"Fortran","rating":2},{"skill":"Foundation","rating":4},{"skill":"Plotly","rating":3}]},
...
```

### GET /users/<id>
Retrieve information on the user with the given id.

#### Return
* On invalid id argument (i.e. not a number), returns 400.
* On other failure, returns 404.
* On success, returns 200 with a body of the following schema:
```json
{
    "name": <string>,
    "company": <string>,
    "email": <string>,
    "phone": <string>,
    "skills": [
        {
            "skill": <string>,
            "rating": <int>
        }
    ]
}
```

### Example
```bash
$ curl localhost:8080/users/1
{"name":"Breanna Dillon","company":"Jackson Ltd","email":"lorettabrown@example.net","phone":"+1-924-116-7963","skills":[{"skill":"OpenCV","rating":1},{"skill":"Swift","rating":4}]}
```

## PUT /users/<id>
Update some information about a user.

#### Body
Provide a JSON in the request body with some or all of the following keys:
```json
{
    "name": <string>,
    "company": <string>,
    "email": <string>,
    "phone": <string>,
    "skills": [
        {
            "skill": <string>,
            "rating": <int>
        },
        ...
    ]
}
```
Keys omitted in the request body will not have their values updated.
If the "skills" key is provided in the request body, the skills in the array will entirely replace the user's original skills. User skills that are not in the array will be removed.

#### Return
* On failure, returns 400.
* On success, returns 200.

#### Example
```bash
# we'll be updating the user with ID=1
$ curl localhost:8080/users/1
{"name":"Breanna Dillon","company":"Jackson Ltd","email":"lorettabrown@example.net","phone":"+1-924-116-7963","skills":[{"skill":"OpenCV","rating":1},{"skill":"Swift","rating":4}]}

# update some fields
$ curl -X PUT -H "Content-Type: application/json" -d '{"company":"", "email":"asdf@example.net", "phone":"123"}' localhost:8080/users/1

# get updated user information
$ curl localhost:8080/users/1
{"name":"Breanna Dillon","company":"","email":"asdf@example.net","phone":"123","skills":[{"skill":"OpenCV","rating":1},{"skill":"Swift","rating":4}]}

# now let's also update some skills
$ curl -X PUT -H "Content-Type: application/json" -d '{"email":"lorettabrown@example.net","skills":[{"skill":"OpenCV","rating":5},{"skill":"Python","rating":3}]}' localhost:8080/users/1

# get updated user information
$ curl localhost:8080/users/1
{"name":"Breanna Dillon","company":"","email":"lorettabrown@example.net","phone":"123","skills":[{"skill":"OpenCV","rating":5},{"skill":"Python","rating":3}]}
```

### GET /skills/
Retrieve aggregate information on user skills.

#### Optional Query Parameters
* min_frequency: only include skills that are possessed by at least min_frequency users (inclusive).
* max_frequency: only include skills that are possessed by at most min_frequency users (inclusive).

#### Return
* On invalid query parameter values (i.e. not a number), returns 400.
* On other failure, returns 404.
* On success, returns 200 with a body of the following schema:
```json
[
    {
        "Skill": <string>,
        "Count": <int>
    },
    ...
]
```

### Example
```bash
$ curl "localhost:8080/skills/"
[{"skill":"Sanic","count":43},{"skill":"React","count":41},{"skill":"Plotly","count":39}, ...

$ curl "localhost:8080/skills/?min_frequency=40"
[{"skill":"Sanic","count":43},{"skill":"React","count":41}]

$ curl "localhost:8080/skills/?min_frequency=19&max_frequency=21"
[{"skill":"Matplotlib","count":21},{"skill":"Aurelia","count":21},{"skill":"Starlette","count":20},{"skill":"Pascal","count":20},{"skill":"Numpy","count":20},{"skill":"Lisp","count":20},{"skill":"Tachyons","count":19}]
```
