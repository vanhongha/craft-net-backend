package handlers

import (
	"craftnet/graph/model"
	"craftnet/internal/app/services"
	"craftnet/internal/aws"
	"craftnet/internal/util"
	"errors"

	"github.com/samber/lo"
)

func GetUser(userID int) (*model.GetUserResponse, error) {
	user, err := services.GetUser(userID)
	if !lo.IsNil(err) {
		return nil, errors.New(err.Message)
	}

	if lo.IsNil(user) {
		errMsg := util.ErrorMessage(util.ERROR_CODE[util.FAIL_TO_GET], "user")
		util.GetLogger().LogErrorWithMsg(errMsg, false)
		return nil, errors.New(errMsg)
	}

	if !lo.IsEmpty(user.AvatarImgPath) {
		url, err := aws.GetFile(*user.AvatarImgPath)
		if !lo.IsNil(err) {
			return nil, errors.New(err.Message)
		}
		*user.AvatarImgPath = url
	}

	response := &model.GetUserResponse{
		User: user,
	}

	return response, nil
}
