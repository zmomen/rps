package main

import (
	"net/http"
	"rps/handler"
	"rps/processor"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func homePage(c echo.Context) error {
	return c.String(http.StatusOK, "Hello! from RPS - Receipt Processor Service!")
}

func main() {
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Initialize Processor
	p := processor.Processor{}

	// Initialize Handler
	h := handler.Handler{
		Processor: p,
	}

	// Routes
	e.GET("/", homePage)
	e.POST("/receipts/process", h.ProcessReceipt)
	e.GET("/receipts/:id/points", h.GetReceiptPoints)

	// Start server
	e.Logger.Fatal(e.Start(":1323"))
}
