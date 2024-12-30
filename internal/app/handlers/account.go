package handlers

import (
	"craftnet/graph/model"
	"craftnet/internal/app/services"
	craftnet_model "craftnet/internal/model"
)

func RegisterAccountHandler(username string, password string) (*model.Account, *craftnet_model.Error) {
	return services.RegisterAccount(username, password)
}
