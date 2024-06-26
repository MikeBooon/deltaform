package main

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/mikebooon/deltaform/internal/db/migrate"
	"github.com/mikebooon/deltaform/internal/email"
	"github.com/mikebooon/deltaform/internal/rest"
	"github.com/mikebooon/deltaform/internal/util"
	"github.com/mikebooon/deltaform/service"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

const (
	defaultAddress = ":3090"
)

func main() {
	log.Println("DeltaForm Starting...")

	if os.Getenv("APP_ENV") != "production" {
		err := godotenv.Load()

		if err != nil {
			log.Fatal("Failed to load .env file")
		}
	}

	e := echo.New()
	e.HideBanner = true

	e.Validator = util.NewValidator()
	e.Use(middleware.CORS())
	e.Use(middleware.Recover())

	e.Use(middleware.RequestLoggerWithConfig(middleware.RequestLoggerConfig{
		LogStatus:  true,
		LogURI:     true,
		LogMethod:  true,
		LogLatency: true,
		LogValuesFunc: func(c echo.Context, v middleware.RequestLoggerValues) error {
			log.Printf("REQ: %v | %v | %v | %v\n", v.Method, v.URI, v.Status, v.Latency)
			return nil
		},
	}))

	dbConnection := getDbConnection()

	db, err := gorm.Open(postgres.Open(dbConnection), &gorm.Config{})

	if err != nil {
		log.Fatal("Failed to open connection to database")
		os.Exit(1)
	}

	migrate.RunMigration(db)

	emailClient := email.NewEmailClient(email.EmailOptions{
		Sender:    os.Getenv("EMAIL_SENDER_ADDRESS"),
		AccessKey: os.Getenv("AWS_ACCESS_KEY"),
		SecretKey: os.Getenv("AWS_SECRET_KEY"),
	})

	serviceRepo := service.NewServiceRepo(*db)

	rest.NewFormHandler(e, *serviceRepo)
	rest.NewAuthHandler(e, *serviceRepo, *emailClient)
	rest.NewUserHandler(e, *serviceRepo)

	address := os.Getenv("SERVER_ADDRESS")

	if address == "" {
		log.Println("SERVER_ADDRESS not specified defaulting to " + defaultAddress)
		address = defaultAddress
	}

	e.Start(address)
}

func getDbConnection() string {
	dbHost := os.Getenv("DB_HOST")
	dbUser := os.Getenv("DB_USER")
	dbPort := os.Getenv("DB_PORT")
	dbPass := os.Getenv("DB_PASS")
	dbName := os.Getenv("DB_NAME")
	dbSslMode := os.Getenv("DB_SSL_MODE")

	return fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=GMT",
		dbHost,
		dbUser,
		dbPass,
		dbName,
		dbPort,
		dbSslMode,
	)
}
