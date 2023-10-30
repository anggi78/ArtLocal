package core

import (
	"art-local/entity/model"
	"art-local/entity/request"
	"art-local/entity/response"
)

func ArtworkDataRequestToArtworkCore(request request.ArtworkRequest, imageUrl string) ArtworkCore {
	artworkCore := ArtworkCore{
		Title:       request.Title,
		Image: 		 imageUrl,
		Description: request.Description,
		UserID:      request.UserID,
	}
	return artworkCore
}

func ArtworkCoreToArtworkModel(coreArtwork ArtworkCore) model.Artwork {
	artworkModel := model.Artwork{
		ID: 		 coreArtwork.ID,
		Title:       coreArtwork.Title,
		Image: 		 coreArtwork.Image,
		Description: coreArtwork.Description,
		UserID:      uint(coreArtwork.UserID),
	}
	return artworkModel
}

func ArtworkCoreToArtworkResponse(coreArtwork ArtworkCore, ID uint) response.ArtworkResponse {
	artworkResponse := response.ArtworkResponse{
		ID: 	 uint(coreArtwork.ID),
		Title:       coreArtwork.Title,
		Image: 		 coreArtwork.Image,
		Description: coreArtwork.Description,
		UserID:      uint(coreArtwork.UserID),
	}
	return artworkResponse
}

func ArtworkModelToArtworkCore(modelArtwork model.Artwork) ArtworkCore {
	artworkCore := ArtworkCore{
		ID: modelArtwork.ID,
		Title:       modelArtwork.Title,
		Image: modelArtwork.Image,
		Description: modelArtwork.Description,
		UserID:      uint(modelArtwork.UserID),
	}
	return artworkCore
}
