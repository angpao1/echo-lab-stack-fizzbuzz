package main

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

//----------
// Handlers
//----------

func fizzBuzz(c echo.Context) error {
	input, _ := strconv.Atoi(c.Param("input"))

	switch {
	case input%15 == 0:
		return c.JSON(http.StatusOK, "fizzbuzz")
	case input%5 == 0:
		return c.JSON(http.StatusOK, "buzz")
	case input%3 == 0:
		return c.JSON(http.StatusOK, "fizz")
	default:
		return c.JSON(http.StatusOK, strconv.Itoa(input))
	}
}

func main() {
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Routes
	e.GET("/:input", fizzBuzz)

	// Start server
	e.Logger.Fatal(e.Start(":1323"))
}
