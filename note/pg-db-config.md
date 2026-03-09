# Setup Portgres Database Connection

1. Install Necessary Package

postgres db connection
```bash
go get github.com/lib/pq
```

sqlx db connection
```bash
go get github.com/jmoiron/sqlx
```

2. Create Database Connection

```go
package config

import (
	"fmt"
	"os"

	"github.com/gofiber/fiber/v3/log"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func StartDatabase() *sqlx.DB {
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")

	pgsqlInfo := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	db, err := sqlx.Connect("postgres", pgsqlInfo)
	if err != nil {
		log.Fatal("cannot connect to database: ", err)
	}

	log.Info("database connected")

	return db
}

```

3. Call Database Connection in main.go

```go
package main

import (
	"oat431/go-fiber-snippets-vol2/internal/config"
	"oat431/go-fiber-snippets-vol2/internal/routes"
	"os"

	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/log"
)

func main() {
	config.LoadEnvConfig()
	
	// call database connection
	db := config.StartDatabase()
	defer db.Close()
	// end of database connection

	app := fiber.New()
	routes.SetupRoutes(app)

	port := os.Getenv("PORT")
	err := app.Listen(":" + port)
	if err != nil {
		log.Fatal("port :+ " + port + " is already in use")
	}
}
```

