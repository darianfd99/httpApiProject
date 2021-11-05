package bootstrap

import (
	"database/sql"
	"fmt"

	creating "github.com/darianfd99/httpApiProject/internal/creating"
	"github.com/darianfd99/httpApiProject/internal/platform/bus/inmemory"
	server "github.com/darianfd99/httpApiProject/internal/platform/server"
	"github.com/darianfd99/httpApiProject/internal/platform/storage/mysql"
	_ "github.com/go-sql-driver/mysql"
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

	var (
		commandBus = inmemory.NewCommandBus()
	)

	courseRepository := mysql.NewCourseRepository(db)
	creatingCourseService := creating.NewCourseService(courseRepository)

	createCourseCommandHandler := creating.NewCourseCommandHandler(creatingCourseService)
	commandBus.Register(creating.CourseCommandType, createCourseCommandHandler)

	srv := server.New(host, port, commandBus)
	return srv.Run()
}
