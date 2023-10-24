package handler

import (
	"art-local/features/core"
	"art-local/features/request"
	"art-local/features/response"
	"art-local/services"
	"fmt"
	"strconv"

	"github.com/labstack/echo/v4"
)

type arthandler struct {
	artService services.ArtServiceInterface
}

func NewArtHandler(artService services.ArtServiceInterface) *arthandler {
	return &arthandler{artService}
}

func (a *arthandler) GetAll(e echo.Context) error {
	arts, err := a.artService.GetAll()
	if err != nil {
		return response.ResponseJSON(e, 500, err.Error(), nil)
	}

	artResponse := []response.ArtworkResponse{}
	for _, v := range arts {
		art := core.ArtworkCoreToArtworkResponse(v)
		artResponse = append(artResponse, art)
	}
	return response.ResponseJSON(e, 200, "success", artResponse)
}

func (a *arthandler) Create(e echo.Context) error {
	artRequest := request.ArtworkRequest{}
	if err := e.Bind(&artRequest); err != nil {
		fmt.Println(artRequest)
		return response.ResponseJSON(e, 400, err.Error(), nil)
	}

	dataArtwork := core.ArtworkDataRequestToArtworkCore(artRequest)
	data, err := a.artService.Create(dataArtwork)
	if err != nil {
		fmt.Println(dataArtwork)
		return response.ResponseJSON(e, 500, err.Error(), nil)
	}

	artResp := core.ArtworkCoreToArtworkResponse(data)
	return response.ResponseJSON(e, 200, "success", artResp)
}

func (a *arthandler) GetById(e echo.Context) error {
	idStr := e.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
    	return response.ResponseJSON(e, 400, "Invalid ID", nil)
	}

	art, err := a.artService.GetById(uint(id))
	if err != nil {
    	return response.ResponseJSON(e, 400, err.Error(), nil)
	}

	artResp := core.ArtworkCoreToArtworkResponse(art)
	return response.ResponseJSON(e, 200, "success", artResp)
}

func (a *arthandler) Delete(e echo.Context) error {
	idStr := e.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
    	return response.ResponseJSON(e, 400, "Invalid ID", nil)
	}

	art, err := a.artService.Delete(uint(id))
	if err != nil {
    	return response.ResponseJSON(e, 400, err.Error(), nil)
	}
	if !art {
    	return response.ResponseJSON(e, 400, err.Error(), nil)
	}
	return response.ResponseJSON(e, 200, "success", nil)
}

func (a *arthandler) Update(e echo.Context) error {
	idStr := e.Param("id")
	id , err := strconv.Atoi(idStr)
	if err != nil {
		return response.ResponseJSON(e, 400, "invalid id", nil)
	}

	newArtwork := request.ArtworkRequest{}
	err = e.Bind(&newArtwork)
	if err != nil {
		return response.ResponseJSON(e, 400, err.Error(), nil)
	}
	newArtworks := core.ArtworkDataRequestToArtworkCore(newArtwork)

	dataArt, err := a.artService.Update(uint(id), newArtworks)
	if err != nil {
		return response.ResponseJSON(e, 400, err.Error(), nil)
	}

	artRespon := core.ArtworkCoreToArtworkResponse(dataArt)
	return response.ResponseJSON(e, 200, "succes", artRespon)
}