package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func processReceipt(c echo.Context) error {
	return c.JSON(http.StatusCreated, 201)
}

func getReceiptPoints(c echo.Context) error {

	return c.JSON(http.StatusOK, "points response")
}

func homePage(c echo.Context) error {
	return c.String(http.StatusOK, "Hello! from RPS - Receipt Processor Service!")
}

func main() {
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Routes
	e.GET("/", homePage)
	e.POST("/receipts/process", processReceipt)
	e.GET("/receipts/:id/points", getReceiptPoints)

	// Start server
	e.Logger.Fatal(e.Start(":1323"))
}
