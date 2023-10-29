package services

import (
	"art-local/features/core"
	"art-local/repositories"
)

type EventServiceInterface interface {
	GetAll() ([]core.EventCore, error)
	GetById(id uint) (core.EventCore, string, error)
	GetAllFollowEvent(UserID uint) ([]core.EventCore, error)
	GetByIdFollowEvent(EventID uint) (core.FollowEventCore, error)
	Create(event core.EventCore, Name string) (core.EventCore, error)
	Delete(id uint) (bool, error)
	Update(id uint, updateEvent core.EventCore) (core.EventCore, string, error)
	FindEventsFollow(UserID uint) []core.EventCore
	CreateFollow(event core.FollowEventCore, UserID uint) (core.FollowEventCore, error)
}

type eventService struct {
	eventRepo repositories.EventRepoInterface
}

func NewEventService(eventRepo repositories.EventRepoInterface) *eventService {
	return &eventService{eventRepo}
}

func (e *eventService) GetAll() ([]core.EventCore, error) {
	events, err := e.eventRepo.GetAll()
	if err != nil {
		return nil, err
	}
	return events, nil
}

func (e *eventService) GetAllFollowEvent(UserID uint) ([]core.EventCore, error) {
	event, err := e.eventRepo.GetAllFollowEvent(UserID)
	if err != nil {
		return nil, err
	}
	return event, nil
}

func (e *eventService) GetById(id uint) (core.EventCore, string, error) {
	events, err := e.eventRepo.GetById(id)
	if err != nil {
		return events, "", err
	}
	name := e.eventRepo.FindName(events.ID)
	return events, name, nil
}

func (e *eventService) GetByIdFollowEvent(EventID uint) (core.FollowEventCore, error) {
	events, err := e.eventRepo.GetByIdFollowEvent(EventID)
	if err != nil {
		return events, nil
	}
	return events, nil
}

func (e *eventService) Create(event core.EventCore, Name string) (core.EventCore, error) {
	events, err := e.eventRepo.Create(event)
	if err != nil {
		return event, err
	}
	return events, nil
}

func (e *eventService) Delete(id uint) (bool, error) {
	events, err := e.eventRepo.Delete(id)
	if err != nil {
		return false, err
	}
	return events, nil
}

func (e *eventService) Update(id uint, updateEvent core.EventCore) (core.EventCore, string, error) {
	existingEvent, _:= e.eventRepo.GetById(id)

	existingEvent.Title = updateEvent.Title
	existingEvent.Date = updateEvent.Date
	existingEvent.Description = updateEvent.Description
	existingEvent.Location = updateEvent.Location

	if err := e.eventRepo.Update(id, existingEvent); err != nil {
		return core.EventCore{}, "", err
	}
	return existingEvent, "", nil
}

func (e *eventService) FindEventsFollow(UserID uint) []core.EventCore {
	followEvents := e.eventRepo.FindEventID(UserID)
	events := []core.EventCore{}
	for _, v := range followEvents {
		dataEvent := e.eventRepo.FindEventByID(v.EventID)
		events = append(events,dataEvent)
	}
	return events
}

func (e *eventService) CreateFollow(event core.FollowEventCore, UserID uint) (core.FollowEventCore, error) {
	follows, err := e.eventRepo.CreateFollow(event, UserID)
	if err != nil {
		return follows, err
	}
	e.eventRepo.FindEventByID(UserID)
	return follows, nil
} 