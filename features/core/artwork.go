package core

import "art-local/features/request"

func ArtworkRequestToUserCore(art request.ArtworkRequest) request.UserRequest {
	dataArt := request.UserRequest{
		Name: art.Title,
	}
	return dataArt
}