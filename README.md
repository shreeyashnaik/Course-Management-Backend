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


## DB Schema Details

<strong>Table 1</strong>: `users`<br/>
All the users including `superadmin`, `admin`, `employee`.

<strong>Table 2</strong>: `courses`<br/>
All the courses designed by a `user (admin)`. Hence `user_id`
in courses will refer to `id` in `users`.

<strong>Table 3</strong>: `viewed_courses` (join table between `users` and `courses`)<br/>
This table is a result of many-to-many relationship between
users and courses. If the user U has viewed course C, we insert U & C as a
composite primary key in `viewed_courses`. If U has completed C, we mark
`is_completed` to be `true`. We get to know the course completion status by
`completed_duration` attribute in the table.

### <strong>DB Schema Diagram</strong>
![image](/db_schema.png)