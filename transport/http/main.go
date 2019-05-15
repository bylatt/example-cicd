package main

import (
	"fmt"
	"net/http"
	"os"
	"strconv"

	"github.com/clozed2u/example-cicd/book"
	"github.com/clozed2u/example-cicd/transport/http/handler"
	"github.com/clozed2u/example-cicd/usecase"
	"github.com/go-redis/redis"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main() {
	redisEndpoint := os.Getenv("REDISENDPOINT")
	redisPort := os.Getenv("REDISPORT")
	redisDB, _ := strconv.Atoi(os.Getenv("REDISDB"))
	redisClient := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", redisEndpoint, redisPort),
		Password: "",
		DB:       redisDB,
	})
	bookRedisAdapter := book.NewRedisAdapter(redisClient)
	bookRepo := book.NewRepo(bookRedisAdapter)
	bookUsecase := usecase.NewBookUsecase(bookRepo)
	bookHandler := handler.NewBookHandler(bookUsecase)

	e := echo.New()
	e.Use(middleware.Gzip())
	e.Use(middleware.Logger())

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	// bookRoute := e.Group("/books")
	e.GET("/books", func(c echo.Context) error {
		return c.String(http.StatusNotImplemented, "Not supported")
	})
	e.POST("/books", bookHandler.CreateBook)
	e.GET("/books/:id", bookHandler.FindBook)
	e.PUT("/books/:id", bookHandler.UpdateBook)
	e.DELETE("/books/:id", bookHandler.DeleteBook)

	e.Logger.Fatal(e.Start(":8000"))
}
