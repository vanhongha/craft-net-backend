package sql

import (
	"craftnet/graph/model"
	"craftnet/internal/db"
	craftnet_model "craftnet/internal/model"
	"craftnet/internal/util"
	"database/sql"
	"errors"

	"github.com/samber/lo"
)

func GetMedia(mediaID int) (*model.Media, *craftnet_model.Error) {
	var media model.Media
	query := `
	SELECT
		id,
		path
	FROM media
	WHERE id = ?`
	if err := db.Instance.QueryRow(query, mediaID).
		Scan(&media.ID,
			&media.URLPath); !lo.IsNil(err) {
		if errors.Is(err, sql.ErrNoRows) {
			errMsg := util.ErrorMessage(util.ERROR_CODE[util.FAIL_TO_GET], "media")
			util.GetLogger().LogErrorWithMsgAndError(errMsg, err, false)
			return nil, nil
		}

		errMsg := util.ErrorMessage(util.ERROR_CODE[util.INTERNAL_SERVER], nil)
		util.GetLogger().LogErrorWithMsgAndError(errMsg, err, false)

		return nil, &craftnet_model.Error{
			Message: errMsg,
			Error:   err,
		}
	}
	return &media, nil
}
