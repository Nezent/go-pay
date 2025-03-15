package config

// import (
// 	"context"
// 	"fmt"
// 	"log"
// 	"os"

// 	"github.com/jackc/pgx/v5"
// 	"github.com/joho/godotenv"
// )

// func ConnectDB() *pgx.Conn {

// 	err := godotenv.Load(".env")
// 	if err != nil {
// 		log.Fatalf("Error loading .env file")
// 	}
// 	// Create the connection string
// 	dsn := fmt.Sprintf("postgres://%s:%s@%s:%s/%s", os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_HOST"), os.Getenv("DB_PORT"), os.Getenv("DB_NAME"))
// 	conn, err := pgx.Connect(context.Background(), dsn)

// 	if err != nil {
// 		log.Fatalf("Error Connecting to Database: %v\n", err)
// 		os.Exit(2)

// 	}
// 	// defer conn.Close(context.Background())
// 	log.Print("Successfully connected to database")
// 	return conn
// }
