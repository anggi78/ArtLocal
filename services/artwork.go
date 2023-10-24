package services

import (
	"art-local/features/core"
	"art-local/repositories"
)

type ArtServiceInterface interface {
	GetAll() ([]core.ArtworkCore, error)
	GetById(id uint) (core.ArtworkCore, error)
	Create(art core.ArtworkCore) (core.ArtworkCore, error)
	Delete(id uint) (bool, error)
	Update(id uint, updateArt core.ArtworkCore) (core.ArtworkCore, error)
}

type artService struct {
	artRepo repositories.ArtworkRepoInterface
}

func NewArtService(artRepo repositories.ArtworkRepoInterface) *artService {
	return &artService{artRepo}
}

func (a *artService) GetAll() ([]core.ArtworkCore, error) {
	arts, err := a.artRepo.GetAll()
	if err != nil {
		return nil, err
	}
	return arts, nil
}

func (a *artService) GetById(id uint) (core.ArtworkCore, error) {
	arts, err := a.artRepo.GetById(id)
	if err != nil {
		return arts, err
	}
	return arts, nil
}

func (a *artService) Create(art core.ArtworkCore) (core.ArtworkCore, error) {
	arts, err := a.artRepo.Create(art)
	if err != nil {
		return art, err
	}
	return arts, nil
}

func (a *artService) Delete(id uint) (bool, error) {
	arts, err := a.artRepo.Delete(id)
	if err != nil {
		return false, err
	}
	return arts, nil
}

func (a *artService) Update(id uint, updateArt core.ArtworkCore) (core.ArtworkCore, error) {
	arts, err := a.artRepo.Update(id, updateArt)
	if err != nil {
		return arts, err
	}
	return arts, nil
}