package handlers

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/ncalamsyah/go-echo/cmd/api/service"

)

func PostHandler(c echo.Context) error {
	data, err := service.GetAll()
	if err != nil {
		c.String(http.StatusBadGateway, "Unable to process data")
	}

	res := make(map[string]any)
	res["status"] = "ok"
	res["data"] = data
	return c.JSON(http.StatusOK, res)
}

func PostDetailHandler(c echo.Context) error {
	id := c.Param("id")
	idx, err := strconv.Atoi(id)
	if err != nil {
		c.String(http.StatusBadRequest, "Unable to convert param")
	}

	data, err := service.GetById(idx)
	if err != nil {
		c.String(http.StatusInternalServerError, "Unable to fetch data")
	}

	res := make(map[string]any)
	res["status"] = "ok"
	res["data"] = data
	return c.JSON(http.StatusOK, res)
}
