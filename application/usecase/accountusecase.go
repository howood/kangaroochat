package usecase

import (
	"context"
	"errors"

	"github.com/howood/kangaroochat/application/actor"
	"github.com/howood/kangaroochat/application/actor/cacheservice"
	"github.com/howood/kangaroochat/domain/entity"
	log "github.com/howood/kangaroochat/infrastructure/logger"
)

type AccountUsecase struct {
	Ctx context.Context
}

func (au AccountUsecase) Login(identifier string, form entity.LoginRoomForm) (token string, err error) {
	err = au.loginRoom(identifier, form.Password)
	if err != nil {
		return token, err
	}
	token, err = au.createToken(identifier, form.UserName)
	return token, err
}

func (au AccountUsecase) loginRoom(identifier, password string) error {
	cacheAssessor := cacheservice.NewCacheAssessor(au.Ctx)
	if cachedvalue, cachedfound := cacheAssessor.Get(identifier); cachedfound {
		chatroom := actor.NewChatRoomOperator(au.Ctx)
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
			log.Error(au.Ctx, err.Error())
			return err
		}
		if err := chatroom.ComparePassword(password); err != nil {
			return err
		}
		return nil
	}

	return errors.New("No Room")
}

func (au AccountUsecase) SetRoom(form entity.CreateRoomForm) (string, error) {
	chatroom := actor.NewChatRoomOperator(au.Ctx)
	chatroom.Set(form.RoomName, form.Password)
	encodedcached, err := chatroom.GobEncode()
	if err != nil {
		return "", err

	}

	cacheAssessor := cacheservice.NewCacheAssessor(au.Ctx)
	cacheAssessor.Set(chatroom.GetIdentifier(), encodedcached, cacheservice.GetChacheExpired())
	return chatroom.GetIdentifier(), nil
}

func (au AccountUsecase) createToken(identifier, username string) (string, error) {
	jwtinstance := actor.NewJwtOperator(au.Ctx, username, false, identifier)
	tokenstr := jwtinstance.CreateToken(actor.TokenSecret)
	return tokenstr, nil
}
