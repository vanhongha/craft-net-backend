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

	response := &model.GetUserResponse{
		User: user,
	}

	if !lo.IsEmpty(user.AvatarMediaID) {
		mediaData, err := services.GetMedia(*user.AvatarMediaID)
		if !lo.IsNil(err) {
			return nil, errors.New(err.Message)
		}

		url, err := aws.GetFile(mediaData.URLPath)
		if !lo.IsNil(err) {
			return nil, errors.New(err.Message)
		}
		response.AvatarURL = &url
	}

	if !lo.IsEmpty(user.CoverMediaID) {
		mediaData, err := services.GetMedia(*user.CoverMediaID)
		if !lo.IsNil(err) {
			return nil, errors.New(err.Message)
		}

		url, err := aws.GetFile(mediaData.URLPath)
		if !lo.IsNil(err) {
			return nil, errors.New(err.Message)
		}
		response.CoverURL = &url
	}

	return response, nil
}
