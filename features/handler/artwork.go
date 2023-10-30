package handler

import (
	"art-local/entity/core"
	"art-local/entity/request"
	"art-local/entity/response"
	"art-local/helpers"
	"art-local/features/services"
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

func (a *arthandler) GetAllArt(e echo.Context) error {
	arts, err := a.artService.GetAll()
	if err != nil {
		return response.ResponseJSON(e, 500, err.Error(), nil)
	}

	artResponse := []response.ArtworkResponse{}
	for _, v := range arts {
		art := core.ArtworkCoreToArtworkResponse(v, v.ID)
		artResponse = append(artResponse, art)
	}
	return response.ResponseJSON(e, 200, "success", artResponse)
}

func (a *arthandler) CreateArt(e echo.Context) error {
	artRequest := request.ArtworkRequest{}
	if err := e.Bind(&artRequest); err != nil {
		fmt.Println(artRequest)
		return response.ResponseJSON(e, 400, err.Error(), nil)
	}

	file, err := e.FormFile("image")
	if err != nil {
		return e.JSON(400, "Failed to receive file")
	}

	client := helpers.ConfigCloud()
	imageurl := helpers.UploadFile(file, client)
	dataArt := core.ArtworkDataRequestToArtworkCore(artRequest, imageurl)
	data, err := a.artService.Create(dataArt)
	if err != nil {
		return response.ResponseJSON(e, 500, err.Error(), nil)
	}

	artResp := core.ArtworkCoreToArtworkResponse(data, data.ID)
	return response.ResponseJSON(e, 200, "success", artResp)
}

func (a *arthandler) GetByIdArt(e echo.Context) error {
	idStr := e.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return response.ResponseJSON(e, 400, "Invalid ID", nil)
	}

	art, err := a.artService.GetById(uint(id))
	if err != nil {
		return response.ResponseJSON(e, 400, err.Error(), nil)
	}

	artResp := core.ArtworkCoreToArtworkResponse(art, art.ID)
	return response.ResponseJSON(e, 200, "success", artResp)
}

func (a *arthandler) DeleteArt(e echo.Context) error {
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

func (a *arthandler) UpdateArt(e echo.Context) error {
    idStr := e.Param("id")
    id, err := strconv.Atoi(idStr)
    if err != nil {
        return response.ResponseJSON(e, 400, "Invalid ID", nil)
    }

    newArt := request.ArtworkRequest{}

    err = e.Bind(&newArt)
    if err != nil {
        return response.ResponseJSON(e, 400, err.Error(), nil)
    }

    file, err := e.FormFile("image")
    if err != nil {
        return e.JSON(400, "Failed to receive file")
    }

    client := helpers.ConfigCloud()
    imageurl := helpers.UploadFile(file, client)

    NewArt := core.ArtworkDataRequestToArtworkCore(newArt, imageurl)
    dataArt, err := a.artService.Update(uint(id), NewArt)
    if err != nil {
        return response.ResponseJSON(e, 400, err.Error(), nil)
    }
    dataArt.Image = imageurl
    artRespon := core.ArtworkCoreToArtworkResponse(dataArt, dataArt.ID)
    return response.ResponseJSON(e, 200, "success", artRespon)
}

