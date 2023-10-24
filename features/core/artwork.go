package core

import (
	"art-local/features/model"
	"art-local/features/request"
	"art-local/features/response"
)

func ArtworkDataRequestToArtworkCore(request request.ArtworkRequest) ArtworkCore {
	artworkCore := ArtworkCore{
		Title:       request.Title,
		//Image: 		 imageurl,
		Description: request.Description,
		UserID:      request.UserID,
	}
	return artworkCore
}

func ArtworkCoreToArtworkModel(coreArtwork ArtworkCore) model.Artwork {
	artworkModel := model.Artwork{
		Title:       coreArtwork.Title,
		//Image: 		 coreArtwork.Image,
		Description: coreArtwork.Description,
		UserID:      coreArtwork.UserID,
	}
	return artworkModel
}

func ArtworkCoreToArtworkResponse(coreArtwork ArtworkCore) response.ArtworkResponse {
	artworkResponse := response.ArtworkResponse{
		Title:       coreArtwork.Title,
		Description: coreArtwork.Description,
		UserID:      coreArtwork.UserID,
	}
	return artworkResponse
}

func ArtworkModelToArtworkCore(modelArtwork model.Artwork) ArtworkCore {
	artworkCore := ArtworkCore{
		Title:       modelArtwork.Title,
		Description: modelArtwork.Description,
		UserID:      modelArtwork.UserID,
	}
	return artworkCore
}
