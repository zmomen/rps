package handler

import (
	"net/http"
	"rps/processor"
	"rps/model"

	"github.com/labstack/echo/v4"
)

type Handler struct {
	Processor processor.Processor
}

func (h *Handler) ProcessReceipt(c echo.Context) error {
	
	var request model.ReceiptRequest
	if err := c.Bind(&request); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	result, err := h.Processor.CalculatePoints(request)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusCreated, result)
}

func (h *Handler) GetReceiptPoints(c echo.Context) error {
	idParam := c.Param("id")
	response := h.Processor.GetPoints(idParam)
	return c.JSON(http.StatusOK, response)
}
