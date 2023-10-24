package services

import (
	"art-local/features/core"
	"art-local/repositories"
)

type EventServiceInterface interface {
	GetAll() ([]core.EventCore, error)
	GetById(id uint) (core.EventCore, error)
	Create(event core.EventCore) (core.EventCore, error)
	Delete(id uint) (bool, error)
	Update(id uint, updateEvent core.EventCore) (core.EventCore, error)
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

func (e *eventService) GetById(id uint) (core.EventCore, error) {
	events, err := e.eventRepo.GetById(id)
	if err != nil {
		return events, err
	}
	return events, nil
}

func (e *eventService) Create(event core.EventCore) (core.EventCore, error){
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

func (e *eventService) Update(id uint, updateEvent core.EventCore) (core.EventCore, error) {
	events, err := e.eventRepo.Update(id, updateEvent)
	if err != nil {
		return events, err
	}
	return events, nil
}