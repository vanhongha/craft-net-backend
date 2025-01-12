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

func GetUserByID(userId int) (*model.User, *craftnet_model.Error) {
	var user model.User
	query := `
	SELECT
		id,
		last_name,
		first_name,
		date_of_birth,
		email,
		phone_number,
		avatar_img_path,
		status
	FROM users
	WHERE id = ? AND (status = ? OR status = ?)`
	if err := db.Instance.QueryRow(query, userId, util.STATUS_ACTIVATE, util.STATUS_IDLE).
		Scan(&user.ID,
			&user.LastName,
			&user.FirstName,
			&user.DateOfBirth,
			&user.Email,
			&user.PhoneNumber,
			&user.AvatarImgPath,
			&user.Status); !lo.IsNil(err) {
		if errors.Is(err, sql.ErrNoRows) {
			errMsg := util.ErrorMessage(util.ERROR_CODE[util.FAIL_TO_GET], "user")
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
	return &user, nil
}
