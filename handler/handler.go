package handler

import (
	"net/http"
	"rps/processor"

	"github.com/labstack/echo/v4"
)

type Handler struct {
	Processor processor.Processor
}

func (h *Handler) ProcessReceipt(c echo.Context) error {
	//TODO: implementation here
	h.Processor.CalculatePoints()

	return c.JSON(http.StatusCreated, 201)
}

func (h *Handler) GetReceiptPoints(c echo.Context) error {
	//TODO: implementation here
	response := h.Processor.GetPoints("receiptID from request")
	return c.JSON(http.StatusOK, response)
}
