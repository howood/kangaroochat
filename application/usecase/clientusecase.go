package usecase

import (
	"context"
	"errors"

	"github.com/howood/kangaroochat/application/actor"
	"github.com/howood/kangaroochat/application/actor/cacheservice"
	log "github.com/howood/kangaroochat/infrastructure/logger"
)

type ClientUsecase struct {
	Ctx context.Context
}

func (cu ClientUsecase) GetRoomname(identifier string) (string, error) {
	cacheAssessor := cacheservice.NewCacheAssessor(cu.Ctx)
	if cachedvalue, cachedfound := cacheAssessor.Get(identifier); cachedfound {
		chatroom := actor.NewChatRoomOperator(cu.Ctx)
		var err error
		switch xi := cachedvalue.(type) {
		case []byte:
			err = chatroom.GobDecode(xi)
		case string:
			err = chatroom.GobDecode([]byte(xi))
		default:
			err = errors.New("get cache error")
		}
		if err != nil {
			log.Error(cu.Ctx, err.Error())
			return "", err
		}
		return chatroom.GetRoomName(), nil
	}

	return "", errors.New("No Room")
}
