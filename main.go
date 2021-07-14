package main

import (
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	s := http.Server{
		Addr:    "8080",
		Handler: e,
	}
	initRoutes(e)
	if err := s.ListenAndServe(); err != http.ErrServerClosed {
		log.Fatal(err)
	}
}

func initRoutes(e *echo.Echo) {
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Start page!")
	})
	e.GET("/hello", helloUser)
	e.Logger.Fatal(e.Start(":8080"))
}

func helloUser(context echo.Context) error {
	userName := context.QueryParam("name")
	return context.String(http.StatusOK, "Hi "+userName)
}
