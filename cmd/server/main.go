package main

import (
	"chatapp/internal/config"
	"chatapp/internal/handler"
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/pressly/goose/v3"
)

func main() {
	godotenv.Load()
	Port := ":" + os.Getenv("PORT")

	db, err := config.ConnectPostgres()
	if err != nil {
		panic(err)
	}
	defer db.Close()
	migration := "./migrations"
	if err := goose.Up(db, migration); err != nil {
		panic(err)
	}

	app := gin.Default()
	handler.Router(app)

	server := &http.Server{
		Addr:    Port,
		Handler: app.Handler(),
	}

	go func() {
		if err := server.ListenAndServe(); err != nil {
			panic(err)
		}
	}()

	fmt.Println("Server is running on port", Port)

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	<-quit
	fmt.Println("Shutting down server...")
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)

	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		fmt.Println("Server forced to shutdown:", err)
	}
	fmt.Println("Server exited properly")
}
