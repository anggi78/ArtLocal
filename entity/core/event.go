package core

import (
	"art-local/entity/model"
	"art-local/entity/request"
	"art-local/entity/response"
)

func EventCoreToEventModel(event EventCore) model.Event {
	dataEvent := model.Event {
		ID: event.ID,
		Title: event.Title,
		Date: event.Date,
		Description: event.Description,
		Location: event.Location,
	}

	for _, v := range event.FollowEvent {
		data := FollowEventCoreToModelFollowEvent(v)
		dataEvent.FollowEvent = append(dataEvent.FollowEvent, data)
	}
	return dataEvent
}

func EventModelToEventCore(event model.Event) EventCore {
	dataEvent := EventCore {
		ID: event.ID,
		Title: event.Title,
		Date: event.Date,
		Description: event.Description,
		Location: event.Location,
	}
	return dataEvent
}

func EventCoreToEventRespon(event EventCore, ID uint) response.EventResponse {
	dataEvent := response.EventResponse {
		ID: uint(event.ID),
		Title: event.Title,
		Date: event.Date,
		Description: event.Description,
		Location: event.Location,
	}
	return dataEvent
}

func EventRequestToEventCore(event request.EventRequest) EventCore {
	dataEvent := EventCore {
		Title: event.Title,
		Date: event.Date,
		Description: event.Description,
		Location: event.Location,
	}
	return dataEvent
}

func EventCoreToEventAll(event EventCore) EventAll {
	allEvent := EventAll {
		ID: event.ID,
		Name:		"",
		Title:		event.Title,
		Description: event.Description,
		Date: 		 event.Date,
	}
	return allEvent
}