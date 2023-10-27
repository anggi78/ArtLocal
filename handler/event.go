package handler

import (
	"art-local/features/core"
	"art-local/features/request"
	"art-local/features/response"
	//"art-local/helpers"

	"art-local/helpers"
	"art-local/services"
	"strconv"

	"github.com/labstack/echo/v4"
)

type eventHandler struct {
	eventService services.EventServiceInterface
}

func NewEventHandler(eventService services.EventServiceInterface) *eventHandler {
	return &eventHandler{eventService}
}

func (ev *eventHandler) GetAllEvent(e echo.Context) error {
	events, err := ev.eventService.GetAll()
	if err != nil {
		return response.ResponseJSON(e, 500, err.Error(), nil)
	}

	eventResponse := []response.EventResponse{}
	for _, v := range events {
		event := core.EventCoreToEventRespon(v, v.ID)
		eventResponse = append(eventResponse, event)
	}
	return response.ResponseJSON(e, 200, "success", eventResponse)
}

func (ev *eventHandler) CreateEvent(e echo.Context) error {
	eventRequest := request.EventRequest{}
	if err := e.Bind(&eventRequest); err != nil {
		return response.ResponseJSON(e, 400, err.Error(), nil)
	}
	dataEvent := core.EventRequestToEventCore(eventRequest)
	data, err := ev.eventService.Create(dataEvent, "")
	if err != nil {
		return response.ResponseJSON(e, 500, err.Error(), nil)
	}

	eventResp := core.EventCoreToEventRespon(data, data.ID)
	return response.ResponseJSON(e, 200, "success", eventResp)
}

func (ev *eventHandler) GetByIdEvent(e echo.Context) error {
    idStr := e.Param("id")
    id, err := strconv.Atoi(idStr)
    if err != nil {
        return response.ResponseJSON(e, 400, "Invalid ID", nil)
    }

    event, _, err := ev.eventService.GetById(uint(id))
    if err != nil {
        return response.ResponseJSON(e, 400, err.Error(), nil)
    }

    eventResp := core.EventCoreToEventRespon(event, event.ID)
    return response.ResponseJSON(e, 200, "success", eventResp)
}

func (ev *eventHandler) DeleteEvent(e echo.Context) error {
	idStr := e.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		return response.ResponseJSON(e, 400, "Invalid ID", nil)
	}

	event, err := ev.eventService.Delete(uint(id))
	if err != nil {
		return response.ResponseJSON(e, 400, err.Error(), nil)
	}
	if !event {
		return response.ResponseJSON(e, 400, err.Error(), nil)
	}
	return response.ResponseJSON(e, 200, "success", nil)
}

func (ev *eventHandler) UpdateEvent(e echo.Context) error {
	idStr := e.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return response.ResponseJSON(e, 400, "invalid id", nil)
	}

	newEvent := request.EventRequest{}
	err = e.Bind(&newEvent)
	if err != nil {
		return response.ResponseJSON(e, 400, err.Error(), nil)
	}

	newEvents := core.EventRequestToEventCore(newEvent)
	dataEvent,_, err := ev.eventService.Update(uint(id), newEvents)
	if err != nil {
		return response.ResponseJSON(e, 400, err.Error(), nil)
	}

	eventRespon := core.EventCoreToEventRespon(dataEvent, dataEvent.ID)
	return response.ResponseJSON(e, 200, "succes", eventRespon)
}

func (ev *eventHandler) GetAllFollowEvent(e echo.Context) error {
	UserIDStr := e.QueryParam("users")
	UserID, err := strconv.Atoi(UserIDStr)
	if err != nil {
		return response.ResponseJSON(e, 400, "Invalid UserID", nil)
	}

	event, err := ev.eventService.GetAllFollowEvent(uint(UserID))
	if err != nil {
	   return response.ResponseJSON(e, 500, err.Error(), nil)
	}

	historiResp := []response.EventResponse{}

	for _, v := range event {
	   eventResp := core.EventCoreToEventRespon(v, v.ID)
	   historiResp = append(historiResp, eventResp)
	}

	return response.ResponseJSON(e, 200, "success", historiResp)
}


func (ev *eventHandler) GetByIdFollowEvent(e echo.Context) error {
	eventIDStr := e.Param("id")
	eventID, err := strconv.Atoi(eventIDStr)
	if err != nil {
   		return response.ResponseJSON(e, 400, "Invalid ID", nil)
	}

	history, err := ev.eventService.GetByIdFollowEvent(uint(eventID))
	if err != nil {
		return response.ResponseJSON(e, 400, "Invalid ID", nil)
	}

	historyResp := core.FollowEventCoreToFollowEventResp(history)
	return response.ResponseJSON(e, 200, "success", historyResp)
}

func (ev *eventHandler) GetAllUserEvent(e echo.Context) error {
	userID := helpers.ExtractTokenUserId(e)

	data := ev.eventService.FindEventsFollow(uint(userID))
	return response.ResponseJSON(e, 200, "success", data)
}

func (ev *eventHandler) CreateFollowEvent(e echo.Context) error {
	userID := helpers.ExtractTokenUserId(e)
	followRequest := request.FollowEventRequest{}
	if err := e.Bind(&followRequest); err != nil {
		return response.ResponseJSON(e, 400, err.Error(), nil)
	}
	dataFollow := core.FollowEventReqToFollowEventCore(followRequest, followRequest.UserID)
	data, err := ev.eventService.CreateFollow(dataFollow, uint(userID))
	if err != nil {
		return response.ResponseJSON(e, 500, err.Error(), nil)
	}

	followResp := core.FollowEventCoreToFollowEventResp(data)
	return response.ResponseJSON(e, 200, "success", followResp)
}