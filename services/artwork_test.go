package services

import (
	"art-local/features/core"
	"art-local/mocks"
	"errors"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestCreateArtwork_Success(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockArtworkRepo := mocks.NewMockArtworkRepoInterface(ctrl)
	artService := NewArtService(mockArtworkRepo)

	inputArtwork := core.ArtworkCore{ 
		Title: "art jog",
		Image: "art",
		Description: "art",
		UserID: 1,
	}
	expectedArtwork := core.ArtworkCore{ 
		Title: "art jog",
		Image: "art",
		Description: "art",
		UserID: 1,
	 }

	mockArtworkRepo.EXPECT().
		Create(inputArtwork).
		Return(expectedArtwork, nil) 

	createdArtwork, err := artService.Create(inputArtwork)

	assert.NoError(t, err)
	assert.Equal(t, expectedArtwork, createdArtwork)
}

func TestCreateArtwork_Failure(t *testing.T) {
    ctrl := gomock.NewController(t)
    defer ctrl.Finish()

    mockArtworkRepo := mocks.NewMockArtworkRepoInterface(ctrl)
    artService := NewArtService(mockArtworkRepo)

    inputArtwork := core.ArtworkCore{
		Title: "art jog",
		Image: "art",
		Description: "art",
		UserID: 1,
	}
    expectedError := errors.New("Simulated error") 

    mockArtworkRepo.EXPECT().Create(gomock.Any()).Return(core.ArtworkCore{}, expectedError)

    _, err := artService.Create(inputArtwork)

    assert.Error(t, err)
    assert.Equal(t, expectedError, err)
}

func TestDeleteArtwork_Success(t *testing.T) {
    ctrl := gomock.NewController(t)
    defer ctrl.Finish()

    mockArtRepo := mocks.NewMockArtworkRepoInterface(ctrl)
    artService := NewArtService(mockArtRepo)

    artworkID := uint(1)

    mockArtRepo.EXPECT().Delete(artworkID).Return(true, nil)

    deleted, err := artService.Delete(artworkID)

    assert.NoError(t, err)
    assert.True(t, deleted)
}

func TestDeleteArtwork_Failure(t *testing.T) {
    ctrl := gomock.NewController(t)
    defer ctrl.Finish()

    mockArtRepo := mocks.NewMockArtworkRepoInterface(ctrl)
    artService := NewArtService(mockArtRepo)

    artworkID := uint(1)
    expectedError := errors.New("Simulated error")

    mockArtRepo.EXPECT().Delete(artworkID).Return(false, expectedError)

    deleted, err := artService.Delete(artworkID)

    assert.Error(t, err)
    assert.False(t, deleted)
    assert.Equal(t, expectedError, err)
}

func TestGetAllArtworks_Success(t *testing.T) {
    ctrl := gomock.NewController(t)
    defer ctrl.Finish()

    mockArtRepo := mocks.NewMockArtworkRepoInterface(ctrl)
    artService := NewArtService(mockArtRepo)

    expectedArtworks := []core.ArtworkCore{
        {Title: "art jog",
		Image: "art",
		Description: "art",
		UserID: 1,},
		{Title: "art",
		Image: "jog",
		Description: "jog",
		UserID: 2,},
    }

    mockArtRepo.EXPECT().
        GetAll().
        Return(expectedArtworks, nil)

    artworks, err := artService.GetAll()

    assert.NoError(t, err)
    assert.Equal(t, expectedArtworks, artworks)
}

func TestGetAllArtworks_Failure(t *testing.T) {
    ctrl := gomock.NewController(t)
    defer ctrl.Finish()

    mockArtRepo := mocks.NewMockArtworkRepoInterface(ctrl)
    artService := NewArtService(mockArtRepo)

    expectedError := errors.New("Simulated error")

    mockArtRepo.EXPECT().
        GetAll().
        Return(nil, expectedError)

    artworks, err := artService.GetAll()

    assert.Error(t, err)
    assert.Equal(t, expectedError, err)
    assert.Empty(t, artworks)
}

func TestGetArtworkCoreById_Success(t *testing.T) {
    ctrl := gomock.NewController(t)
    defer ctrl.Finish()

    mockArtRepo := mocks.NewMockArtworkRepoInterface(ctrl)
    artService := NewArtService(mockArtRepo)

    artworkID := uint(1)
    expectedArtworkCore := core.ArtworkCore{ID: 1, Title: "Artwork 1"}

    mockArtRepo.EXPECT().GetById(artworkID).Return(expectedArtworkCore, nil)

    artworkCore, err := artService.GetById(artworkID)

    assert.NoError(t, err)
    assert.Equal(t, expectedArtworkCore, artworkCore)
}

func TestGetArtworkCoreById_Failure(t *testing.T) {
    ctrl := gomock.NewController(t)
    defer ctrl.Finish()

    mockArtRepo := mocks.NewMockArtworkRepoInterface(ctrl)
    artService := NewArtService(mockArtRepo)

    artworkID := uint(1)
    expectedError := errors.New("Artwork not found")

    mockArtRepo.EXPECT().GetById(artworkID).Return(core.ArtworkCore{}, expectedError)

    artworkCore, err := artService.GetById(artworkID)

    assert.Error(t, err)
    assert.Equal(t, expectedError, err)
    assert.Empty(t, artworkCore)
}

func TestUpdateArtwork_Success(t *testing.T) {
    ctrl := gomock.NewController(t)
    defer ctrl.Finish()

    mockArtworkRepo := mocks.NewMockArtworkRepoInterface(ctrl)
    artService := NewArtService(mockArtworkRepo)

    artworkID := uint(1)
    artwork := core.ArtworkCore{
        Title: "art jog",
		Image: "art",
		Description: "art",
		UserID: 1,
    }

    mockArtworkRepo.EXPECT().Update(artworkID, artwork).Return(artwork, nil)

    updatedArtwork, err := artService.Update(artworkID, artwork)

    assert.NoError(t, err)
    assert.Equal(t, artwork, updatedArtwork)
}

func TestUpdateArtwork_Failure(t *testing.T) {
    ctrl := gomock.NewController(t)
    defer ctrl.Finish()

    mockArtworkRepo := mocks.NewMockArtworkRepoInterface(ctrl)
    artService := NewArtService(mockArtworkRepo)

    artworkID := uint(1)
    artwork := core.ArtworkCore{
        Title: "art jog",
		Image: "art",
		Description: "art",
		UserID: 1,
    }
    expectedError := errors.New("Simulated error")

    mockArtworkRepo.EXPECT().Update(artworkID, artwork).Return(core.ArtworkCore{}, expectedError)

    updatedArtwork, err := artService.Update(artworkID, artwork)

    assert.Error(t, err)
    assert.Equal(t, expectedError, err)
    assert.Empty(t, updatedArtwork)
}