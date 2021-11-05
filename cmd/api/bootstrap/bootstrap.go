package bootstrap

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	creating "github.com/darianfd99/httpApiProject/internal/creating"
	"github.com/darianfd99/httpApiProject/internal/platform/bus/inmemory"
	server "github.com/darianfd99/httpApiProject/internal/platform/server"
	"github.com/darianfd99/httpApiProject/internal/platform/storage/mysql"
	_ "github.com/go-sql-driver/mysql"
)

const (
	host            = "localhost"
	port            = 8080
	shutdownTimeout = 10 * time.Second

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

	ctx, srv := server.New(context.Background(), host, port, shutdownTimeout, commandBus)
	return srv.Run(ctx)
}
