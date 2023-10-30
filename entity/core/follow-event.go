package core

import (
	"art-local/entity/model"
	"art-local/entity/request"
	"art-local/entity/response"
)

func FollowEventReqToFollowEventCore(follow request.FollowEventRequest, UserID uint) FollowEventCore {
	dataFollowEvent := FollowEventCore {
	 	UserID: UserID,
		EventID: follow.EventID,
	}
	return dataFollowEvent
}

func FollowEventCoreToModelFollowEvent(follow FollowEventCore) model.FollowEvent {
	dataFollow := model.FollowEvent {
		UserID: follow.UserID,
		EventID: uint(follow.EventID),
	}
	return dataFollow
}

func FollowEventCoreToFollowEventResp(follow FollowEventCore) response.FollowEventResponse {
	dataFollow := response.FollowEventResponse {
		UserID: uint(follow.UserID),
		EventID: uint(follow.EventID),
	}
	return dataFollow
}

func FollowEventModeltoFollowEventCore(follow model.FollowEvent) FollowEventCore {
	dataFollow := FollowEventCore {
		UserID: uint(follow.UserID),
		EventID: uint(follow.EventID),
	}
	return dataFollow
}

