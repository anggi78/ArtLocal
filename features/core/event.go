package core

import (
	"art-local/features/model"
	"art-local/features/request"
	//"art-local/features/request"
	"art-local/features/response"
)

func EventCoreToEventModel(event EventCore) model.Event {
	dataEvent := model.Event {
		Title: event.Title,
		Date: event.Date,
		Description: event.Description,
		Location: event.Location,
	}

	// for _, v := range event.FollowEvent {
	// 	data := FollowEventCoreToModelFollowEvent(v)
	// 	dataEvent.Follow_event = append(dataEvent.Follow_event, data)
	// }
	return dataEvent
}

func EventModelToEventCore(event model.Event) EventCore {
	dataEvent := EventCore {
		Title: event.Title,
		Date: event.Date,
		Description: event.Description,
		Location: event.Location,
	}
	return dataEvent
}

func EventCoreToEventRespon(event EventCore) response.EventResponse {
	dataEvent := response.EventResponse {
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