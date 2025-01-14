package services

import (
	"craftnet/graph/model"
	"craftnet/internal/app/sql"
	craftnet_model "craftnet/internal/model"

	"github.com/samber/lo"
)

func GetMedia(mediaID int) (*model.Media, *craftnet_model.Error) {
	if lo.IsNil(mediaID) {
		return nil, nil
	}

	media, craftnetError := sql.GetMedia(mediaID)

	if !lo.IsNil(craftnetError) {
		return nil, craftnetError
	}

	return media, nil
}
