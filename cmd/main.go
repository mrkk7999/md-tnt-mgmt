package main

import (
	"fmt"
	"md-tnt-mgmt/controller"
	"md-tnt-mgmt/implementation"
	"md-tnt-mgmt/repository"
	httpHandler "md-tnt-mgmt/transport/http"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {

	// Logger
	log := logrus.New()
	log.SetFormatter(&logrus.JSONFormatter{
		TimestampFormat: "2006-01-02 15:04:05",
	})

	err := godotenv.Load("../.env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	var (
		httpAddr   = os.Getenv("HTTP_ADDR")
		dbHost     = os.Getenv("DB_HOST")
		dbUser     = os.Getenv("DB_USER")
		dbPassword = os.Getenv("DB_PASSWORD")
		dbName     = os.Getenv("DB_NAME")
		dbPort     = os.Getenv("DB_PORT")
		dbSSLMode  = os.Getenv("DB_SSLMODE")
		dbTimeZone = os.Getenv("DB_TIMEZONE")
	)

	// PostgreSQL DSN
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=%s",
		dbHost, dbUser, dbPassword, dbName, dbPort, dbSSLMode, dbTimeZone)

	// Connect to Database
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	repo := repository.New(db)

	svc := implementation.New(repo)

	controller := controller.New(svc, log)

	handler := httpHandler.SetUpRouter(controller, log)

	errs := make(chan error)

	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
		errs <- fmt.Errorf("%s", <-c)
	}()

	log.Info("Server is running " + httpAddr)

	go func() {
		server := &http.Server{
			Addr:    httpAddr,
			Handler: handler,
		}
		errs <- server.ListenAndServe()
	}()

	log.Error("exit", <-errs)
}
