package server

import (
	"fmt"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/LucasZatta/ProductCrud/internal/database"

	_ "github.com/joho/godotenv/autoload"
)

type Server struct {
	port int

	db database.Service
}

func NewServer() *http.Server {
	port, _ := strconv.Atoi(os.Getenv("PORT"))
	fmt.Println((port))
	NewServer := &Server{
		port: port,

		db: database.New(),
	}
	server := &http.Server{
		Addr:         fmt.Sprintf(":%d", NewServer.port),
		Handler:      NewServer.RegisterRoutes(),
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	return server
}
