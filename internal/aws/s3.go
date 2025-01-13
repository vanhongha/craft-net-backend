package aws

import (
	"craftnet/config"
	"craftnet/internal/util"
	"fmt"
	"time"

	craftnet_model "craftnet/internal/model"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"github.com/samber/lo"
)

var Instance *s3.S3
var Sess *session.Session
var Downloader *s3manager.Downloader

func InitAws() {
	region := config.GetAwsRegion()

	var err error
	Sess, err = session.NewSession(&aws.Config{
		Region: aws.String(region),
	})

	if !lo.IsNil(err) {
		errMsg := util.ErrorMessage(util.ERROR_CODE[util.FAIL_TO_GENERATE], "session")
		fmt.Println(errMsg, err)
		util.GetLogger().LogErrorWithMsgAndError(errMsg, err, false)
		return
	}

	Instance = s3.New(Sess)
	Downloader = s3manager.NewDownloader(Sess)

	util.GetLogger().LogInfo("Init AWS S3 successfully")
}

func GetFile(fileName string) (string, *craftnet_model.Error) {
	bucket := config.GetAwsBucket()
	item := fileName

	params := &s3.GetObjectInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(item),
	}

	req, _ := Instance.GetObjectRequest(params)

	url, err := req.Presign(24 * 60 * time.Minute)

	if !lo.IsNil(err) {
		errMsg := util.ErrorMessage(util.ERROR_CODE[util.FAIL_TO_GET], "file")
		fmt.Println(errMsg, err)
		util.GetLogger().LogErrorWithMsgAndError(errMsg, err, false)
		return "", &craftnet_model.Error{
			Message: errMsg,
			Error:   err,
		}
	}

	return url, nil
}
