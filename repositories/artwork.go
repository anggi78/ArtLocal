package repositories

import (
	"art-local/features/core"
	"art-local/features/model"

	"gorm.io/gorm"
)

type ArtworkRepoInterface interface {
	GetAll() ([]core.ArtworkCore, error)
	GetById(id uint) (core.ArtworkCore, error)
	Create(art core.ArtworkCore) (core.ArtworkCore, error)
	Delete(id uint) (bool, error)
	Update(id uint, updateArt core.ArtworkCore) (core.ArtworkCore, error)
}

type ArtRepositories struct {
	db *gorm.DB
}

func NewArtRepositories(db *gorm.DB) *ArtRepositories {
	return &ArtRepositories{db}
}

func (a *ArtRepositories) GetAll() ([]core.ArtworkCore, error) {
	var arts []model.Artwork
	var dataArt []core.ArtworkCore

	err := a.db.Find(&arts).Error
	if err != nil {
		return nil, err
	}

	for _, v := range arts {
		artCore := core.ArtworkModelToArtworkCore(v)
		dataArt = append(dataArt, artCore)
	}
	return dataArt, nil
}

func (a *ArtRepositories) GetById(id uint) (core.ArtworkCore, error) {
	art := model.Artwork{}
	dataArt := core.ArtworkCore{}

	err := a.db.Where("id = ?", id).First(&art).Error
	if err != nil {
		return dataArt, err
	}

	dataArt = core.ArtworkModelToArtworkCore(art)
	return dataArt, nil
}

func (a *ArtRepositories) Create(art core.ArtworkCore) (core.ArtworkCore, error) {
	artInput := core.ArtworkCoreToArtworkModel(art)
	err := a.db.Create(&artInput).Error
	if err != nil {
		return art, err
	}

	result := core.ArtworkModelToArtworkCore(artInput)
	return result, nil
}

func (a *ArtRepositories) Delete(id uint) (bool, error) {
	art, err := a.GetById(id)

	dataArts := core.ArtworkCoreToArtworkModel(art)
	if err != nil {
		return false, err
	}

	err = a.db.Where("id = ?", id).Delete(&dataArts).Error
	if err != nil {
		return false, err
	}
	return true, nil
}

func (a *ArtRepositories) Update(id uint, updateArt core.ArtworkCore) (core.ArtworkCore, error) {
	art := core.ArtworkCoreToArtworkModel(updateArt)
	data, err := a.GetById(id)
	if err != nil {
		return data, err
	}

	err = a.db.Where("id = ?", id).Updates(&art).Error
	if err != nil {
		return data, err
	}

	data.ID = id
	return data, nil
}