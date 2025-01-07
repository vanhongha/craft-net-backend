package services

import (
	"craftnet/graph/model"
	"craftnet/internal/app/sql"
	craftnet_model "craftnet/internal/model"

	"github.com/samber/lo"
)

func GetUser(userID int) (*model.User, *craftnet_model.Error) {
	user, craftnetError := sql.GetUserByID(userID)

	if !lo.IsNil(craftnetError) {
		return nil, craftnetError
	}

	return user, nil
}
