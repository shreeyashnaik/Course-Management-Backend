# Course Management System Backend

## About the project

This repository consists of Course Management System Backend code. It follows clean architecture.

## Postman docs

[Postman](https://documenter.getpostman.com/view/18647190/VVdf5kWQ)

### Tree Layout

```tree
.
├── .gitignore
├── dockercompose.yml
├── Dockerfile
├── go.mod
├── go.sum
├── LICENSE
├── main.go
├── README.md
├── sample.env
├── common
│   ├── db
│   │   ├── db.go
│   │   └── services.go
│   ├── migrations
│   │   └── migrations.go
│   ├── schemas
│   │   ├── schemas.go
│   │   └── validations.go
│   ├── utils
│   │   ├── jwt_utils.go
│   │   ├── list_utils.go
│   │   ├── validator_utils.go
│   │   └── viper_utils.go
│   └── views
│       └── views.go
├── config
│   └── config.go
├── pkg
│   ├── models
│   │   ├── course.go
│   │   └── user.go
│   └── services
│       ├── course
│       │   ├── postgres.go
│       │   └── service.go
│       └── user
│           ├── postgres.go
│           └── service.go
├── README.md
└── src
    └── core
        ├── controllers
        │   └── controllers.go
        ├── middlewares
        │   └── middlewares.go
        ├── routers
        │   └── routers.go
        └── server
            └── server.go
```

Overview of the layout:

* `.gitignore` contains all files to be ignored for github.
* `README.md` is a detailed description of the project.
* `src` contains all micro services.
* `pkg` comprises business logic.
* `common` consists of common code, important to the application

### Steps to run

Before running make sure to add `.env` file. Refer `sample.env`

`go run main.go`