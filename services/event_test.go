package services

import (
	"art-local/features/core"
	"art-local/mocks"
	"errors"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestCreateEvents_Success(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockEventRepo := mocks.NewMockEventRepoInterface(ctrl)
	eventService := NewEventService(mockEventRepo)

	event := core.EventCore{ 
		Title: "events local",
		Date: "20-10-2023",
		Description: "ini event yang sangat menarik",
		Location: "Palu",
	 }
	expectedEvent := core.EventCore{
		Title: "events local",
		Date: "20-10-2023",
		Description: "ini event yang sangat menarik",
		Location: "Palu",
	}

	mockEventRepo.EXPECT().Create(event).Return(expectedEvent, nil)

	createdEvent, err := eventService.Create(event, "arg1")

	assert.NoError(t, err)
	assert.Equal(t, expectedEvent, createdEvent)
}

func TestCreateEvents_Failure(t *testing.T) {
    ctrl := gomock.NewController(t)
    defer ctrl.Finish()

    mockEventRepo := mocks.NewMockEventRepoInterface(ctrl)
    eventService := NewEventService(mockEventRepo)

    event := core.EventCore{ 
		Title: "events local",
		Date: "20-10-2023",
		Description: "ini event yang sangat menarik",
		Location: "Palu",
	 }
    expectedError := errors.New("Simulated error")

    mockEventRepo.EXPECT().
        Create(event).
        Return(core.EventCore{}, expectedError)

    createdEvent, err := eventService.Create(event, "arg1")

    assert.Error(t, err)
    assert.Equal(t, expectedError, err)
    assert.Empty(t, core.EventCore{}, createdEvent) 
}

func TestCreateFollow_Success(t *testing.T) {
    ctrl := gomock.NewController(t)
    defer ctrl.Finish()

    mockEventRepo := mocks.NewMockEventRepoInterface(ctrl)
    eventService := NewEventService(mockEventRepo)

    userID := uint(1)
    followEvent := core.FollowEventCore{
        EventID: 1,
    }

    mockEventRepo.EXPECT().FindEventByID(followEvent.EventID).Return(core.EventCore{ID: followEvent.EventID})

    mockEventRepo.EXPECT().CreateFollow(followEvent, userID).Return(followEvent, nil)

    createdEvent, err := eventService.CreateFollow(followEvent, userID)

    assert.NoError(t, err)
    assert.Equal(t, followEvent, createdEvent)
}


func TestCreateFollow_Failure(t *testing.T) {
    ctrl := gomock.NewController(t)
    defer ctrl.Finish()

    mockEventRepo := mocks.NewMockEventRepoInterface(ctrl)
    eventService := NewEventService(mockEventRepo)

    userID := uint(1)
    followEvent := core.FollowEventCore{
        EventID: 1,
    }
    expectedError := errors.New("Simulated error")

    mockEventRepo.EXPECT().CreateFollow(followEvent, userID).Return(core.FollowEventCore{}, expectedError)

    createdEvent, err := eventService.CreateFollow(followEvent, userID)

    assert.Error(t, err)
    assert.Equal(t, expectedError, err)
    assert.Empty(t, createdEvent)
}

func TestDeleteEvent_Success(t *testing.T) {
    ctrl := gomock.NewController(t)
    defer ctrl.Finish()

    mockEventRepo := mocks.NewMockEventRepoInterface(ctrl)
    eventService := NewEventService(mockEventRepo)

    eventID := uint(1)

    mockEventRepo.EXPECT().
        Delete(eventID).
        Return(true, nil)

    success, err := eventService.Delete(eventID)

    assert.NoError(t, err)
    assert.True(t, success)
}

func TestDeleteEvent_Failure(t *testing.T) {
    ctrl := gomock.NewController(t)
    defer ctrl.Finish()

    mockEventRepo := mocks.NewMockEventRepoInterface(ctrl)
    eventService := NewEventService(mockEventRepo)

    eventID := uint(1)
    expectedError := errors.New("Simulated error")

    mockEventRepo.EXPECT().
        Delete(eventID).
        Return(false, expectedError)

    success, err := eventService.Delete(eventID)

    assert.Error(t, err)
    assert.False(t, success)
}

func TestGetAllEvents_Success(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockEventRepo := mocks.NewMockEventRepoInterface(ctrl)
	eventService := NewEventService(mockEventRepo)

	expectedEvents := []core.EventCore{
		{ID: 1, 
		 Title: "event",
		 Date: "20-10-2023",
		 Description: "halo",
		 Location: "Palu",
		},
		{ID: 2, 
		 Title: "events",
		 Date: "20-10-2024",
		 Description: "haloo",
		 Location: "Manado",
		},
	}
	mockEventRepo.EXPECT().GetAll().Return(expectedEvents, nil)

	events, err := eventService.GetAll()

	assert.NoError(t, err)
	assert.Len(t, events, 2)
}

func TestGetAllEvents_Failure(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockEventRepo := mocks.NewMockEventRepoInterface(ctrl)
	eventService := NewEventService(mockEventRepo)

	expectedError := errors.New("Simulated error")

	mockEventRepo.EXPECT().GetAll().Return(nil, expectedError)

	events, err := eventService.GetAll()

	assert.Error(t, err)
	assert.Len(t, events, 0)
}

func TestUpdateEvent_Success(t *testing.T) {
    ctrl := gomock.NewController(t)
    defer ctrl.Finish()

    mockEventRepo := mocks.NewMockEventRepoInterface(ctrl)
    eventService := NewEventService(mockEventRepo)

    eventID := uint(1)
    event := core.EventCore{
		ID: 2, 
		Title: "events",
		Date: "20-10-2024",
		Description: "haloo",
		Location: "Manado",
	}

    mockEventRepo.EXPECT().GetById(eventID).Return(event, nil)
	mockEventRepo.EXPECT().Update(eventID, gomock.Any()).Return(nil)

    updatedEvent, _, err := eventService.Update(eventID, event)

    assert.NoError(t, err)
    assert.Equal(t, event, updatedEvent)
}

func TestUpdateEvent_Failure(t *testing.T) {
    ctrl := gomock.NewController(t)
    defer ctrl.Finish()

    mockEventRepo := mocks.NewMockEventRepoInterface(ctrl)
    eventService := NewEventService(mockEventRepo)

    eventID := uint(1)
    event := core.EventCore{
		ID: 2, 
		Title: "events",
		Date: "20-10-2024",
		Description: "haloo",
		Location: "Manado",
	}
	expectedError := errors.New("Simulated error")

    mockEventRepo.EXPECT().GetById(eventID).Return(event, nil)
	mockEventRepo.EXPECT().Update(eventID, gomock.Any()).Return(expectedError)

    updatedEvent, _, err := eventService.Update(eventID, event)

    assert.Error(t, err)
    assert.Equal(t, expectedError, err)
	assert.Empty(t, updatedEvent)
}

func TestGetAllFollowEvent_Success(t *testing.T) {
    ctrl := gomock.NewController(t)
    defer ctrl.Finish()

    mockEventRepo := mocks.NewMockEventRepoInterface(ctrl)
    eventService := NewEventService(mockEventRepo)

    userID := uint(1)
    fakeEvents := []core.EventCore{
        {ID: 1, Title: "Event 1", Date: "07-05-2023", Description: "events", Location: "Gorontalo"},
        {ID: 2, Title: "Event 2", Date: "07-05-2022", Description: "event", Location: "Kotamobagu"},
    }

    mockEventRepo.EXPECT().GetAllFollowEvent(userID).Return(fakeEvents, nil)

    events, err := eventService.GetAllFollowEvent(userID)

    assert.NoError(t, err)
    assert.Equal(t, fakeEvents, events)
}

func TestGetAllFollowEvent_Failure(t *testing.T) {
    ctrl := gomock.NewController(t)
    defer ctrl.Finish()

    mockEventRepo := mocks.NewMockEventRepoInterface(ctrl)
    eventService := NewEventService(mockEventRepo)

    userID := uint(1)
    expectedError := errors.New("Simulated error")

    mockEventRepo.EXPECT().GetAllFollowEvent(userID).Return(nil, expectedError)

    events, err := eventService.GetAllFollowEvent(userID)

    assert.Error(t, err)
    assert.Equal(t, expectedError, err)
    assert.Empty(t, events)
}

func TestGetByIdFollowEvent_Success(t *testing.T) {
    ctrl := gomock.NewController(t)
    defer ctrl.Finish()

    mockEventRepo := mocks.NewMockEventRepoInterface(ctrl)
    eventService := NewEventService(mockEventRepo)

    eventID := uint(1)
    expectedFollowEvent := core.FollowEventCore{
        EventID: 1,
    }

    mockEventRepo.EXPECT().
        GetByIdFollowEvent(eventID).
        Return(expectedFollowEvent, nil)

    followEvent, err := eventService.GetByIdFollowEvent(eventID)

    assert.NoError(t, err)
    assert.Equal(t, expectedFollowEvent, followEvent)
}

func TestGetByIdFollowEvent_Error(t *testing.T) {
    ctrl := gomock.NewController(t)
    defer ctrl.Finish()

    mockEventRepo := mocks.NewMockEventRepoInterface(ctrl)
    eventService := NewEventService(mockEventRepo)

    eventID := uint(1)
    expectedError := errors.New("Simulated error")

    mockEventRepo.EXPECT().
        GetByIdFollowEvent(eventID).
        Return(core.FollowEventCore{}, expectedError)

    followEvent, err := eventService.GetByIdFollowEvent(eventID)

    assert.Empty(t, err) 
    assert.Empty(t, followEvent) 
}

func TestFindEventsFollow_Success(t *testing.T) {
    ctrl := gomock.NewController(t)
    defer ctrl.Finish()

    mockEventRepo := mocks.NewMockEventRepoInterface(ctrl)
    eventService := NewEventService(mockEventRepo)

    userID := uint(123)
    eventID := uint(456)

    mockFollowEvent := core.FollowEventCore{
        ID:     1,
        UserID: userID,
        EventID: eventID,
    }

    mockEvent := core.EventCore{
        ID:          eventID,
        Title:       "Mock Event Title",
        Date:        "Mock Event Date",
        Description: "Mock Event Description",
        Location:    "Mock Event Location",
        FollowEvent: []core.FollowEventCore{mockFollowEvent},
    }

    mockEventRepo.EXPECT().FindEventID(userID).Return([]core.FollowEventCore{mockFollowEvent}).Times(1)
    mockEventRepo.EXPECT().FindEventByID(eventID).Return(mockEvent).Times(1)

    event := eventService.FindEventsFollow(userID)
    expectedEvent := []core.EventCore{mockEvent}

    assert.NotNil(t, event)
    assert.Equal(t, expectedEvent, event)
}

func TestGetEvent_Success(t *testing.T) {
    ctrl := gomock.NewController(t)
    defer ctrl.Finish()

    mockEventRepo := mocks.NewMockEventRepoInterface(ctrl)
    expectedEvent := core.EventCore{
        ID:          1,
        Title:       "Sample Event",
        Date:        "2023-10-29",
        Description: "This is a sample event",
        Location:    "Sample Location",
    }

    mockEventRepo.EXPECT().GetById(gomock.Any()).Return(expectedEvent, nil)
    mockEventRepo.EXPECT().FindName(expectedEvent.ID).Return("John Doe")

    eventService := NewEventService(mockEventRepo)

    event, name, err := eventService.GetById(1)
    
    assert.Nil(t, err)
    assert.NotNil(t, event)
    assert.Equal(t, expectedEvent, event)
    assert.Equal(t, "John Doe", name)
}

func TestGetEvent_Failure(t *testing.T) {
    ctrl := gomock.NewController(t)
    defer ctrl.Finish()

    mockEventRepo := mocks.NewMockEventRepoInterface(ctrl)
    mockEventRepo.EXPECT().GetById(gomock.Any()).Return(core.EventCore{}, errors.New("Database error"))

    eventService := NewEventService(mockEventRepo)
    event, _, err := eventService.GetById(1)

    expectedError := errors.New("Database error")
    assert.Error(t, err)
    assert.Equal(t, expectedError, err)
    assert.Empty(t, event)
}