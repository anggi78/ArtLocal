package handler

import (
	"art-local/features/core"
	"art-local/features/request"
	"art-local/features/response"
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
		event := core.EventCoreToEventRespon(v)
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
	data, err := ev.eventService.Create(dataEvent)
	if err != nil {
		return response.ResponseJSON(e, 500, err.Error(), nil)
	}

	eventResp := core.EventCoreToEventRespon(data)
	return response.ResponseJSON(e, 200, "success", eventResp)
}

func (ev *eventHandler) GetByIdEvent(e echo.Context) error {
	idStr := e.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return response.ResponseJSON(e, 400, "Invalid ID", nil)
	}

	event, err := ev.eventService.GetById(uint(id))
	if err != nil {
		return response.ResponseJSON(e, 400, err.Error(), nil)
	}

	eventResp := core.EventCoreToEventRespon(event)
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

	eventRespon := core.EventCoreToEventRespon(dataEvent)
	return response.ResponseJSON(e, 200, "succes", eventRespon)
}