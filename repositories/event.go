package repositories

import (
	"art-local/features/core"
	"art-local/features/model"

	"gorm.io/gorm"
)

type EventRepoInterface interface {
	GetAll() ([]core.EventCore, error)
	GetById(id uint) (core.EventCore, error)
	Create(event core.EventCore) (core.EventCore, error)
	Delete(id uint) (bool, error)
	Update(id uint, updateEvent core.EventCore) error
}

type EventRepo struct {
	db *gorm.DB
}

func NewEventRepositories(db *gorm.DB) *EventRepo {
	return &EventRepo{db}
}

func (e *EventRepo) GetAll() ([]core.EventCore, error) {
	var events []model.Event
	var dataEvent []core.EventCore

	err := e.db.Find(&events).Error
	if err != nil {
		return nil, err
	}

	for _, v := range events {
		eventCore := core.EventModelToEventCore(v)
		dataEvent = append(dataEvent, eventCore)
	}
	return dataEvent, nil
}

func (e *EventRepo) GetById(id uint) (core.EventCore, error) {
	event := model.Event{}
	dataEvent := core.EventCore{}

	err := e.db.Where("id = ?", id).First(&event).Error
	if err != nil {
		return dataEvent, err
	}

	dataEvent = core.EventModelToEventCore(event)
	return dataEvent, nil
}

func (e *EventRepo) Create(event core.EventCore) (core.EventCore, error) {
	eventInput := core.EventCoreToEventModel(event)
	err := e.db.Create(&eventInput).Error
	if err != nil {
		return event, err
	}

	result := core.EventModelToEventCore(eventInput)
	return result, nil
}

func (e *EventRepo) Delete(id uint) (bool, error) {
	event, err := e.GetById(id)

	dataEvents := core.EventCoreToEventModel(event)
	if err != nil {
		return false, err
	}

	err = e.db.Where("id = ?", id).Delete(&dataEvents).Error
	if err != nil {
		return false, err
	}
	return true, nil
}

func (e *EventRepo) Update(id uint, updateEvent core.EventCore) error {
	var event model.Event

	if err := e.db.First(&event, id).Error; err != nil {
		return err
	}

	event.Title 		= updateEvent.Title
	event.Date 			= updateEvent.Date
	event.Description 	= updateEvent.Description
	event.Location 		= updateEvent.Location

	if err := e.db.Save(&event).Error; err != nil {
		return err
	}
	return nil
}