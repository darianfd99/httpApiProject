package bootstrap

import (
	"database/sql"
	"fmt"

	server "github.com/darianfd99/httpApiProject/internal/platform/server"
)

const (
	host = "localhost"
	port = 8080

	dbUser = "darian"
	dbPass = "data"
	dbHost = "localhost"
	dbPort = "3306"
	dbName = "apiGo"
)

func Run() error {
	mysqlURI := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", dbUser, dbPass, dbHost, dbPort, dbName)
	db, err := sql.Open("mysql", mysqlURI)
	if err != nil {
		return err
	}

	courseRepository := mysql.NewCourseRepository(db)

	srv := server.New(host, port, courseRepository)
	return srv.Run()
}
