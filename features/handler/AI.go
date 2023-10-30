package handler

import (
	"art-local/entity/request"
	"art-local/entity/response"
	"art-local/features/services"
	"fmt"
	"os"

	"github.com/labstack/echo/v4"
)

type evtHandler struct {
	evtRecService services.EventReccomInterface
}

func NewEventRecc(evtRecService services.EventReccomInterface) *evtHandler {
	return &evtHandler{evtRecService}
}

func (h *evtHandler) Reccomend(e echo.Context) error {
	req := request.AskAboutEvent{}
	err := e.Bind(&req)
	if err != nil {
		return response.ResponseJSON(e, 400, err.Error(), nil)
	}

	userInput := fmt.Sprintf("Event ini masuk dalam kategori %s dan akan berlangsung di %s. Dengan tingkat popularitas yang populer, acara ini tidak boleh Anda lewatkan!",  req.Category, req.Location)
	answer, err := h.evtRecService.EventReccomend(userInput, req.Category, req.Location, os.Getenv("OPEN_AI_KEY"))
	if err != nil {
		return response.ResponseJSON(e, 500, err.Error(), nil)
	}
	return response.ResponseJSON(e, 200, "success", answer)
}