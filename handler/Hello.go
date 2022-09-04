package handler

import (
	"github.com/labstack/echo/v4"
	models "go-echo-prototype/model"
	"net/http"
)

func Hello(c echo.Context) error {
	var model models.HelloModel
	err := c.Bind(&model)
	if err != nil {
		return c.String(http.StatusBadRequest, "bad request")
	}

	if len(model.Message) > 0 {
		return c.String(http.StatusOK, model.Message)
	} else {
		return c.String(http.StatusOK, "Hello, World!")
	}
}
