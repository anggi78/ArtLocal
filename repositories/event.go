package repositories

import (
	"art-local/features/core"
	"art-local/features/model"
	// "art-local/features/response"

	"gorm.io/gorm"
)

type EventRepoInterface interface {
	GetAllFollowEvent(UserID uint) ([]core.EventCore, error)
	GetByIdFollowEvent(EventID uint) (core.FollowEventCore, error)
	GetAll() ([]core.EventCore, error)
	GetById(id uint) (core.EventCore, error)
	Create(event core.EventCore) (core.EventCore, error)
	Delete(id uint) (bool, error)
	Update(id uint, updateEvent core.EventCore) error
	FindName(UserID uint) string
	FindEventID(UserID uint) []core.FollowEventCore
	FindEventByID(EventID uint) core.EventCore
	CreateFollow(event core.FollowEventCore, UserID uint) (core.FollowEventCore, error)
	//CombineEventAndUser(event response.EventResponse, user response.UserResponse) response.UserEventResponse
}

type EventRepo struct {
	db *gorm.DB
}

func NewEventRepositories(db *gorm.DB) *EventRepo {
	return &EventRepo{db}
}

func (e *EventRepo) GetAllFollowEvent(UserID uint) ([]core.EventCore, error) {
	eventsData := []model.Event{}

	err := e.db.Where("user_id = ?", UserID).Preload("FollowEvent").Find(&eventsData).Error
	if err != nil {
		return nil, err
	}

	EventsCore := []core.EventCore{}
	for _, v := range eventsData {
		data := core.EventModelToEventCore(v)

		for _, follows := range v.FollowEvent {
			followsCore := core.FollowEventModeltoFollowEventCore(follows)
			data.FollowEvent = append(data.FollowEvent, followsCore)
		}
		EventsCore = append(EventsCore, data)
	}
	return EventsCore, nil
}

func (e *EventRepo) GetByIdFollowEvent(EventID uint) (core.FollowEventCore, error) {
	eventData := model.FollowEvent{}

	err := e.db.Where("id = ?", EventID).Preload("Event").First(&eventData).Error
	if err != nil {
		return core.FollowEventCore{}, err
	}

	data := core.FollowEventModeltoFollowEventCore(eventData)
	return data, nil
}

func (e *EventRepo) GetAll() ([]core.EventCore, error) {
	data := []model.Event{}
	err := e.db.Find(&data).Error
	if err != nil {
		return nil, err
	}

	dataResp := []core.EventCore{}
	for _, v := range data {
		dataCore := core.EventModelToEventCore(v)
		dataResp = append(dataResp, dataCore)
	}
	return dataResp, nil
}

func (e *EventRepo) GetById(id uint) (core.EventCore, error) {
	event := model.Event{}
	err := e.db.Where("id = ?", id).Preload("FollowEvent").First(&event).Error
	if err != nil {
		return core.EventCore{}, err
	}

	data := core.EventModelToEventCore(event)
	return data, nil
}

func (e *EventRepo) Create(event core.EventCore) (core.EventCore, error) {
	data := e.db.Begin()
	if data.Error != nil {
		return core.EventCore{}, data.Error
	}

	InsertEvent := core.EventCoreToEventModel(event)
	dataEvent := core.EventCore{}

	err := data.Create(&InsertEvent).Error
	if err != nil {
		data.Rollback()
		return dataEvent, err
	}

	dataEvent = core.EventModelToEventCore(InsertEvent)
	data.Commit()

	for _, v := range InsertEvent.FollowEvent {
		followEvent := core.FollowEventModeltoFollowEventCore(v)
		dataEvent.FollowEvent = append(dataEvent.FollowEvent, followEvent)
	}
	return dataEvent, nil
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

	event.Title = updateEvent.Title
	event.Date = updateEvent.Date
	event.Description = updateEvent.Description
	event.Location = updateEvent.Location

	if err := e.db.Save(&event).Error; err != nil {
		return err
	}
	return nil
}

func (e *EventRepo) FindName(UserID uint) string {
	user := model.User{}
	e.db.Where("id = ?", UserID).First(&user)
	return user.Name
}

func (e *EventRepo) FindEventID(UserID uint) []core.FollowEventCore {
	event := []model.FollowEvent{}
	e.db.Where("user_id = ?", UserID).Find(&event)
	followevents := []core.FollowEventCore{}
	for _, v := range event {
		data := core.FollowEventModeltoFollowEventCore(v)
		followevents = append(followevents, data)
	}
	return followevents
}

func (e *EventRepo) FindEventByID(EventID uint) core.EventCore {
	follow := model.Event{}
	
	e.db.Where("id = ?", EventID).First(&follow)

	data := core.EventModelToEventCore(follow)

	return data
}

func (e *EventRepo) CreateFollow(event core.FollowEventCore, UserID uint) (core.FollowEventCore, error) {
	eventInput := core.FollowEventCoreToModelFollowEvent(event)
	eventInput.UserID = UserID
	err := e.db.Create(&eventInput).Error
	if err != nil {
		return event, err
	}

	result := core.FollowEventModeltoFollowEventCore(eventInput)
	return result, nil
}