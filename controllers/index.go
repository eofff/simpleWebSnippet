package controllers

import (
	"gosample/repository"
	"gosample/services"
	"log"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

func IndexGetAll(ctx echo.Context) error {
	samples, err := services.SampleGet()
	if err != nil {
		log.Println(err)
		return ctx.JSON(http.StatusInternalServerError, nil)
	}

	return ctx.JSON(http.StatusOK, samples)
}

func IndexGetById(ctx echo.Context) error {
	id, err := strconv.Atoi(ctx.Param("id"))
	if err != nil {
		log.Println(err)
		return ctx.JSON(http.StatusBadRequest, nil)
	}

	sample, err := services.SampleGetById(id)
	if err != nil {
		log.Println(err)
		return ctx.JSON(http.StatusInternalServerError, nil)
	}

	return ctx.JSON(http.StatusOK, sample)
}

func IndexPost(ctx echo.Context) error {
	var sample repository.Sample

	err := ctx.Bind(&sample)
	if err != nil {
		log.Println(err)
		return ctx.JSON(http.StatusBadRequest, nil)
	}

	err = services.SampleInsert(&sample)
	if err != nil {
		log.Println(err)
		return ctx.JSON(http.StatusBadRequest, nil)
	}

	return ctx.JSON(http.StatusOK, sample)
}
