package core

import (
	"art-local/features/model"
	"art-local/features/request"
	"art-local/features/response"
)

func EventRequestToFollowEventCore(follow request.EventRequest) FollowEventCore {
	dataFollowEvent := FollowEventCore {
		UserID: follow.AdminID,
	}
	return dataFollowEvent
}

func FollowEventCoreToModelFollowEvent(follow FollowEventCore) model.FollowEvent {
	dataFollow := model.FollowEvent {
		UserID: follow.UserID,
		EventID: follow.EventID,
	}
	return dataFollow
}

func FollowEventCoreToFollowEventResp(follow FollowEventCore) response.FollowEventResponse {
	dataFollow := response.FollowEventResponse {
		UserID: follow.UserID,
		EventID: follow.EventID,
	}
	return dataFollow
}

func FollowEventModeltoFollowEventCore(follow model.FollowEvent) FollowEventCore {
	dataFollow := FollowEventCore {
		UserID: follow.UserID,
		EventID: follow.EventID,
	}
	return dataFollow
}