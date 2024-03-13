<div id="top"></div>


<!-- PROJECT LOGO -->
<br />
<div align="center">
  <a href="https://github.com/othneildrew/Best-README-Template">
    <img src="https://seeklogo.com/images/G/go-logo-046185B647-seeklogo.com.png" alt="Logo" width="70">
  </a>

<h3 align="center">Go Base</h3>
<p>Building a Web API with Clean Architecture, GIN, and GORM</p>
</div>

## Prerequisites
Before getting started, ensure you have the following dependencies installed:
- Redis
  - Install Redis based on your operating system.
- Golang Migrate Tool (global installation):
  - Install the Golang Migrate tool globally by running:
  ```
  go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest
  ```

## Quick Start Guide
 - Clone the Project:
    ``` 
    git clone https://github.com/yousifnimah/Go-Base.git
    ```
- Navigate to the Project Directory: ```cd Go-Base```

- Create Environment File:
  - Create a .env file in the project directory.
  - You can refer to .env.example for the required properties.
- Update Go Modules: ```go mod tidy```

- Run the Application: ```go run .```

That's it! You're ready to start working with the project.

<br />

## Database Migrations
Inside the "DB/migrations" directory, you will find all the database migration files. These files contain SQL scripts that represent changes to the database schema over time. Each migration file corresponds to a specific database modification, allowing for seamless tracking and versioning of database changes.

### Create Migration File 
Use the following command to create a migration file: 

```
migrate create -ext sql -dir DB/migrations -seq create_users_table
```


### Migrate Up
To apply migrations and update the database schema, run:
```
go run DB/migrate.go up
```

### Migrate Down
To roll back migrations, use the following command:
```
go run DB/migrate.go down
```

### Seeds
Run the seeder to populate the database with initial data:
```
go run DB/seeder.go
```

<br/>


## Packages Used
The project utilizes the following Go packages:

- Gin Web Framework:
  - GitHub Repository: github.com/gin-gonic/gin

- bcrypt (golang.org/x/crypto):
  Provides functions for hashing and comparing passwords securely.

- Golang Migrate: Used for managing database migrations.
  - GitHub Repository: github.com/golang-migrate/migrate

- GORM: An ORM (Object-Relational Mapping) library for Go, used for database operations.
  - Official Website: gorm.io
- JWT (JSON Web Tokens): Used for generating and verifying JSON Web Tokens for authentication and authorization.
  - GitHub Repository: github.com/golang-jwt/jwt
